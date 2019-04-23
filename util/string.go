package util

import (
	"strings"
)

// SpaceDelimitedStringNotGreater ...
func SpaceDelimitedStringNotGreater(first, second string) bool {
	if first == "" {
		return true
	}
	secondParts := strings.Split(second, " ")
	for _, firstPart := range strings.Split(first, " ") {
		if !StringInSlice(firstPart, secondParts) {
			return false
		}
	}
	return true
}

// StringInSlice ...
func StringInSlice(s string, list []string) bool {
	for _, b := range list {
		if s == b {
			return true
		}
	}
	return false
}
