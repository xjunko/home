module processor

pub struct DiscordCDNProcessor {}

pub fn (mut discord DiscordCDNProcessor) process(text string) string {
	mut new_content := text

	if new_content.contains('cdn.discordapp.com') && !new_content.contains('emojis') {
		new_content = new_content.replace('cdn.discordapp.com', 'cdn.discordapp.xyz')
	}

	if new_content.contains('media.discordapp.net') && !new_content.contains('emojis') {
		new_content = new_content.replace('media.discordapp.net', 'cdn.discordapp.xyz')
	}

	return new_content
}

pub fn DiscordCDNProcessor.create() !DiscordCDNProcessor {
	return DiscordCDNProcessor{}
}
