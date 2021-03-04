package scrapping

import "strings"

func Clean(value string) string {
	value = strings.TrimSpace(value)
	value = strings.ReplaceAll(value, "\t", "")
	value = strings.ReplaceAll(value, "\n", "")
	value = strings.ReplaceAll(value, "\v", "")
	value = strings.ReplaceAll(value, "\f", "")
	value = strings.ReplaceAll(value, "\r", "")
	value = cleanDoubleSpaces(value)
	return value
}

func cleanDoubleSpaces(value string) string {
	cleaned := value
	for {
		cleaned = strings.ReplaceAll(value, "  ", " ")
		if value == cleaned {
			break
		}
		value = cleaned
	}
	return value
}
