package processor

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/webp"
	"gopkg.in/vansante/go-ffprobe.v2"

	"gorm.io/gorm"
)

type MediaType int

const (
	Image MediaType = iota
	Video
	Emoji
)

type MediaInfo struct {
	gorm.Model

	URL    string `gorm:"primaryKey"`
	Width  int
	Height int
	Mime   string
}

func (info *MediaInfo) Setup() {
	// get file and handle based on mimetypes
	// first get mimetype
	log.Println("[MediaProcessor] Fetching media info for URL:", info.URL)
	mmt_req, err := http.NewRequest("GET", info.URL, nil)
	if err != nil {
		log.Fatal(err)
	}
	mmt_req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/114.0.0.0 Safari/537.36")
	mmt_req.Header.Set("Range", "bytes=0-1")

	mmt_resp, err := http.DefaultClient.Do(mmt_req)
	if err != nil {
		log.Fatal(err)
	}
	info.Mime = strings.TrimSpace(mmt_resp.Header.Get("Content-Type"))

	if len(info.Mime) == 0 {
		log.Printf("[MediaProcessor] Invalid mimetype, defaulting to image/png.")
		info.Mime = "image/png"
	}

}

func (info *MediaInfo) resolveImage() {
	// simply download the image in memory and read
	resp, err := http.Get(info.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.DecodeConfig(bytes.NewReader(buf))
	if err != nil {
		log.Fatal(err)
	}

	info.Width = img.Width
	info.Height = img.Height
}

func (info *MediaInfo) resolveVideo() {
	ctx, cancel_fn := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel_fn()

	data, err := ffprobe.ProbeURL(ctx, info.URL)
	if err != nil {
		log.Fatal(err)
	}

	if len(data.Streams) > 0 {
		info.Width = int(data.Streams[0].Width)
		info.Height = int(data.Streams[0].Height)
	}
}

func (info *MediaInfo) Resolve() {
	if info.IsVideo() {
		info.resolveVideo()
	} else {
		info.resolveImage()
	}
}

func IsVideoURL(url string) bool {
	VIDEO_FORMAT := []string{"mp4", "webm", "mov"}

	for _, fmt := range VIDEO_FORMAT {
		if strings.HasSuffix(url, fmt) {
			return true
		}
	}

	return false
}

func (info *MediaInfo) IsVideo() bool {
	return IsVideoURL(info.URL)
}

func (info *MediaInfo) GetType() MediaType {
	if info.IsVideo() {
		return Video
	}

	if strings.Contains(info.URL, "/emojis/") {
		return Emoji
	}

	return Image
}

type MediaProcessor struct {
	BaseProcessor

	database *gorm.DB
	pattern  *regexp.Regexp
}

func NewMediaProcessor(database *gorm.DB) (*MediaProcessor, error) {
	if err := database.AutoMigrate(&MediaInfo{}); err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`https?://\S+\.(?:(png)|(jpe?g)|(gif)|(svg)|(webp)|(mp4)|(webm)|(mov))(?:\?\S*)?
`)
	return &MediaProcessor{pattern: re, database: database}, nil
}

func (p *MediaProcessor) Process(text string) string {
	return p.pattern.ReplaceAllStringFunc(text, func(match string) string {
		link := strings.TrimSpace(match)
		info := p.getMediaInfo(link)

		switch info.GetType() {
		case Video:
			return fmt.Sprintf("\n<video muted autoplay loop controls preload=metadata width=\"%d\" height=\"auto\" src=\"%s\"></video>\n", info.Width, link)
		case Emoji:
			return fmt.Sprintf("\n<img class=\"discord-emoji\" loading=lazy alt=\"\" src=\"%s?size=32&quality=lossless\">\n", link)
		case Image:
			return fmt.Sprintf("\n<img loading=lazy alt=\"\" width=\"%d\" height=\"auto\" src=\"%s\">\n", info.Width, link)
		default:
			panic(fmt.Sprintf("Unknown media type for URL: %s", link))
		}
	})
}

func (p *MediaProcessor) getMediaInfo(url string) MediaInfo {
	info := p.getMediaInfoFromDB(url)

	if len(info.URL) == 0 {
		info.URL = url

		// first get the mimetype
		info.Setup()

		// then we handle the file based on mimetype
		info.Resolve()

		p.database.Create(&info)
	}

	return info
}

func (p *MediaProcessor) getMediaInfoFromDB(url string) MediaInfo {
	var info MediaInfo
	p.database.First(&info, "url = ?", url)
	return info
}
