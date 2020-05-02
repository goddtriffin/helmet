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

// Empty returns whether the X-Frame-Options is empty.
func (xfo XFrameOptions) Empty() bool {
	return xfo.String() == ""
}

// Header adds the X-Frame-Options HTTP header to the given http.ResponseWriter.
func (xfo XFrameOptions) Header(w http.ResponseWriter) {
	if !xfo.Empty() {
		w.Header().Set(HeaderXFrameOptions, xfo.String())
	}
}
