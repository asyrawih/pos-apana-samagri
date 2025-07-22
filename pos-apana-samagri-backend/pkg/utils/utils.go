package utils

import "github.com/google/uuid"

// GenerateUUID generates a new UUID v4
func GenerateUUID() string {
	return uuid.New().String()
}

// ContainsString checks if a string exists in a slice of strings
func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// RemoveString removes a string from a slice of strings
func RemoveString(slice []string, s string) []string {
	result := []string{}
	for _, item := range slice {
		if item != s {
			result = append(result, item)
		}
	}
	return result
}
