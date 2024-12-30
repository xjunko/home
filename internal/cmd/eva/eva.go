package main

import (
	"eva/internal/config"
	"eva/internal/export"
)

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

	if err := config.Load(); err != nil {
		panic(err)
	}

	if err := export.Execute(config); err != nil {
		panic(err)
	}
}
