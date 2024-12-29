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
		if startsWith(line, ">") && !startsWith(line, ">>") {
			results = append(results, "<a style='color: var(--green-text)'>\\"+line+"</a><br/>")
		} else if startsWith(line, "<<") {
			results = append(results, "<a style='color: var(--red-text)'>\\"+line[1:]+"</a><br/>")
		} else {
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
