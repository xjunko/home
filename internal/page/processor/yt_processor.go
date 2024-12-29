package processor

import (
	"fmt"
	"net/http"
	"regexp"
)

type YoutubeProcessor struct {
	pattern *regexp.Regexp
}

func (youtube *YoutubeProcessor) PreProcess(text string) string {
	return text
}

// Process finds the video ID from the URL and fetches the thumbnail.
func (youtube *YoutubeProcessor) Process(text string) string {
	return youtube.pattern.ReplaceAllStringFunc(text, func(match string) string {
		videoID := youtube.extractVideoID(match)
		thumbnailURL := youtube.getVideoThumbnailFromID(videoID)
		return fmt.Sprintf(`<img src="%s" alt="YouTube Video Thumbnail">`, thumbnailURL)
	})
}

func (youtube *YoutubeProcessor) PostProcess(text string) string {
	return text
}

func (youtube *YoutubeProcessor) extractVideoID(url string) string {
	matches := youtube.pattern.FindStringSubmatch(url)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func (youtube *YoutubeProcessor) getVideoThumbnailFromID(videoID string) string {
	qualities := []string{"maxresdefault.jpg", "mqdefault.jpg", "0.jpg"}

	for _, quality := range qualities {
		thumbnailURL := fmt.Sprintf("https://i3.ytimg.com/vi/%s/%s", videoID, quality)
		if isValidThumbnail(thumbnailURL) {
			return thumbnailURL
		}
	}

	return fmt.Sprintf("https://i3.ytimg.com/vi/%s/0.jpg", videoID)
}

func isValidThumbnail(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200
}

func NewYoutubeProcessor() (*YoutubeProcessor, error) {
	pattern := regexp.MustCompile(`https?://(?:www\.)?youtu(?:be\.com/watch\?v=|\.be/)(\S+)`)
	return &YoutubeProcessor{pattern: pattern}, nil
}
