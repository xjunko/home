module main

import internal { Configuration }
import magi

pub const used = internal.used

fn main() {
	mut config := Configuration.create('config.json')

	// General
	config.expects('instance.name', 'xjunko')
	config.expects('instance.type', 'magi')
	config.expects('instance.version', '0.2.3')

	// Domain
	config.expects('instance.domain', 'https://konno.ovh')

	// Folder
	config.expects('folder.root', 'static')
	config.expects('folder.entry', 'entry')

	// Discord
	config.expects('discord.token', '<INSERT DISCORD BOT TOKEN>')
	config.expects('discord.channel', '<INSERT CHANNEL ID>')

	// Paging
	config.expects('website.page.limit', f32(20)) // 20 per page.

	// Misc flag
	config.expects('website.channel.enable', true)
	config.expects('website.channel.media', true)

	// Discord
	config.expects('website.channel.discord', true)

	//
	config.load()
	config.save()

	//
	magi.execute(config)
}
