package asciiart

import "strings"

func InputFormatter(rawInput string) string {
	formattedText := strings.ReplaceAll(rawInput, "\\n", "\n")      // Line Feed
	formattedText = strings.ReplaceAll(formattedText, "\r\n", "\n") // Carriage Return
	formattedText = strings.ReplaceAll(formattedText, "\"", string('"'))
	formattedText = strings.ReplaceAll(formattedText, "\\'", "'")
	return formattedText
}
