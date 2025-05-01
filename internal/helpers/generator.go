package helpers

import (
	"crypto/rand"
	"math/big"
	"regexp"
	"strings"
)

const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenerateRandomLetters generates a cryptographically secure random string of letters
// with the specified length using crypto/rand
func GenerateRandomLetters(length int) (string, error) {
	b := make([]byte, length)
	for i := range b {
		// Generate a random index between 0 and len(letterBytes)-1
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
		if err != nil {
			return "", err
		}
		b[i] = letterBytes[n.Int64()]
	}
	return string(b), nil
}

// GenerateSlug converts a title string into a URL-friendly slug with random 5 letters
// Example: "Hello World!" -> "hello-world"
func GenerateSlug(title string, randomNumberLetter int) (string, error) {
	// Convert to lowercase
	slug := strings.ToLower(title)

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove all non-alphanumeric characters except hyphens
	// This will remove symbols like /?;:@&=+$,#
	reg := regexp.MustCompile(`[^a-z0-9\-]`)
	slug = reg.ReplaceAllString(slug, "")

	// Replace multiple hyphens with a single hyphen
	reg = regexp.MustCompile(`[\-]+`)
	slug = reg.ReplaceAllString(slug, "-")

	// Trim hyphens from start and end
	slug = strings.Trim(slug, "-")

	randomLetters, err := GenerateRandomLetters(randomNumberLetter)
	if err != nil {
		return "", err
	}

	return slug + "-" + randomLetters, err
}
