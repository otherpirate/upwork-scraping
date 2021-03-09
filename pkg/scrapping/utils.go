package scrapping

import (
	"strconv"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
)

func cleanString(value string) string {
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

func cleanID(links []soup.Root) string {
	for _, link := range links {
		if cleanString(link.Text()) == "Profile" {
			url := strings.ReplaceAll(link.Attrs()["href"], "?viewMode=1", "")
			return strings.ReplaceAll(url, "https://www.upwork.com/freelancers/", "")
		}
	}
	return ""
}

func formatDateTime(date string) string {
	if date == "" {
		return ""
	}
	months := map[string]int{
		"january":   1,
		"february":  2,
		"march":     3,
		"april":     4,
		"may":       5,
		"june":      6,
		"july":      7,
		"august":    8,
		"september": 9,
		"october":   10,
		"november":  11,
		"december":  12,
	}
	values := strings.Split(date, " ")
	month := months[strings.ToLower(values[0])]
	year, _ := strconv.Atoi(values[1])
	if month == 0 || year == 0 {
		return ""
	}
	dateTime := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	return dateTime.Format("2006-01-02T15:04:05.999999Z")
}
