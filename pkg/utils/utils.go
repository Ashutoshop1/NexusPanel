package utils

// Package utils provides utility functions

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// GenerateRandomString generates a random string of specified length
func GenerateRandomString(length int) (string, error) {
	// Generate enough bytes to produce the required length after base64 encoding
	// Base64 encodes 3 bytes to 4 characters, so we need (length * 3 / 4) + 1 bytes
	bytesNeeded := (length*3)/4 + 1
	bytes := make([]byte, bytesNeeded)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	encoded := base64.URLEncoding.EncodeToString(bytes)
	if len(encoded) < length {
		return "", fmt.Errorf("failed to generate string of length %d", length)
	}
	return encoded[:length], nil
}

// TODO: Add more utility functions
// - String manipulation
// - Time formatting
// - File operations
// - Network utilities
