package export

import (
	"eva/internal/config"
	"eva/internal/page"
	"eva/internal/page/processor"
	"eva/templates"
	"fmt"
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

	// Exports
	manager.ExportRSS()
	manager.ExportPage()

	return nil
}
