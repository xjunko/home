// Copyright (c) 2023 l-m.dev. All rights reserved.
// Use of this source code is governed by an AGPL license
// that can be found in the LICENSE file.
//
// Modified by [xJunko | konno.ovh] on 2024
// Changes made: integrate database cache directly into the processor

module processor

import json
import regex
import net.http
import encoding.base64
import db.sqlite

pub struct RootCoverArt {
pub:
	url    string
	width  int
	height int
}

pub struct RootArtist {
pub:
	id      string
	profile struct {
	pub:
		name string
	}
}

pub struct RootTrack {
pub:
	id             string
	name           string
	uri            string
	album_of_track struct {
		cover_art struct {
			sources []RootCoverArt
		} @[json: coverArt]
	} @[json: albumOfTrack]

	previews struct {
		audio_previews struct {
		pub:
			items []struct {
			pub:
				url string
			}
		} @[json: audioPreviews]
	}

	first_artist struct {
		items []RootArtist
	} @[json: firstArtist]
}

pub struct Root {
pub:
	entities struct {
		items map[string]RootTrack
	}
}

@[table: 'spotify']
pub struct Track {
pub:
	source_url        string @[primary; required]
	id                string @[required; sql_type: 'TEXT']
	name              string @[required; sql_type: 'TEXT']
	artist            string @[required; sql_type: 'TEXT']
	artist_id         string @[required; sql_type: 'TEXT']
	cover_art_url     string @[required; sql_type: 'TEXT']
	audio_preview_url string @[default: ''; required; sql_type: 'TEXT']
}

const sptfy_url_re = r'https?://open\.spotify\.com/track/(\w+)'
const sptfy_script_re = r'<script\s+id="initial-state"\s+type="text/plain">([^<]+)</script>'
const sptfy_i_fucked_up = '[Spotify]: Failed to embed url. Reason: '

pub struct SpotifyProcessor {
pub mut:
	pattern        regex.RE
	script_pattern regex.RE
}

pub fn (spotify SpotifyProcessor) largest_cover_art(sources []RootCoverArt) ?string {
	mut root_art := sources[0] or { return none }

	for i := 1; i < sources.len; i++ {
		size := sources[i].width * sources[i].height
		root_size := root_art.width * root_art.height

		if size > root_size {
			root_art = sources[i]
		}
	}

	return root_art.url
}

pub fn (mut spotify SpotifyProcessor) get_track_from_url(url string) !Track {
	response := http.get(url) or {
		return error(processor.sptfy_i_fucked_up + 'Failed to reach url.')
	}

	script_index, _ := spotify.script_pattern.find_from(response.body, 0)

	if script_index < 0 {
		return error(processor.sptfy_i_fucked_up + 'Embedded script not found.')
	}

	base64_data := spotify.script_pattern.get_group_by_id(response.body, 0)
	music_data := base64.decode_str(base64_data)

	decoded_data := json.decode(Root, music_data) or {
		return error(processor.sptfy_i_fucked_up + 'Embedded data contains no valid data.')
	}

	for _, track in decoded_data.entities.items {
		if audio_preview := track.previews.audio_previews.items[0] {
			if artist := track.first_artist.items[0] {
				thumbnail := spotify.largest_cover_art(track.album_of_track.cover_art.sources) or {
					return error(processor.sptfy_i_fucked_up + 'Failed to get thumbnail.')
				}

				fetched_track := Track{
					source_url: url
					id: track.id
					name: track.name
					artist: artist.profile.name
					artist_id: artist.id
					cover_art_url: thumbnail
					audio_preview_url: audio_preview.url
				}

				return fetched_track
			}
		}
	}

	return error('Failed, unhandled exception!')
}

pub fn (mut spotify SpotifyProcessor) get_track(url string, mut db sqlite.DB) !Track {
	// Fetch from database.
	tracks := sql db {
		select from Track where source_url == url limit 1
	}!

	if db_track := tracks[0] {
		return db_track
	}

	// Fetch from URL.
	if fetched_track := spotify.get_track_from_url(url) {
		cloned_fetched_track := fetched_track

		// HACK: honestly no clue why, but this crashes the V compiler without the explicit copy.
		sql db {
			insert cloned_fetched_track into Track
		}!

		return fetched_track
	}

	return error(processor.sptfy_i_fucked_up + 'Failed to get data from both DB and Online.')
}

pub fn (mut spotify SpotifyProcessor) handle_url(url string, mut db sqlite.DB) string {
	if track := spotify.get_track(url, mut db) {
		return $tmpl('../templates/embed/spotify.html')
	} else {
		return err.str()
	}
}

pub fn (mut spotify SpotifyProcessor) process(text string, mut db sqlite.DB) string {
	return spotify.pattern.replace_by_fn(text, fn [mut spotify, mut db] (_ regex.RE, content string, b1 int, b2 int) string {
		return spotify.handle_url(content[b1..b2], mut db)
	})
}

pub fn SpotifyProcessor.create() !SpotifyProcessor {
	return SpotifyProcessor{
		pattern: regex.regex_opt(processor.sptfy_url_re)!
		script_pattern: regex.regex_opt(processor.sptfy_script_re)!
	}
}
