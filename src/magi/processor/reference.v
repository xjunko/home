module processor

import regex
import pcre

pub const ref_reference_re = r'^(>>.+?)(?:\n|$)'
pub const ref_placeholder_prefix = '[[REF_DATA'
pub const ref_placeholder_re = r'\[\[REF_DATA*.*]]'

// Two pass
// First pass, replace with just dumb placeholder.
// Second pass, insert the actual reference.
pub struct ReferenceProcessor {
pub mut:
	pattern regex.RE
}

pub fn (mut reference ReferenceProcessor) first_pass(text string) []string {
	mut do_we_really_need_to_check := false

	for line in text.split_into_lines() {
		if line.trim_space().starts_with('>>') {
			do_we_really_need_to_check = true
			break
		}
	}

	if !do_we_really_need_to_check {
		return [text, do_we_really_need_to_check.str()]
	}

	// NOTE: This is some heavy shit.
	current_re := pcre.new_regex(processor.ref_reference_re, 0) or {
		return [text, do_we_really_need_to_check.str()]
	}

	mut results := []string{}

	// NEWLINE pass
	for line in text.split_into_lines() {
		matching := current_re.match_str(line, 0, 0) or {
			results << line
			continue
		}

		if reference_text := matching.get(1) {
			results << '${processor.ref_placeholder_prefix}:${reference_text.split('>>')[1].trim_space()}]]'
		} else {
			results << line
		}
	}

	defer {
		current_re.free()
	}

	return [results.join('\n'), do_we_really_need_to_check.str()]
}

pub fn (mut reference ReferenceProcessor) final_pass(text string, mut posts []IPost) string {
	return reference.pattern.replace_by_fn(text, fn [posts] (_ regex.RE, text string, b1 int, b2 int) string {
		mut ref_id := text[b1..b2].split_nth(':', 2)[1].split(']')[0]
		mut is_simple_ref := false

		if ref_id.starts_with('>') || ref_id.starts_with('&gt;') {
			is_simple_ref = true

			if ref_id.starts_with('&gt;') {
				ref_id = ref_id[4..]
			} else {
				ref_id = ref_id[1..]
			}
		}

		for post in posts {
			if ref_id == post.id {
				return post.reference(is_simple_ref)
			}
		}

		return '[Reference to #${ref_id}]'
	})
}

pub fn ReferenceProcessor.create() !ReferenceProcessor {
	return ReferenceProcessor{
		pattern: regex.regex_opt(processor.ref_placeholder_re)!
	}
}
