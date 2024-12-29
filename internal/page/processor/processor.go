package processor

type IProcessor interface {
	PreProcess(string) string
	Process(string) string
	PostProcess(string) string
}

type Processor struct {
	BaseProcessor

	processors []IProcessor
}

func (p *Processor) PreProcess(content string) string {
	for _, proc := range p.processors {
		content = proc.PreProcess(content)
	}

	return content
}

func (p *Processor) Process(content string) string {
	for _, proc := range p.processors {
		content = proc.Process(content)
	}

	return content
}

func (p *Processor) PostProcess(content string) string {
	for _, proc := range p.processors {
		content = proc.PostProcess(content)
	}

	return content
}

func NewProcessor() *Processor {
	proc := &Processor{
		processors: make([]IProcessor, 0),
	}

	if mediaProc, err := NewMediaProcessor(); err == nil {
		proc.processors = append(proc.processors, mediaProc)
	}

	if ytProc, err := NewYoutubeProcessor(); err == nil {
		proc.processors = append(proc.processors, ytProc)
	}

	if spotifyProc, err := NewSpotifyProcessor(); err == nil {
		proc.processors = append(proc.processors, spotifyProc)
	}

	return proc
}
