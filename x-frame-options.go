package helmet

import "net/http"

// HeaderXFrameOptions is the X-Frame-Options HTTP security header.
const HeaderXFrameOptions = "X-Frame-Options"

// X-Frame-Options options.
const (
	XFrameOptionsDeny       XFrameOptions = "DENY"
	XFrameOptionsSameOrigin XFrameOptions = "SAMEORIGIN"
)

// XFrameOptions represents the X-Frame-Options HTTP security header.
type XFrameOptions string

func (xfo XFrameOptions) String() string {
	return string(xfo)
}

// Exists returns whether the X-Frame-Options has been set.
func (xfo XFrameOptions) Exists() bool {
	if xfo.String() == "" {
		return false
	}

	return true
}

// Header adds the X-Frame-Options HTTP header to the given http.ResponseWriter.
func (xfo XFrameOptions) Header(w http.ResponseWriter) {
	if xfo.Exists() {
		w.Header().Set(HeaderXFrameOptions, xfo.String())
	}
}
