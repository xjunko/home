module processor

import pcre

pub const ct_green_re = r'^(>.+?)(?:\n|$)'
pub const ct_red_re = r'^(<<.+?)(?:\n|$)'

pub struct ChanStyleTextProcessor {
}

pub fn (mut chan_text ChanStyleTextProcessor) process(text string) string {
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
	gt_re := pcre.new_regex(processor.ct_green_re, 0) or { return text }
	rt_re := pcre.new_regex(processor.ct_red_re, 0) or { return text }

	mut results := []string{}

	for line in text.split_into_lines() {
		if matching_green_text := gt_re.match_str(line, 0, 0) {
			if green_text := matching_green_text.get(1) {
				if green_text.starts_with('>>') {
					results << line
					continue
				}

				results << '<a style="color: var(--green-text)">\\${green_text}</a><br/>'
				continue
			}
		}

		if matching_red_text := rt_re.match_str(line, 0, 0) {
			if red_text := matching_red_text.get(1) {
				results << '<a style="color: var(--red-text)">\\${red_text[1..]}</a><br/>'
				continue
			}
		}

		results << line
	}

	defer {
		gt_re.free()
		rt_re.free()
	}

	return results.join('\n')
}

pub fn ChanStyleTextProcessor.create() !ChanStyleTextProcessor {
	return ChanStyleTextProcessor{}
}
