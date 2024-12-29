package export

import (
	"fmt"
	"os"
	"strings"
)

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

		if strings.HasPrefix(currentPage.Metadata["tags"], "notes") {
			m.Mode = "NoteList"
		}

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

func (m *Magi) ExportNote() {
	for _, currentNote := range m.Notes {
		m.CurrentPage = &currentNote
		m.Mode = "Note"

		noteFile, err := os.OpenFile("dist/note/"+currentNote.Metadata["slog"]+".html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

		if err != nil {
			fmt.Printf("[Magi] Failed to open destination note file: %v \n", err)
			return
		}

		if _, err := m.Template.New("note_" + currentNote.Metadata["slug"]).Parse(currentNote.Content); err != nil {
			fmt.Printf("[Magi] Failed to parse note template: %v \n", err)
			return
		}

		if err := m.Template.ExecuteTemplate(noteFile, "page", m); err != nil {
			fmt.Printf("[Magi] Failed to generate the note template: %v \n", err)
			return
		}
	}
}

func (m *Magi) ExportChannel() {
	perPage := 20
	totalPages := (len(m.Channels) / perPage) + 1

	for _, currentPage := range m.Pages {
		if currentPage.Metadata["tags"] == "channel" {
			m.CurrentPage = &currentPage
		}
	}

	m.Mode = "Channel"

	for i := 0; i < totalPages; i++ {
		m.CurrentChannel = i

		channelFile, err := os.OpenFile("dist/chan/"+fmt.Sprintf("%d", i+1)+".html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

		if err != nil {
			fmt.Printf("[Magi] Failed to open destination channel file: %v \n", err)
			return
		}

		if err := m.Template.ExecuteTemplate(channelFile, "page", m); err != nil {
			fmt.Printf("[Magi] Failed to generate the channel page: %v \n", err)
			return
		}
	}
}
