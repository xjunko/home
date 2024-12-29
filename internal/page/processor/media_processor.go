package processor

import (
	"fmt"
	"regexp"
	"strings"
)

var VIDEO_FORMAT = []string{
	"mp4", "webm", "mov",
}

func IsVideoURL(url string) bool {
	for _, format := range VIDEO_FORMAT {
		if strings.HasSuffix(url, format) {
			return true
		}
	}
	return false
}

type MediaProcessor struct {
	BaseProcessor

	pattern *regexp.Regexp
}

func NewMediaProcessor() (*MediaProcessor, error) {
	re := regexp.MustCompile(`https?://\S+\.(?:(png)|(jpe?g)|(gif)|(svg)|(webp)|(mp4)|(webm)|(mov))(?:\?\S*)?
`)
	return &MediaProcessor{pattern: re}, nil
}

func (p *MediaProcessor) Process(text string) string {
	return p.pattern.ReplaceAllStringFunc(text, func(match string) string {
		link := strings.TrimSpace(match)

		if IsVideoURL(link) {
			return fmt.Sprintf("\n<video muted autoplay loop controls preload=metadata src=\"%s\"></video>\n", link)
		}

		if strings.Contains(link, "/emojis/") {
			return fmt.Sprintf("\n<img class=\"discord-emoji\" loading=lazy alt=\"\" src=\"%s?size=32&quality=lossless\">\n", link)
		}

		return fmt.Sprintf("\n<img loading=lazy alt=\"\" src=\"%s\">\n", link)
	})
}
