package processor

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
	"text/template"

	"gorm.io/gorm"
)

type YoutubeInfo struct {
	gorm.Model

	ID           string `gorm:"primaryKey"`
	ThumbnailURL string
	URL          string
}

type YoutubeProcessor struct {
	BaseProcessor

	database *gorm.DB
	pattern  *regexp.Regexp
}

func (youtube *YoutubeProcessor) Process(text string) string {
	return youtube.pattern.ReplaceAllStringFunc(text, func(match string) string {
		videoID := youtube.extractVideoID(match)
		videoInfo := youtube.getVideoInfo(videoID)

		templateEngine, err := template.New("youtube.tmpl").ParseFiles("templates/widget/socials/youtube.tmpl")

		if err != nil {
			return fmt.Sprintf("Error: %v", err)
		}

		var buf bytes.Buffer
		if err := templateEngine.Execute(&buf, videoInfo); err != nil {
			return fmt.Sprintf("Error: %v", err)
		}

		return buf.String()
	})
}

func (youtube *YoutubeProcessor) extractVideoID(url string) string {
	matches := youtube.pattern.FindStringSubmatch(url)

	if len(matches) > 1 {
		return matches[1]
	}

	return ""
}

func (youtube *YoutubeProcessor) getVideoInfo(videoID string) YoutubeInfo {
	info := youtube.getVideoThumbnailFromDB(videoID)

	if len(info.ThumbnailURL) == 0 {
		info = youtube.getVideoThumbnailFromID(videoID)
		youtube.database.Create(&info)
		fmt.Printf("[Youtube] Added %v into the database! \n", info.ID)
	}

	return info
}

func (youtube *YoutubeProcessor) getVideoThumbnailFromID(videoID string) YoutubeInfo {
	info := YoutubeInfo{ID: videoID, URL: fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)}
	info.ThumbnailURL = fmt.Sprintf("https://i3.ytimg.com/vi/%s/0.jpg", videoID) // Fallback to the worst quality

	// Then gradually find the best quality.
	qualities := []string{"maxresdefault.jpg", "mqdefault.jpg", "0.jpg"}
	for _, quality := range qualities {
		thumbnailURL := fmt.Sprintf("https://i3.ytimg.com/vi/%s/%s", videoID, quality)
		if isValidThumbnail(thumbnailURL) {
			info.ThumbnailURL = thumbnailURL
			break
		}
	}

	return info
}

func (youtube *YoutubeProcessor) getVideoThumbnailFromDB(videoID string) YoutubeInfo {
	var info YoutubeInfo

	youtube.database.First(&info, "id = ?", videoID)

	return info
}

func isValidThumbnail(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200
}

func NewYoutubeProcessor(database *gorm.DB) (*YoutubeProcessor, error) {
	database.AutoMigrate(&YoutubeInfo{})

	pattern := regexp.MustCompile(`https?://(?:www\.)?youtu(?:be\.com/watch\?v=|\.be/)(\S+)`)
	return &YoutubeProcessor{pattern: pattern, database: database}, nil
}
