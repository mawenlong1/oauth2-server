package util

import (
	"strings"
)

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

func StringInSlice(s string, list []string) bool {
	for _, b := range list {
		if a == b {
			return true
		}
	}
	return false
}
