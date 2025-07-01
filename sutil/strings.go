package sutil

import (
	"log"
	"strings"
)

// Ensures string begins with `http(s)://` and does not end with a slash
func NormalizeHttpUrl(str string) string {
	if str == "" {
		return str
	}

	if !strings.HasPrefix(str, "http://") && !strings.HasPrefix(str, "https://") {
		log.Panicf("String must begin with the http(s) protocol: %v", str)
	} else if strings.HasSuffix(str, "/") {
		return str[:len(str)-1]
	}
	return str
}
