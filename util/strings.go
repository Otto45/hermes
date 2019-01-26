package util

// SliceContainsString searches a string slice for a given string
func SliceContainsString(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// BuildHTTPResponseBodyForSuccess returns a JSON response body as a string containing the message given as an argument
func BuildHTTPResponseBodyForSuccess(msg string) string {
	return "{\"success\":\"" + msg + "\"}"
}

// BuildHTTPResponseBodyForError returns a JSON response body as a string containing the message given as an argument
func BuildHTTPResponseBodyForError(msg string) string {
	return "{\"error\":\"" + msg + "\"}"
}
