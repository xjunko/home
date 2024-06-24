module magi

import db.sqlite
import processor

pub struct Casper {
mut:
	database sqlite.DB
pub mut:
	spotify   processor.SpotifyProcessor
	youtube   processor.YoutubeProcessor
	media     processor.MediaProcessor
	greentext processor.GreentextProcessor
	reference processor.ReferenceProcessor
	discord   processor.DiscordCDNProcessor
}

pub fn Casper.create() !Casper {
	mut casper := Casper{
		database: sqlite.connect('db.sqlite')!
		spotify: processor.SpotifyProcessor.create()!
		youtube: processor.YoutubeProcessor.create()!
		media: processor.MediaProcessor.create()!
		greentext: processor.GreentextProcessor.create()!
		reference: processor.ReferenceProcessor.create()!
		discord: processor.DiscordCDNProcessor.create()!
	}

	// Create table if doesnt exists
	sql casper.database {
		create table processor.Track
	}!

	sql casper.database {
		create table processor.YoutubeThumbnail
	}!

	return casper
}

// Do
pub fn (mut casper Casper) preprocess(mut post Post) string {
	// Even more simpler
	discord_cdn_fix := casper.discord.process(post.content)

	// Simpler handler first.
	gt_text := casper.greentext.process(discord_cdn_fix)
	reference_info := casper.reference.first_pass(gt_text)

	// Bit more complex
	m_text := casper.media.process(reference_info[0])
	yt_text := casper.youtube.process(m_text, mut casper.database)
	sptfy_text := casper.spotify.process(yt_text, mut casper.database)

	post.has_reference = reference_info[1] == 'true'

	return sptfy_text
}

pub fn (mut casper Casper) postprocess(mut posts []Post) {
	mut ref := []processor.IPost{}

	for mut post in posts {
		ref << processor.IPost(post)
	}

	for mut post in posts {
		if post.has_reference {
			post.content = casper.reference.final_pass(post.content, mut ref)
		}
	}
}
