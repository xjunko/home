package export

import (
	"fmt"
	"os"
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
