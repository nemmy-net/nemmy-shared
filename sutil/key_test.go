package sutil

import (
	"encoding/base64"
	"testing"
)

func TestKeyStuff(t *testing.T) {
	key := MakeKey()

	input := []byte("blehhh")
	hash, err := SignedHash(key, input, nil)
	if err != nil {
		t.Fatal(err)
	}

	ok, err := VerifySignature(key, input, hash)
	if err != nil {
		t.Fatal(err)
	} else if !ok {
		t.Fatalf("Failed to verify signature")
	}

	hashStr := base64.StdEncoding.EncodeToString(hash)
	ok, err = VerifySignatureString(key, input, hashStr)
	if err != nil {
		t.Fatal(err)
	} else if !ok {
		t.Fatalf("Failed to verify signature")
	}
}
