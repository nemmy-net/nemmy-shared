package fileproxy

import "testing"

func Test_Filename(t *testing.T) {
	if IsValidFileName("a/b") {
		t.Fatal()
	} else if IsValidFileName("dir?recursive=true") {
		t.Fatal()
	} else if SanitizeFileName("a/b") != "a_b" {
		t.Fatal()
	} else if SanitizeFileName("dir?recursive=true") != "dir_recursive_true" {
		t.Fatal()
	} else if SanitizeFileName("Filer-Server-API#filer-server") != "Filer-Server-API_filer-server" {
		t.Fatal()
	} else if !IsValidFileKey("test/key") {
		t.Fatal()
	} else if IsValidFileKey("test//key") {
		t.Fatal()
	} else if IsValidFileKey("/test/key") {
		t.Fatal()
	} else if IsValidFileKey("test/key/") {
		t.Fatal()
	} else if IsValidFileName("") {
		t.Fatal()
	}
}
