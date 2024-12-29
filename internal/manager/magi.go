package manager

import (
	"eva/internal/config"
	"eva/internal/page"
	"eva/internal/page/processor"
	"eva/templates"
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
	Template  *template.Template

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
		curPage := page.NewPage(m.Template, file)

		if err := curPage.Load(m.Processor); err != nil {
			fmt.Printf("[Magi] Page failed to load: %v \n", err)
			continue
		}

		if !curPage.ShouldExclude() {
			m.Pages = append(m.Pages, *curPage)
		}
	}
}

func (m *Magi) ResolveNote() {
	files, err := filepath.Glob("web/entries/notes/*.md")

	if err != nil {
		fmt.Println("[Magi] Failed to resolve Note!")
		return
	}

	for _, file := range files {
		curNote := page.NewPage(m.Template, file)

		if err := curNote.Load(m.Processor); err != nil {
			fmt.Printf("[Magi] Note failed to load: %v \n", err)
			continue
		}

		if !curNote.ShouldExclude() {
			m.Notes = append(m.Notes, *curNote)
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
		curPage := page.NewPage(m.Template, file)

		if err := curPage.Load(m.Processor); err != nil {
			fmt.Printf("[Magi] Channel Posts failed to resolved: %v \n", err)
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
	rssFile, err := os.OpenFile("dist/feed.xml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		fmt.Printf("[Magi] RSS template failed to load: %v \n", err)
		return
	}

	if err := m.Template.ExecuteTemplate(rssFile, "service/rss", m); err != nil {
		fmt.Printf("[Magi] Failed generating RSS: %v \n", err)
		return
	}
}

func (m *Magi) ExportPage() {
	for _, currentPage := range m.Pages {
		m.CurrentPage = &currentPage
		m.Mode = "Normal"

		pageFile, err := os.OpenFile("dist"+currentPage.Metadata["route"], os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

		if err != nil {
			fmt.Printf("[Magi] Failed to open destination page file: %v \n", err)
			return
		}

		if _, err := m.Template.New("page_" + currentPage.Metadata["slug"]).Parse(currentPage.Content); err != nil {
			fmt.Printf("[Magi] Failed to parse page template: %v \n", err)
			return
		}

		if err := m.Template.ExecuteTemplate(pageFile, "page", m); err != nil {
			fmt.Printf("[Magi] Failed to generate the template: %v \n", err)
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
	manager.Template = template.New("")

	templates.BindFunctions(manager.Template)

	if err := templates.BindTemplates(manager.Template); err != nil {
		panic(err)
	}

	is_channel_enabled, _ := config.GetAsBool("Instance.Channel.Enabled")

	fmt.Println("[Magi] Starting!")

	if is_channel_enabled {
		manager.ResolveChannel()
		fmt.Printf("[Magi] Channels: %v posts \n", len(manager.Channels))
	} else {
		manager.Channels = append(manager.Channels, page.EvaPage{})
	}

	manager.ResolvePage()
	fmt.Printf("[Magi] Page: %v pages \n", len(manager.Pages))

	manager.ResolveNote()
	fmt.Printf("[Magi] Notes: %v notes \n", len(manager.Notes))

	manager.ProcessChannel()

	// Exports
	manager.ExportRSS()
	manager.ExportPage()

	return nil
}
