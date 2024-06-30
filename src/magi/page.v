module magi

import os
import markdown
import processor

pub const c_special = '@'
pub const c_supported = [
	// Web info (for embeds)
	'title',
	'description',
	'thumbnail',
	// Common
	'tags',
	'outer',
	'author',
	'priority',
	'route',
	// Channel & Blog
	'style',
	'outline',
	'outline-style',
]

pub struct Page {
mut:
	path string @[required]
pub mut:
	metadata   map[string]string
	number     int = -5
	max_number int = -5
	content    string
	outer      string
}

pub fn (mut page Page) load() {
	for line in os.read_lines(page.path) or { [] } {
		if line.starts_with(magi.c_special) {
			for supported in magi.c_supported {
				if line.to_lower().starts_with('${magi.c_special}${supported}') {
					items := line.split_nth('=', 2)
					page.metadata[items[0].replace(magi.c_special, '')] = items[1]

					if supported == 'outline' {
						page.metadata['style'] += ';border: ' + (page.metadata['outline-style'] or {
							'solid'
						}) + ' 2px ${page.metadata['outline']}'
					}
				}
			}
		} else {
			if page.metadata['outer'] == 'start' {
				page.outer += line + '\n'
			} else {
				page.content += line + '\n'
			}
		}
	}
}

pub fn Page.create(path string) Page {
	mut page := Page{
		path: path
	}

	mut template_processor := processor.SimpleTemplateProcessor.create() or { panic(err) }

	page.load()
	page.content = markdown.to_html(template_processor.process(page.content))

	return page
}
