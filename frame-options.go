package helmet

import "net/http"

// HeaderFrameOptions is the X-Frame-Options HTTP security header.
const HeaderFrameOptions = "X-Frame-Options"

// X-Frame-Options options.
const (
	FrameOptionsDeny       FrameOptions = "DENY"
	FrameOptionsSameOrigin FrameOptions = "SAMEORIGIN"
)

// FrameOptions represents the X-Frame-Options HTTP security header.
type FrameOptions string

func (fo FrameOptions) String() string {
	return string(fo)
}

// Exists returns whether the X-Frame-Options has been set.
func (fo FrameOptions) Exists() bool {
	if fo.String() == "" {
		return false
	}

	return true
}

// Header adds the X-Frame-Options HTTP header to the given http.ResponseWriter.
func (fo FrameOptions) Header(w http.ResponseWriter) {
	if fo.Exists() {
		w.Header().Set(HeaderFrameOptions, fo.String())
	}
}
