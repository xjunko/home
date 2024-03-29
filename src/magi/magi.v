module magi

import os
import internal { Configuration }

pub struct Magi {
mut:
	pages_cache map[int][]Post
	casper      Casper
pub mut:
	config Configuration @[required]

	page  []Page
	posts []Post // Similar to Page but boxed into a Post type.
}

pub fn (mut magi Magi) resolve_pages() {
	for file in os.glob('src/magi/templates/pages/*.md') or { [] } {
		println('[Magi] Creating page: ${os.base(file)}')
		magi.page << Page.create(file)
	}

	magi.page.sort(a.metadata['priority'] > b.metadata['priority'])
}

pub fn (mut magi Magi) resolve_blog() {
	mut files := []string{}

	files << os.glob('static/entry/written/*.md') or { [] }
	files << os.glob('static/entry/discord/*.md') or { [] }

	for file in files {
		magi.posts << Post.create(file, mut magi.casper)
	}

	magi.casper.postprocess(mut magi.posts)

	magi.posts.sort(a.date > b.date)
}

//
pub fn execute(config Configuration) {
	mut magi := Magi{
		config: config
		casper: Casper.create() or { panic(err) }
	}

	println('[Internal] Magi, created.')

	magi.resolve_blog()
	magi.resolve_pages()

	// RSS
	println('[Magi] Saving RSS Feed.')
	os.write_file('feed.xml', $tmpl('templates/rss.xml')) or { panic(err) }

	// Page
	println('[Magi] Saving pages!')
	for mut page in magi.page {
		if os.base(page.path) == 'blog.md' {
			posts := magi.posts.clone()
			posts_per_page := 20

			for i in 0 .. (posts.len / posts_per_page) + 1 {
				page.number = i
				page.max_number = posts.len / posts_per_page
				magi.posts = posts#[i * posts_per_page..(i * posts_per_page) + posts_per_page]

				if posts#[i * posts_per_page..(i * posts_per_page) + posts_per_page].len == 0 {
					continue
				}

				os.write_file('blog/${i + 1}.html', $tmpl('templates/base.html')) or { panic(err) }
			}

			continue
		}

		os.write_file('${os.base(page.path).split_nth('.', 2)[0]}.html', $tmpl('templates/base.html')) or {
			panic(err)
		}
	}
}

// Called fropm templates/component/blog-redirect.html
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
