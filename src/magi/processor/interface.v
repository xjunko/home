module processor

pub interface IPost {
	reference(bool) string
mut:
	id            string
	content       string
	has_reference bool
}
