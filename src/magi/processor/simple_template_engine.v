module processor

import os

// dumb fucking template engine cuz im dumb
pub struct SimpleTemplateProcessor {
pub mut:
	cache map[string]string
}

pub fn (mut template SimpleTemplateProcessor) process(text string) string {
	mut result := []string{}

	for line in text.split_into_lines() {
		if line.trim_space().starts_with('@include') || line.contains('@include') {
			template_path := @VMODROOT + '/src/magi/templates/' +
				line.split_nth('@include', 2)[1].replace('"', '').trim_space()

			if os.exists(template_path) {
				println('[Template] File include: ${template_path} | from ${line.trim_space()}')

				result << os.read_file(template_path) or { panic(err) }
			} else {
				println('[Template] Error: Failed to include: ${template_path}')
				result << '<a> [SYSTEM: Failed to include template: ${template_path}] </a> <br>'
			}

			continue
		}

		result << line
	}

	return result.join('\n')
}

pub fn SimpleTemplateProcessor.create() !SimpleTemplateProcessor {
	return SimpleTemplateProcessor{}
}
