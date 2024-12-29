package export

import (
	"eva/internal/page"
	"fmt"
	"path/filepath"
	"sort"
)

func (m *Magi) ResolvePage() {
	files, err := filepath.Glob("templates/pages/*.md")

	if err != nil {
		fmt.Println("[Magi] Failed to resolve page!")
		return
	}

	for _, file := range files {
		curPage := page.NewPage(m, file)

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
	files, err := filepath.Glob("entries/notes/*.md")

	if err != nil {
		fmt.Println("[Magi] Failed to resolve Note!")
		return
	}

	for _, file := range files {
		curNote := page.NewPage(m, file)

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
	files, err := filepath.Glob("entries/channels/*.md")

	if err != nil {
		fmt.Println("[Magi] Failed to resolve channel!")
		return
	}

	for _, file := range files {
		curPage := page.NewPage(m, file)

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
