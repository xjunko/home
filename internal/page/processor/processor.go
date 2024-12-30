package processor

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type IProcessor interface {
	PreProcess(string) string
	Process(string) string
	PostProcess(string) string
}

type Processor struct {
	BaseProcessor

	Database   *gorm.DB
	processors []IProcessor
}

func (p *Processor) InitializeProcessor() error {
	if chanProc, err := NewChanStyleProcessor(); err == nil {
		p.processors = append(p.processors, chanProc)
	}

	if mediaProc, err := NewMediaProcessor(); err == nil {
		p.processors = append(p.processors, mediaProc)
	}

	if ytProc, err := NewYoutubeProcessor(p.Database); err == nil {
		p.processors = append(p.processors, ytProc)
	}

	if spotifyProc, err := NewSpotifyProcessor(p.Database); err == nil {
		p.processors = append(p.processors, spotifyProc)
	}

	return nil
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
	databaseInstance, err := gorm.Open(sqlite.Open("eva.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}

	proc := &Processor{
		processors: make([]IProcessor, 0),
		Database:   databaseInstance,
	}

	if err := proc.InitializeProcessor(); err != nil {
		panic(err)
	}

	return proc
}
