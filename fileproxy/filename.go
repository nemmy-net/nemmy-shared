package fileproxy

import "strings"

const _MISC_FILE_CHARS = "!-_.'()"

// Returns a URL-safe S3-compatible ASCII filename without paths.
// Bad bytes are replaced with '_'
func SanitizeFileName(name string) string {
	builder := strings.Builder{}
	builder.Grow(len(name))
	for i := range len(name) {
		b := name[i]
		if !IsValidFileNameByte(b) {
			builder.WriteByte('_')
		} else {
			builder.WriteByte(b)
		}
	}
	return builder.String()
}

// File keys look like paths and may have (non-consecutive) forward slashes.
// A leading or trailing slash is forbidden.
func IsValidFileKey(key string) bool {
	if len(key) == 0 || key[0] == '/' || key[len(key)-1] == '/' {
		return false
	}

	for i := range len(key) {
		b := key[i]
		if b == '/' {
			if i+1 < len(key) && key[i+1] == '/' {
				return false // Cannot have consecutive slashes
			}
		} else if !IsValidFileNameByte(b) {
			return false
		}
	}
	return true
}

func IsValidFileName(name string) bool {
	if name == "" {
		return false
	}

	for i := range len(name) {
		if !IsValidFileNameByte(name[i]) {
			return false
		}
	}
	return true
}

func IsValidFileNameByte(b byte) bool {
	if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') {
		return true
	}
	for i := range len(_MISC_FILE_CHARS) {
		if b == _MISC_FILE_CHARS[i] {
			return true
		}
	}
	return false
}
