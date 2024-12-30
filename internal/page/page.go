package page

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"eva/internal/page/processor"

	"gitlab.com/golang-commonmark/markdown"
)

// Eva specific code
var MARKDOWN = markdown.New(markdown.XHTMLOutput(false), markdown.HTML(true))

func toMarkdown(content string) string {
	return MARKDOWN.RenderToString([]byte(content))
}

var PREFIX = "@"
var PREFIXES = []string{
	// Page Basic Info
	"title",
	"description",
	"thumbnail",
	// Page Data
	"author",
	"date",
	"tags",
	"route",
	// /note/*
	"slog",
	// /channel/*
	"style",
	"outline",
	"outline-style",
	// Misc
	"exclude",
}

type EvaPageType int32

const (
	NOTE EvaPageType = iota
	CHANNEL
)

type EvaPage struct {
	path string

	ID       string
	PostedAt time.Time
	Content  string
	Metadata map[string]string

	Template *template.Template
	Exporter IExporter

	RawContent     string
	BeenReferenced bool
}

func (p *EvaPage) Load(curProcessor processor.IProcessor) error {
	content, err := os.ReadFile(p.path)

	if err != nil {
		return fmt.Errorf("failed to read the page file: %v", err)
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, PREFIX) {
			for _, prefix := range PREFIXES {
				if strings.HasPrefix(line, PREFIX+prefix) {
					p.Metadata[prefix] = strings.Split(line, "=")[1]
				}
			}
		} else {
			p.Content += line + "\n"
		}
	}

	p.RawContent = p.Content
	p.Content = curProcessor.Process(p.Content)

	// Resolve metadata
	if _, exists := p.Metadata["author"]; !exists {
		p.Metadata["author"] = "junko"
	}

	if _, exists := p.Metadata["thumbnail"]; exists {
		p.Metadata["filename"] = getFilename(p.Metadata["thumbnail"])
		p.Metadata["mimetype"] = getMimeType(p.Metadata["thumbnail"])

		if processor.IsVideoURL(p.Metadata["thumbnail"]) {
			p.Metadata["thumbnail-type"] = "video"
		} else {
			p.Metadata["thumbnail-type"] = "image"
		}
	}

	if tags, exists := p.Metadata["tags"]; exists {
		if strings.Contains(tags, "discord-post") {
			p.Metadata["style"] = "border: .1em solid #5865F2;"
		}
	}

	// Date fallbacks
	if p.PostedAt.Year() <= 2000 {
		if unix_time, exists := p.Metadata["date"]; exists {
			i, err := strconv.ParseInt(unix_time, 10, 64)

			if err != nil {
				return fmt.Errorf("failed to parse unix time: %v", err)
			}

			p.PostedAt = time.Unix(i, 0)
		}
	}

	return nil
}

func (p *EvaPage) ToMarkdown() string {
	return toMarkdown(p.Content)
}

func (p *EvaPage) GetContent() string {
	templateName := "internal.page_" + p.ID
	withAllTemplate, err := p.Template.New(templateName).Parse(p.Content)

	if err != nil {
		fmt.Printf("[Page] Failed to parse the page content: %v", err)
		return p.ToMarkdown()
	}

	var buf bytes.Buffer

	if err := withAllTemplate.ExecuteTemplate(&buf, templateName, p.Exporter); err != nil {
		fmt.Printf("[Page] Failed to execute the page content: %v", err)
		return p.ToMarkdown()
	}

	return toMarkdown(buf.String())
}

func (p *EvaPage) GetType() EvaPageType {
	if _, exists := p.Metadata["slog"]; exists {
		return NOTE
	}

	return CHANNEL
}

func (p *EvaPage) GetFormattedPostDate() string {
	return p.PostedAt.Format("Mon, Jan 2nd, 2006")
}

func (p *EvaPage) GetSimpleFormattedPostDate() string {
	return p.PostedAt.Format("2006-01-02")
}

func (p *EvaPage) GetEstimatedReadingTime() string {
	return fmt.Sprintf("%.2f minutes", float32(p.GetWords())/212.0)
}

func (p *EvaPage) ShouldExclude() bool {
	if _, exists := p.Metadata["exclude"]; exists {
		return true
	}

	return false
}

func (p *EvaPage) GetWords() int {
	return len(strings.Fields(p.Content))
}

func NewPage(exporter IExporter, path string) *EvaPage {
	filenameNoExt := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	unixTimefromFilename, err := strconv.ParseInt(filenameNoExt, 10, 64)

	if err != nil {
		unixTimefromFilename = 0
	}

	return &EvaPage{
		path:     path,
		ID:       filenameNoExt,
		PostedAt: time.Unix(unixTimefromFilename, 0),
		Metadata: make(map[string]string),
		Exporter: exporter,
		Template: exporter.GetTemplate(),
	}
}
