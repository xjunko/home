module magi

import os
import time
import internal { Configuration }

pub struct Magi {
mut:
	pages_cache map[int][]Post
	casper      Casper
pub mut:
	config Configuration @[required]

	page  []Page
	posts []Post // Similar to Page but boxed into a Post type.
	notes []Post // No need to make a new type for this.
}

pub fn (mut magi Magi) resolve_pages() {
	for file in os.glob('src/magi/templates/pages/*.md') or { [] } {
		println('[Magi] Creating page: ${os.base(file)}')
		magi.page << Page.create(file)
	}

	magi.page.sort(a.metadata['priority'] > b.metadata['priority'])
}

pub fn (mut magi Magi) resolve_channel() {
	mut files := []string{}

	files << os.glob('static/entry/written/*.md') or { [] }

	println('[Channel] Found ${files.len} written entries.')

	if magi.config.get('website.channel.discord') as bool {
		files << os.glob('static/entry/discord/*.md') or { [] }
		println('[Channel] Found ${files.len} discord entries.')
	}

	for i, file in files {
		// Show progress every 10 files
		if i % 10 == 0 || i == files.len - 1 {
			// Progress bar
			println('[Channel] Processing ${i + 1}/${files.len} files.')
		}

		mut new_post := Post.create(file, mut magi.casper)

		// Thumbnail hack: Discord CDN FIX
		if 'thumbnail' in new_post.metadata {
			new_post.metadata['thumbnail'] = magi.casper.discord.process(new_post.metadata['thumbnail'])
		}

		if 'exclude' !in new_post.metadata {
			magi.posts << new_post
		}
	}

	println('[Channel] Sorting posts and running post-process.')
	magi.casper.postprocess(mut magi.posts)
	magi.posts.sort(a.date > b.date)
	println('[Channel] Done.')
}

pub fn (mut magi Magi) resolve_notes() {
	mut files := []string{}

	files << os.glob('static/entry/notes/*.md') or { [] }

	println('[Note] Found ${files.len} written notes.')

	for i, file in files {
		// Show progress every 10 files
		if i % 10 == 0 || i == files.len - 1 {
			// Progress bar
			println('[Note] Processing ${i + 1}/${files.len} files.')
		}

		mut new_post := Post.create(file, mut magi.casper)

		if 'exclude' !in new_post.metadata {
			magi.notes << new_post
		}
	}

	println('[Note] Sorting Notes and running post-process.')
	magi.casper.postprocess(mut magi.notes)
	magi.notes.sort(a.date > b.date)
	println('[Note] Done.')
}

//
pub fn execute(config Configuration) {
	mut magi := Magi{
		config: config
		casper: Casper.create() or { panic(err) }
	}

	println('[Magi] Starting Magi!')

	c_chan_enabled := (config.get('website.channel.enable') as bool) == true

	println('[Magi] Channel state: ${c_chan_enabled}')

	if c_chan_enabled {
		println('[Magi] Resolving channel.')
		magi.resolve_channel()
		println('[Magi] Channel resolved.')
	} else {
		// HACK: To avoid index=0, len=0 problem.
		magi.posts << Post{
			path: 'NONE'
			id: '0'
			date: time.now()
			original_content: ''
		}

		// We don't need these files for now.
		for chan_html_file in os.glob('chan/*.html') or { [] } {
			os.rm(chan_html_file) or {
				panic('Failed to delete unused channel file, this should never happen!!!!')
			}
		}

		os.write_file('chan/index.html', '<meta http-equiv="refresh" content="0;url=https://konno.ovh">') or {}
		os.write_file('chan/1.html', '<meta http-equiv="refresh" content="0;url=https://konno.ovh">') or {}
	}

	println('[Magi] Resolving pages.')
	magi.resolve_pages()
	println('[Magi] Pages resolved: ${magi.page.len}')

	println('[Magi] Resolving notes.')
	magi.resolve_notes()
	println('[Magi] Notes resolved: ${magi.notes.len}')

	// RSS
	println('[Magi] Saving RSS Feed.')
	os.write_file('feed.xml', $tmpl('templates/rss.xml')) or { panic(err) }

	// Page
	println('[Magi] Saving pages!')
	for mut page in magi.page {
		// Channel
		if os.base(page.path) == 'channel.md' {
			if !c_chan_enabled {
				continue
			}

			posts := magi.posts.clone()
			posts_per_page := 20

			for i in 0 .. (posts.len / posts_per_page) + 1 {
				page.number = i
				page.max_number = posts.len / posts_per_page
				magi.posts = posts#[i * posts_per_page..(i * posts_per_page) + posts_per_page]

				if posts#[i * posts_per_page..(i * posts_per_page) + posts_per_page].len == 0 {
					continue
				}

				os.write_file('chan/${i + 1}.html', $tmpl('templates/base.html')) or { panic(err) }
			}

			magi.posts = posts
			continue
		}

		// Note
		if os.base(page.path) == 'notes.md' {
			// Entry page
			os.write_file('${os.base(page.path).split_nth('.', 2)[0]}.html', $tmpl('templates/base.html')) or {
				panic(err)
			}

			page.metadata['step'] = 'notes'

			// Generate notes
			for current_note in magi.notes {
				page.metadata['step-data'] = current_note.metadata['slog']

				os.write_file('notes/${current_note.metadata['slog']}.html', $tmpl('templates/base.html')) or {
					panic(err)
				}
			}

			continue
		}

		if os.base(page.path) == 'index.md' {
			// TODO: This is a hack, we should have a better way to do this.
			mut recent_post_template := '<a style="color: #96c83b;">>none, unfortunately.</a>'
			mut recent_note_template := '<a style="color: #96c83b;">>none, unfortunately.</a>'

			if post := magi.posts[0] {
				recent_post_template = $tmpl('templates/component/mini/post.html')
				recent_post_template += '\n<link rel="stylesheet" type="text/css" href="/static/css/channel.css">'
			}

			if note := magi.notes[0] {
				recent_note_template = $tmpl('templates/component/mini/note.html')
			}

			page.content = page.content.replace('++RECENT_POST', recent_post_template)
			page.content = page.content.replace('++RECENT_NOTE', recent_note_template)

			os.write_file('${os.base(page.path).split_nth('.', 2)[0]}.html', $tmpl('templates/base.html')) or {
				panic(err)
			}

			continue
		}

		os.write_file('${os.base(page.path).split_nth('.', 2)[0]}.html', $tmpl('templates/base.html')) or {
			panic(err)
		}
	}

	// Done
	magi.casper.database.close() or { panic(err) }
}

// Called from templates/component/channel-redirect.html
pub fn (mut magi Magi) get_pages() map[int][]Post {
	if magi.pages_cache.len == 0 {
		// Maybe cache this?
		mut pages := map[int][]Post{}

		posts := magi.posts.clone()
		posts_per_page := 20

		for i in 0 .. (posts.len / posts_per_page) + 1 {
			pages[i + 1] = posts#[i * posts_per_page..(i * posts_per_page) + posts_per_page]
		}

		magi.pages_cache = unsafe { pages }
	}

	return magi.pages_cache
}

pub fn (mut magi Magi) escape_xml(content string) string {
	return content.replace_each(['"', '&quot;', "'", '&apos;', '<', '&lt;', '>', '&gt;', '&', '&amp;'])
}
