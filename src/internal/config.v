module internal

import v.vmod
import v.util.version
import json
import os

// The good old v warning hack
pub const used = 0

type ConfigurationValue = bool | f32 | string

pub struct Configuration {
mut:
	path string
pub mut:
	options map[string]ConfigurationValue

	manifest         vmod.Manifest
	compiler_version string = version.full_v_version(false)
}

// Config operation
pub fn (mut config Configuration) expects(key string, value ConfigurationValue) {
	if key !in config.options {
		config.options[key] = value
	}
}

pub fn (config &Configuration) get(key string) ConfigurationValue {
	return config.options[key] or { panic('Config: Key not found => ' + key) }
}

// File Operation
pub fn (mut config Configuration) save() {
	os.write_file(config.path, json.encode_pretty(config)) or { panic(err) }
}

pub fn (mut config Configuration) load() {
	if content := os.read_file(config.path) {
		mut decoded_config := json.decode(Configuration, content) or { panic(err) }

		config.path = decoded_config.path

		reference_option := config.options.clone()
		config.options = decoded_config.options.move()

		// Ensures new config value is set, if doesn't exists.
		for key, value in reference_option {
			if key !in config.options {
				config.options[key] = value
			}
		}
	}
}

pub fn Configuration.create(path string) Configuration {
	mut configuration := Configuration{
		path: path
	}

	println('[Internal] Finding v.mod file, please wait!')
	print('[Internal] Found v.mod file,')
	configuration.manifest = vmod.decode(@VMOD_FILE) or { panic('Failed to read manifest: ${err}') }
	println(' running version ${configuration.manifest.version}!')

	return configuration
}
