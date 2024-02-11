module processor

import pcre

pub const gt_green_re = r'^(>.+?)(?:\n|$)'

pub struct GreentextProcessor {
}

pub fn (mut greentext GreentextProcessor) process(text string) string {
	mut do_we_really_need_to_check := false

	for line in text.split_into_lines() {
		if line.trim_space().starts_with('>') && !line.trim_space().starts_with('>>') {
			do_we_really_need_to_check = true
			break
		}
	}

	if !do_we_really_need_to_check {
		return text
	}

	// NOTE: This is some heavy shit.
	regex := pcre.new_regex(processor.gt_green_re, 0) or { return text }

	mut results := []string{}

	for line in text.split_into_lines() {
		matching := regex.match_str(line, 0, 0) or {
			results << line
			continue
		}

		if green_text := matching.get(1) {
			if green_text.starts_with('>>') {
				results << line
				continue
			}

			results << '<a style="color: var(--green-text)">\\${green_text}</a><br/>'
		} else {
			results << line
		}
	}

	defer {
		regex.free()
	}

	return results.join('\n')
}

pub fn GreentextProcessor.create() !GreentextProcessor {
	return GreentextProcessor{}
}
