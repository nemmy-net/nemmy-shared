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

func IsValidFileName(name string) bool {
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
