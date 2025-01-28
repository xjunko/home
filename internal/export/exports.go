package export

import (
	"fmt"
	"log"
	"os"
)

func (m *Magi) ExportRSS() {
	rssFile, err := os.OpenFile("dist/feed.xml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		log.Printf("[Magi] RSS template failed to load: %v \n", err)
		return
	}

	if err := m.Template.ExecuteTemplate(rssFile, "service/rss", m); err != nil {
		log.Printf("[Magi] Failed generating RSS: %v \n", err)
		return
	}
}

func (m *Magi) ExportPage() {
	for _, currentPage := range m.Pages {
		m.CurrentPage = &currentPage

		pageFile, err := os.OpenFile("dist"+currentPage.Metadata["route"], os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

		if err != nil {
			log.Printf("[Magi] Failed to open destination page file: %v \n", err)
			return
		}

		if _, err := m.Template.New("page_" + currentPage.Metadata["slug"]).Parse(currentPage.Content); err != nil {
			log.Printf("[Magi] Failed to parse page template: %v \n", err)
			return
		}

		if err := m.Template.ExecuteTemplate(pageFile, "page", m); err != nil {
			log.Printf("[Magi] Failed to generate the template: %v \n", err)
			return
		}
	}

}

func (m *Magi) ExportBlog() {
	for _, currentNote := range m.Notes {
		m.CurrentPage = &currentNote
		m.CurrentPage.Metadata["tags"] = "blog-read"

		noteFile, err := os.OpenFile("dist/blog/"+currentNote.Metadata["slog"]+".html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

		if err != nil {
			log.Printf("[Magi] Failed to open destination blog file: %v \n", err)
			return
		}

		if _, err := m.Template.New("note_" + currentNote.Metadata["slug"]).Parse(currentNote.Content); err != nil {
			log.Printf("[Magi] Failed to parse blog template: %v \n", err)
			return
		}

		if err := m.Template.ExecuteTemplate(noteFile, "page", m); err != nil {
			log.Printf("[Magi] Failed to generate the blog template: %v \n", err)
			return
		}
	}
}

func (m *Magi) GetChannelPages() map[int][]string {
	perPage := 20
	totalPages := (len(m.Channels) / perPage) + 1

	channelPages := make(map[int][]string)

	for i := 0; i < totalPages; i++ {
		channelPages[i+1] = make([]string, 0)
	}

	for i, currentPost := range m.Channels {
		postPage := (i / perPage) + 1
		channelPages[postPage] = append(channelPages[postPage], currentPost.ID)
	}

	return channelPages
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
		m.Mode = "ContentOnly"
		m.CurrentChannel = i

		channelFile, err := os.OpenFile("dist/chan/"+fmt.Sprintf("%d", i+1)+".html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

		if err != nil {
			log.Printf("[Magi] Failed to open destination channel file: %v \n", err)
			return
		}

		if err := m.Template.ExecuteTemplate(channelFile, "page", m); err != nil {
			log.Printf("[Magi] Failed to generate the channel page: %v \n", err)
			return
		}
	}
}
