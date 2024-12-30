package export

import (
	"eva/internal/config"
	"eva/internal/page"
	"eva/internal/page/processor"
	"eva/templates"
	"log"
	"text/template"
)

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

	log.Println("[Magi] Starting!")

	if is_channel_enabled {
		manager.ResolveChannel()
		log.Printf("[Magi] Channels: %v posts \n", len(manager.Channels))
	} else {
		manager.Channels = append(manager.Channels, page.EvaPage{})
	}

	manager.ResolvePage()
	log.Printf("[Magi] Page: %v pages \n", len(manager.Pages))

	manager.ResolveNote()
	log.Printf("[Magi] Notes: %v notes \n", len(manager.Notes))

	// Exports
	manager.ExportRSS()
	manager.ExportPage()
	manager.ExportNote()
	manager.ExportChannel()

	return nil
}
