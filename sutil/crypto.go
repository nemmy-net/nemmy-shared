package sutil

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"os"

	"golang.org/x/crypto/sha3"
)

const _KEY_LEN = 32
const _SIGN_HASH_LEN = 256 / 8

type Key []byte

func MakeKey() Key {
	key := make([]byte, _KEY_LEN)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatal(err)
	}
	return key
}

// Load key bytes from file.
// A new file is created if `makeKey` is true and the file does not exist.
// Panics on failure.
func MakeOrLoadKey(filename string, makeKey bool) []byte {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if file != nil {
		defer file.Close()
	}

	if errors.Is(err, os.ErrNotExist) { // Write random key to disk
		if !makeKey {
			log.Panicf("Missing key file: %v", filename)
		}

		key := MakeKey()
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0)
		if file != nil {
			defer file.Close()
		}

		if err != nil {
			log.Panicf("Error creating key file: %v", err)
		}
		_, err = file.Write(key)
		if err != nil {
			log.Panicf("Error writing key: %v", err)
		}
		return key
	} else if err != nil {
		log.Panicf("Error opening key: %v", err)
	}

	key := make([]byte, _KEY_LEN)
	n, err := file.Read(key)
	if err != nil {
		log.Panic(err)
	} else if n != len(key) {
		log.Panicf("Expected session key length of %v, got %v", len(key), n)
	}
	return key
}

// Generate a signed hash.
// If output is nil then a new array is returned.
func SignedHash(key Key, input []byte, output []byte) (hash []byte, err error) {
	hasher := sha3.NewShake256()
	if output == nil {
		output = make([]byte, _SIGN_HASH_LEN)
	}
	if _, err := hasher.Write(input); err != nil {
		return nil, err
	} else if _, err := hasher.Write(key); err != nil {
		return nil, err
	} else if _, err := hasher.Read(output); err != nil {
		return nil, err
	}
	return output, nil
}

func VerifySignature(key Key, input []byte, givenHash []byte) (ok bool, err error) {
	correctHash, err := SignedHash(key, input, nil)
	if err != nil {
		return false, err
	}
	return bytes.Equal(givenHash, correctHash), nil
}

// Like VerifySignature except givenHash is encoded in std base64
func VerifySignatureString(key Key, input []byte, givenHash string) (ok bool, err error) {
	decoded, err := base64.StdEncoding.DecodeString(givenHash)
	if err != nil {
		return false, err
	}
	return VerifySignature(key, input, decoded)
}

// `input` must include a signed hash at the end.
// Returns the input without the trailing hash
func VerifySignatureTail(key Key, input []byte) (data []byte, ok bool, err error) {
	if len(input) < _SIGN_HASH_LEN {
		return nil, false, nil
	}

	givenHash := input[len(input)-_SIGN_HASH_LEN:]
	data = input[:len(input)-_SIGN_HASH_LEN]

	ok, err = VerifySignature(key, data, givenHash)
	return data, ok, err
}
