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

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"

	"eva/internal/config"
	"eva/internal/page/processor"
)

// Eva specific code
var MARKDOWN = goldmark.New(
	goldmark.WithExtensions(extension.GFM),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
	goldmark.WithRendererOptions(
		html.WithXHTML(),
		html.WithUnsafe(),
	),
)

func toMarkdown(content string) string {
	var buf bytes.Buffer

	if err := MARKDOWN.Convert([]byte(content), &buf); err != nil {
		fmt.Printf("[Page] Failed to convert markdown: %v", err)
		return content
	}

	return buf.String()
}

var PREFIX = "@"
var PREFIXES = []string{
	// Page Basic Info
	"title",
	"description",
	"thumbnail",
	// Page Data
	"author",
	"data",
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
	if p.PostedAt.Year() == 1970 {
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

func (p *EvaPage) GetContent(curConfig *config.Config, postOnChannels []EvaPage) string {
	templates := []string{}

	err := filepath.Walk("web/templates/html", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".tmpl" {
			templates = append(templates, path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("[Magi] No templates found!")
		return p.ToMarkdown()
	}

	funcs := template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
	}

	templateEngine := template.New(p.ID + ".md").Funcs(funcs)

	// Add all templates
	pageTemplate, err := templateEngine.ParseFiles(templates...)

	if err != nil {
		fmt.Printf("[Page] Failed to include templates: %v", err)
		return p.ToMarkdown()
	}

	withAllTemplate, err := pageTemplate.Parse(p.Content)

	if err != nil {
		fmt.Printf("[Page] Failed to parse the page content: %v", err)
		return p.ToMarkdown()
	}

	var buf bytes.Buffer
	context := struct {
		CurrentPage *EvaPage
		Config      *config.Config
		Channels    []EvaPage
	}{
		CurrentPage: p,
		Config:      curConfig,
		Channels:    postOnChannels,
	}

	if err := withAllTemplate.Execute(&buf, context); err != nil {
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

func (p *EvaPage) ShouldExclude() bool {
	if _, exists := p.Metadata["exclude"]; exists {
		return true
	}

	return false
}

func NewPage(path string) *EvaPage {
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
	}
}

func getFilename(url string) string {
	// Split URL by '/' and take the last part
	parts := strings.Split(url, "/")
	filename := parts[len(parts)-1]

	// Split by '?' to remove any query parameters
	if idx := strings.Index(filename, "?"); idx != -1 {
		filename = filename[:idx]
	}

	return filename
}

func getMimeType(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}
