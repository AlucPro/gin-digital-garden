package utils

import "strings"

func GenerateSlug(title string) string {
	return strings.ReplaceAll(strings.ToLower(title), " ", "-")
}
