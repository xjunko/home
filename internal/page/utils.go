package page

import "strings"

// Utils
func getFilename(url string) string {
	parts := strings.Split(url, "/")
	filename := parts[len(parts)-1]

	if idx := strings.Index(filename, "?"); idx != -1 {
		filename = filename[:idx]
	}

	return filename
}

func getMimeType(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}
