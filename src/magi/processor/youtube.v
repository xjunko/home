module processor

import regex
import net.http
import db.sqlite

@[table: 'youtube']
pub struct YoutubeThumbnail {
pub:
	video_id      string @[primary; required]
	thumbnail_url string @[required; sql_type: 'TEXT']
}

pub struct YoutubeProcessor {
pub mut:
	pattern regex.RE
}

pub fn (youtube YoutubeProcessor) get_video_thumbnail_from_id(video_id string) string {
	for quality in ['maxresdefault.jpg', 'mqdefault.jpg', '0.jpg'] {
		current_thumbnail := 'https://i3.ytimg.com/vi/${video_id}/${quality}'

		if (http.get(current_thumbnail) or { panic(err) }).status_code == 200 {
			return current_thumbnail
		}
	}

	return 'https://i3.ytimg.com/vi/${video_id}/0.jpg' // Most shit quality
}

pub fn (youtube YoutubeProcessor) get_video_thumbnail(video_id string, mut db sqlite.DB) string {
	default_thumbnail := 'https://i3.ytimg.com/vi/${video_id}/0.jpg' // Most shit quality

	// Fetch DB
	db_thumbnails := sql db {
		select from YoutubeThumbnail where video_id == video_id limit 1
	} or { return default_thumbnail }

	if thumbnail := db_thumbnails[0] {
		return thumbnail.thumbnail_url
	}

	// Fetch online
	online_thumbnail := YoutubeThumbnail{
		video_id: video_id
		thumbnail_url: youtube.get_video_thumbnail_from_id(video_id)
	}

	sql db {
		insert online_thumbnail into YoutubeThumbnail
	} or {
		println('[Database] Failed to save thumbnail info: ${err}')
		return default_thumbnail
	}

	return default_thumbnail
}

pub fn (mut youtube YoutubeProcessor) process(text string, mut db sqlite.DB) string {
	return youtube.pattern.replace_by_fn(text, fn [youtube, mut db] (re regex.RE, text string, b1 int, b2 int) string {
		video_url := text[b1..b2]
		video_id := re.get_group_by_id(text, 0)

		video_thumbnail := youtube.get_video_thumbnail(video_id, mut db)

		// Last check, just in case it went through.
		if (http.get(video_thumbnail) or { panic(err) }).status_code != 200 {
			return video_url
		}

		return $tmpl('../templates/embed/youtube.html')
	})
}

pub fn YoutubeProcessor.create() !YoutubeProcessor {
	return YoutubeProcessor{
		pattern: regex.regex_opt(r'https?://(?:www\.)?youtu(?:be\.com/watch\?v=)|(?:\.be/)(\S+)')!
	}
}
