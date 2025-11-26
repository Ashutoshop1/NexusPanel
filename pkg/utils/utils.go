package utils

// Package utils provides utility functions

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomString generates a random string of specified length
func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

// TODO: Add more utility functions
// - String manipulation
// - Time formatting
// - File operations
// - Network utilities
