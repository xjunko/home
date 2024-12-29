package processor

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"regexp"

	"gorm.io/gorm"
)

type RootCoverArt struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type RootArtist struct {
	ID      string `json:"id"`
	Profile struct {
		Name string `json:"name"`
	} `json:"profile"`
}

type RootTrack struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	URI          string `json:"uri"`
	AlbumOfTrack struct {
		CoverArt struct {
			Sources []RootCoverArt `json:"sources"`
		} `json:"coverArt"`
	} `json:"albumOfTrack"`

	Previews struct {
		AudioPreviews struct {
			Items []struct {
				URL string `json:"url"`
			} `json:"items"`
		} `json:"audioPreviews"`
	}

	FirstArtist struct {
		Items []RootArtist `json:"items"`
	} `json:"firstArtist"`
}

type Root struct {
	Entities struct {
		Items map[string]RootTrack `json:"items"`
	} `json:"entities"`
}

type Track struct {
	gorm.Model

	ID              string `gorm:"primaryKey" json:"id"`
	Name            string `json:"name"`
	SourceURL       string `json:"source_url"`
	Artist          string `json:"artist"`
	ArtistID        string `json:"artist_id"`
	CoverArtURL     string `json:"cover_art_url"`
	AudioPreviewURL string `json:"audio_preview_url"`
}

const (
	sptfyURLRe    = `https?://open\.spotify\.com/track/(\w+)`
	sptfyScriptRe = `<script\s+id="initial-state"\s+type="text/plain">([^<]+)</script>`
	sptfyErrorMsg = "[Spotify]: Failed to embed url. Reason: "
)

type SpotifyProcessor struct {
	BaseProcessor

	database *gorm.DB

	Pattern       *regexp.Regexp
	ScriptPattern *regexp.Regexp
}

func (spotify *SpotifyProcessor) LargestCoverArt(sources []RootCoverArt) (string, error) {
	if len(sources) == 0 {
		return "", errors.New("no cover art found")
	}

	largestArt := sources[0]
	for i := 1; i < len(sources); i++ {
		size := sources[i].Width * sources[i].Height
		rootSize := largestArt.Width * largestArt.Height
		if size > rootSize {
			largestArt = sources[i]
		}
	}

	return largestArt.URL, nil
}

func (spotify *SpotifyProcessor) GetTrack(url string) (Track, error) {
	var track Track

	result := spotify.database.Where("source_url = ?", url).First(&track)

	if result.Error != nil {
		fetchedTrack, err := spotify.GetTrackFromURL(url)

		if err != nil {
			return Track{}, err
		}

		spotify.database.Create(&fetchedTrack)

		fmt.Printf("[Spotify]: Fetched track: %s\n", fetchedTrack.Name)
	}

	return track, nil
}

func (spotify *SpotifyProcessor) GetTrackFromURL(url string) (Track, error) {
	resp, err := http.Get(url)
	if err != nil {
		return Track{}, fmt.Errorf(sptfyErrorMsg+"Failed to reach url: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Track{}, fmt.Errorf(sptfyErrorMsg+"Failed to read response body: %v", err)
	}

	matches := spotify.ScriptPattern.FindSubmatch(body)
	if matches == nil {
		return Track{}, errors.New(sptfyErrorMsg + "Embedded script not found")
	}

	base64Data := string(matches[1])
	musicData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return Track{}, fmt.Errorf(sptfyErrorMsg+"Failed to decode base64 data: %v", err)
	}

	var decodedData Root
	err = json.Unmarshal(musicData, &decodedData)
	if err != nil {
		return Track{}, fmt.Errorf(sptfyErrorMsg+"Failed to decode JSON data: %v", err)
	}

	for _, track := range decodedData.Entities.Items {
		if len(track.Previews.AudioPreviews.Items) > 0 {
			audioPreview := track.Previews.AudioPreviews.Items[0]
			if len(track.FirstArtist.Items) > 0 {
				artist := track.FirstArtist.Items[0]

				thumbnail, err := spotify.LargestCoverArt(track.AlbumOfTrack.CoverArt.Sources)
				if err != nil {
					return Track{}, fmt.Errorf(sptfyErrorMsg+"Failed to get thumbnail: %v", err)
				}

				fetchedTrack := Track{
					SourceURL:       url,
					ID:              track.ID,
					Name:            track.Name,
					Artist:          artist.Profile.Name,
					ArtistID:        artist.ID,
					CoverArtURL:     thumbnail,
					AudioPreviewURL: audioPreview.URL,
				}

				return fetchedTrack, nil
			}
		}
	}

	return Track{}, errors.New("failed to parse track data")
}

func (spotify *SpotifyProcessor) HandleURL(url string) (string, error) {
	templateEngine, err := template.New("spotify.tmpl").ParseFiles("templates/widget/socials/spotify.tmpl")

	if err != nil {
		return "", err
	}

	track, err := spotify.GetTrack(url)

	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	if err := templateEngine.Execute(&buf, track); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (spotify *SpotifyProcessor) Process(text string) string {
	return spotify.Pattern.ReplaceAllStringFunc(text, func(url string) string {
		embedInfo, err := spotify.HandleURL(url)

		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}

		return embedInfo
	})
}

func NewSpotifyProcessor(database *gorm.DB) (*SpotifyProcessor, error) {
	database.AutoMigrate(&Track{})

	urlPattern := regexp.MustCompile(sptfyURLRe)
	scriptPattern := regexp.MustCompile(sptfyScriptRe)

	return &SpotifyProcessor{
		database:      database,
		Pattern:       urlPattern,
		ScriptPattern: scriptPattern,
	}, nil
}
