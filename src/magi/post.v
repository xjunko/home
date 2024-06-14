module magi

import os
import time
import markdown

@[heap]
pub struct Post {
	Page
pub mut:
	id   string
	date time.Time

	original_content string
	has_reference    bool
}

pub fn (post &Post) reference() string {
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
	}

	// NOTE: discord-post is exported right from discord, as the name suggests..
	//       and i think it'll look cooler if i changed how the post looks.
	if post.metadata['tags'].contains('discord-post') {
		post.metadata['style'] = 'border: .1em solid #5865F2;'
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
