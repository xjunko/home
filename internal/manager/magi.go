package manager

import (
	"eva/internal/config"
	"eva/internal/page"
	"eva/internal/page/processor"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"text/template"
)

type Magi struct {
	Config *config.Config

	Pages    []page.EvaPage
	Channels []page.EvaPage
	Notes    []page.EvaPage

	Processor *processor.Processor

	CurrentPage *page.EvaPage
	Mode        string
}

func (m *Magi) ResolvePage() {
	files, err := filepath.Glob("web/entries/pages/*.md")

	if err != nil {
		fmt.Println("[Magi] Failed to resolve page!")
		return
	}

	for _, file := range files {
		curPage := page.NewPage(file)

		if err := curPage.Load(m.Processor); err != nil {
			fmt.Printf("[Magi] %v \n", err)
			continue
		}

		if !curPage.ShouldExclude() {
			m.Pages = append(m.Pages, *curPage)
		}
	}
}

func (m *Magi) ResolveChannel() {
	files, err := filepath.Glob("web/entries/channels/*.md")

	if err != nil {
		fmt.Println("[Magi] Failed to resolve channel!")
		return
	}

	for _, file := range files {
		curPage := page.NewPage(file)

		if err := curPage.Load(m.Processor); err != nil {
			fmt.Printf("[Magi] %v \n", err)
			continue
		}

		if !curPage.ShouldExclude() {
			m.Channels = append(m.Channels, *curPage)
		}
	}

	sort.Slice(m.Channels, func(i, j int) bool {
		return m.Channels[i].PostedAt.Before(m.Channels[j].PostedAt)
	})
}

func (m *Magi) ProcessChannel() {
	// TODO
}

func (m *Magi) ExportRSS() {
	funcs := template.FuncMap{
		"sub": func(a, b int) int {
			return a - b
		},
	}

	templateEngine := template.New("feed.tmpl").Funcs(funcs)
	rssTemplate, err := templateEngine.ParseFiles("web/templates/rss/feed.tmpl")

	if err != nil {
		fmt.Printf("[Magi] %v \n", err)
		return
	}

	rssFile, err := os.OpenFile("dist/feed.xml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		fmt.Printf("[Magi] %v \n", err)
		return
	}

	if err := rssTemplate.Execute(rssFile, m); err != nil {
		fmt.Printf("[Magi] %v \n", err)
		return
	}
}

func (m *Magi) ExportPage() {
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
		return
	}

	templateEngine := template.New("base.tmpl")

	pageTemplate, err := templateEngine.ParseFiles(templates...)

	if err != nil {
		fmt.Printf("[Magi] %v \n", err)
		return
	}

	for _, currentPage := range m.Pages {
		m.CurrentPage = &currentPage
		m.Mode = "Normal"

		pageFile, err := os.OpenFile("dist"+currentPage.Metadata["route"], os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

		if err != nil {
			fmt.Printf("[Magi] %v \n", err)
			return
		}

		if err := pageTemplate.Execute(pageFile, m); err != nil {
			fmt.Printf("[Magi] %v \n", err)
			return
		}
	}

}

func Execute(config *config.Config) error {
	manager := &Magi{
		Config:   config,
		Pages:    make([]page.EvaPage, 0),
		Channels: make([]page.EvaPage, 0),
		Notes:    make([]page.EvaPage, 0),
	}

	manager.Processor = processor.NewProcessor()

	is_channel_enabled, _ := config.GetAsBool("Instance.Channel.Enabled")

	fmt.Println("[Magi] Starting!")
	fmt.Printf("[Magi] Channel: %v \n", is_channel_enabled)

	if is_channel_enabled {
		fmt.Println("[Magi] Resolving Channel!")
		manager.ResolveChannel()
		fmt.Println("[Magi] Channel resolved!")
	} else {
		manager.Channels = append(manager.Channels, page.EvaPage{})
	}

	fmt.Println("[Magi] Resolving Page!")
	manager.ResolvePage()
	fmt.Println("[Magi] Page resolved!")

	manager.ProcessChannel()

	// Exports
	manager.ExportRSS()
	manager.ExportPage()

	return nil
}
