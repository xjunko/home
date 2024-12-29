package export

import (
	"eva/internal/config"
	"eva/internal/page"
	"eva/internal/page/processor"
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

func (m *Magi) GetCurrentPage() *page.EvaPage {
	return m.CurrentPage
}

func (m *Magi) GetConfig() *config.Config {
	return m.Config
}

func (m *Magi) GetPages() []page.EvaPage {
	return m.Pages
}

func (m *Magi) GetChannels() []page.EvaPage {
	return m.Channels
}

func (m *Magi) GetNotes() []page.EvaPage {
	return m.Notes
}

func (m *Magi) GetTemplate() *template.Template {
	return m.Template
}

func (m *Magi) GetLatestNote() *page.EvaPage {
	if len(m.Notes) == 0 {
		return nil
	}

	return &m.Notes[len(m.Notes)-1]
}

func (m *Magi) GetLatestPost() *page.EvaPage {
	if len(m.Channels) == 0 {
		return nil
	}

	return &m.Channels[len(m.Channels)-1]
}
