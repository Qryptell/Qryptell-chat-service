package util

import (
	"strings"
)

// Convert string to array spliting by ','
func StringToArray(value string) (keys []string) {
	if value == "" {
		return
	}
	keys = strings.Split(value,",")
	return keys
}
