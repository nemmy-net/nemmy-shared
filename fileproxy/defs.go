package fileproxy

const (
	HOOK_POST_UPLOAD = "post-upload"
)

type HookRequest struct {
	Type     string // One of the `HOOK_` consts
	TokenId  string // The token used
	UserData string // The same string passed to TokenRequest
}
type HookResponse struct {
	HttpOverride bool // Override the HTTP response if true, otherwise all Http fields are ignored
	HttpStatus   int
	HttpHeaders  map[string][]string
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
	Expire             int64
	Userdata           string
}
type TokenResponse struct {
	Token string
}

type ImageInfo struct {
	Width  int64
	Height int64
}
