package fileproxy

import (
	"mime"
	"testing"
)

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

func Test_NovlVersion(t *testing.T) {
	sample := "application/x-nemmy-novl;version=2.5.6"
	mediaType, params, err := mime.ParseMediaType(sample)
	if err != nil {
		t.Fatal(err)
	}

	info, ok := ParseNovlInfo(mediaType, params)
	if !ok {
		t.Fatal("Failed to parse novl info")
	}
	if info.Major != 2 || info.Minor != 5 || info.Patch != 6 {
		t.Fatalf("Expected 2.5.6, got %v.%v.%v", info.Major, info.Minor, info.Patch)
	}

	sample = "application/x-nemmy-novl;version=1.2.3"
	mediaType, params, err = mime.ParseMediaType(sample)
	if err != nil {
		t.Fatal(err)
	}

	info, ok = ParseNovlInfo(mediaType, params)
	if !ok {
		t.Fatal("Failed to parse novl info")
	}
	if info.Major != 1 || info.Minor != 2 || info.Patch != 3 {
		t.Fatalf("Expected 1.2.3, got %v.%v.%v", info.Major, info.Minor, info.Patch)
	}

	if _, ok := ParseNovlVersion("1"); ok {
		t.Fatal("Expected invalid version")
	} else if _, ok := ParseNovlVersion("1.a"); ok {
		t.Fatal("Expected invalid version")
	}
}
