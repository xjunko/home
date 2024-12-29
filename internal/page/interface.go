package page

import (
	"eva/internal/config"
	"text/template"
)

type IExporter interface {
	GetConfig() *config.Config
	GetCurrentPage() *EvaPage
	GetPages() []EvaPage
	GetChannels() []EvaPage
	GetNotes() []EvaPage
	GetTemplate() *template.Template
	GetLatestNote() *EvaPage
	GetLatestPost() *EvaPage
}
