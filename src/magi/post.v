module magi

import os
import time
import markdown
import processor

@[heap]
pub struct Post {
	Page
pub mut:
	id   string
	date time.Time

	original_content string
	has_reference    bool
}

pub fn (post &Post) reference(simple bool) string {
	if simple {
		return $tmpl('templates/embed/reference-simple.html')
	}
	return $tmpl('templates/embed/reference.html')
}

pub fn Post.create(path string, mut casper Casper) Post {
	mut post := Post{
		path: path
		id: os.base(path).split_nth('.', 2)[0]
		date: time.unix(os.base(path).split_nth('.', 2)[0].int())
	}

	post.load()

	post.original_content = post.content.clone()
	post.content = markdown.to_html(casper.preprocess(mut post))

	// Resolve shit
	if 'author' !in post.metadata {
		post.metadata['author'] = 'junko'
	}

	if 'thumbnail' in post.metadata {
		post.metadata['filename'] = get_filename(post.metadata['thumbnail'])
		post.metadata['mimetype'] = get_mimetype(post.metadata['filename'])

		if processor.is_video_url(post.metadata['thumbnail']) {
			post.metadata['thumbnail-type'] = 'video'
		}
	}

	// NOTE: discord-post is exported right from discord, as the name suggests..
	//       and i think it'll look cooler if i changed how the post looks.
	if post.metadata['tags'].contains('discord-post') {
		post.metadata['style'] = 'border: .1em solid #5865F2;'
	}

	// Try to fix date, thru metadata, if it exists.
	if post.date.year == 1970 && 'date' in post.metadata {
		post.date = time.unix(post.metadata['date'].int())
	}

	return post
}

// Utils
fn get_filename(url string) string {
	return url.split('/')#[-1..][0].split_nth('?', 2)[0]
}

fn get_mimetype(filename string) string {
	return filename.split_nth('.', 2)#[-1..][0]
}
