package object

import (
	"regexp"
	"strings"
)

func CleanName(input string) string {
	parsed := regexp.MustCompile("[^a-z0-9-]").ReplaceAllString(strings.ToLower(input), "")
	if len(parsed) > 40 {
		return parsed[:40]
	}
	return parsed
}
