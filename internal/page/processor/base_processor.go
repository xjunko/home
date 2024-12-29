package processor

type BaseProcessor struct{}

func (p *BaseProcessor) PreProcess(content string) string {
	return content
}

func (p *BaseProcessor) Process(content string) string {
	return content
}

func (p *BaseProcessor) PostProcess(content string) string {
	return content
}
