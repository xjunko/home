package main

import (
	"eva/internal/config"
	"eva/internal/export"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func watchFolder(watcher *fsnotify.Watcher, dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != dir {
			err := watcher.Add(path)
			if err != nil {
				log.Printf("Failed to add directory: %v", err)
			}
		}
		return nil
	})

}

func liveFileWatcher(config *config.Config) {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		panic(err)
	}

	defer watcher.Close()

	go func() {
		for {
			select {
			case evt, ok := <-watcher.Events:
				if !ok {
					return
				}

				if evt.Has(fsnotify.Write) {
					log.Println("modified file:", evt.Name)

					if err := export.Execute(config); err != nil {
						log.Panic(err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	if err := watchFolder(watcher, "templates"); err != nil {
		panic(err)
	}

	if err := watchFolder(watcher, "entries"); err != nil {
		panic(err)
	}

	log.Println("Watching for changes...")

	<-make(chan struct{})
}

func main() {
	config := config.NewConfig()

	config.Expect("Instance.Name", "junko")
	config.Expect("Instance.Type", "Eva")
	config.Expect("Instance.Description", "simple microblogging processor")
	config.Expect("Instance.Version", "0.5.0")

	config.Expect("Instance.Domain", "https://kafu.ovh")

	config.Expect("Instance.Channel.Enabled", true)
	config.Expect("Instance.Channel.Media", true)
	config.Expect("Instance.Channel.Page.Limit", 20)

	config.Expect("Instance.Channel.Discord", true)
	config.Expect("Instance.Channel.Discord.Token", "your_token")
	config.Expect("Instance.Channel.Discord.Channel", "channel_id")

	config.Expect("Exporter.Quiet", false)

	if err := config.Load(); err != nil {
		panic(err)
	}

	liveFileWatcher(config)
}
