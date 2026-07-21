package helpers

import "strings"

func FormatLocation(location string) string {

	// Split the location into parts using '-'.
	parts := strings.Split(location, "-")

	// Process each part separately.
	for i, part := range parts {

		// Replace underscores with spaces.
		part = strings.ReplaceAll(part, "_", " ")

		// Split into individual words.
		words := strings.Fields(part)

		// Capitalize every word.
		for j, word := range words {
			words[j] = CapitalizeWord(word)
		}

		// Join the words back together.
		parts[i] = strings.Join(words, " ")
	}

	// Join the location and country using a comma.
	return strings.Join(parts, ", ")
}

func CapitalizeWord(word string) string {

	if word == "" {
		return ""
	}

	// Common abbreviations that should remain uppercase.
	abbreviations := map[string]string{
		"usa": "USA",
		"uk":  "UK",
		"uae": "UAE",
	}

	lower := strings.ToLower(word)

	if value, ok := abbreviations[lower]; ok {
		return value
	}

	return strings.ToUpper(lower[:1]) + lower[1:]
}
