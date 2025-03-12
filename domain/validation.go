package domain

import (
	"regexp"
	"strings"
)

// Validates if the blob storage url uses https and a valid image extension.
// Param: blob storage url.
// Returns: true if the url is valid, false otherwise.
func ValidateBlobStorage(url string) bool {
	url = strings.ToLower(url)
	pattern := regexp.MustCompile(`^https:\/\/([a-z0-9\-\.]+)(\/[a-z0-9\-\.\/]*)*\.(jpg|jpeg|png|gif|webp)$`)
	return pattern.MatchString(url)
}
