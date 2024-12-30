package processor

import (
	"strings"
)

type ChanStyleProcessor struct {
	BaseProcessor
}

func (p *ChanStyleProcessor) Process(text string) string {

	lines := strings.Split(text, "\n")
	results := make([]string, 0)

	for _, line := range lines {
		// Reference
		if startsWith(line, ">>") {
			postID := strings.TrimSpace(strings.Split(line, ">>")[1])

			refHTML := "<a style='color: var(--red-text)' href='/redirect.html?id=" + postID + "'>" + line + "</a>"

			results = append(results, refHTML)
			continue
		}

		// *chan styling
		trimmedLine := strings.TrimSpace(line)

		switch {
		case startsWith(trimmedLine, ">"):
			results = append(results, "<a style='color: var(--green-text)'>\\"+trimmedLine+"</a><br/>")
		case startsWith(trimmedLine, "<<"):
			results = append(results, "<a style='color: var(--red-text)'>\\"+trimmedLine[1:]+"</a><br/>")
		default:
			results = append(results, line)
		}
	}

	return strings.Join(results, "\n")
}

func startsWith(text string, prefix string) bool {
	return strings.HasPrefix(strings.TrimSpace(text), prefix)
}

func NewChanStyleProcessor() (*ChanStyleProcessor, error) {
	return &ChanStyleProcessor{}, nil
}
