module processor

import regex

const fmt_video_list = [
	'mp4',
	'webm',
	'mov',
]

const fmt_media_proc_re = r'https?://\S+\.(?:(png)|(jpe?g)|(gif)|(svg)|(webp)|(mp4)|(webm)|(mov))(?:\?\S*)?
'

pub struct MediaProcessor {
pub mut:
	pattern regex.RE
}

pub fn (mut media MediaProcessor) process(text string) string {
	return media.pattern.replace_by_fn(text, fn (_ regex.RE, text string, b1 int, b2 int) string {
		link := text[b1..b2].trim_space()

		for fmt in processor.fmt_video_list {
			if link.ends_with(fmt) {
				return '\n<video muted autoplay loop controls preload=metadata src="${link}"></video>\n'
			}
		}

		// Discord emote
		if link.contains('/emojis/') {
			return '\n<img class="discord-emoji" loading=lazy alt="" src="${link}">\n'
		}

		return '\n<img loading=lazy alt="" src="${link}">\n'
	})
}

pub fn MediaProcessor.create() !MediaProcessor {
	return MediaProcessor{
		pattern: regex.regex_opt(processor.fmt_media_proc_re)!
	}
}
