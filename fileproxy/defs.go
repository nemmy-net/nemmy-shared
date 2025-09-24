package fileproxy

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	HOOK_POST_UPLOAD = "post-upload"
	HOOK_PRE_UPLOAD  = "pre-upload"

	HEADER_SIGNATURE = "Fdb-Signature"
)

var ErrBadSignature = errors.New(HEADER_SIGNATURE + " header is invalid")

type HookRequest struct {
	Type     string          // One of the `HOOK_` consts
	TokenId  string          // The token used
	UserData json.RawMessage // The same string passed to TokenRequest
}
type HookResponse struct {
	// Override the HTTP response if true, otherwise all Http fields are ignored.
	// Mostly used when rejecting uploads. Only the last-ran hook can decide the response.
	// (An override from `pre-upload` is always ignored if `post-upload` runs after)
	HttpOverride bool
	HttpStatus   int
	HttpHeaders  http.Header
	HttpBody     string
	// Cancel and delete the upload.
	// If this upload is overwriting a file, then that file is deleted.
	RejectUpload bool
}

type TokenRequest struct {
	Key                string
	ContentLength      int64
	ContentType        string
	ContentDisposition string
	Expire             int64 // Milliseconds since unix epoch
	UserData           json.RawMessage
}
type TokenResponse struct {
	Token string
}

type DeleteRequest struct {
	Key string
}
type DeleteResponse = DeleteRequest
type ImageInfoRequest = DeleteRequest

type ImageInfoResponse struct {
	Width  int64
	Height int64
}

type CopyRequest struct {
	SrcKey string
	DstKey string
}
