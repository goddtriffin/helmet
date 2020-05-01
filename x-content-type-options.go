package helmet

import "net/http"

// HeaderXContentTypeOptions is the X-Content-Type-Options HTTP header.
const HeaderXContentTypeOptions = "X-Content-Type-Options"

// XContentTypeOptionsNoSniff represents the X-Content-Type-Options No Sniff option.
const XContentTypeOptionsNoSniff XContentTypeOptions = "nosniff"

// XContentTypeOptions represents the X-Content-Type-Options HTTP security header.
type XContentTypeOptions string

func (xcto XContentTypeOptions) String() string {
	return string(xcto)
}

// Exists returns whether the X-Content-Type-Options has been set.
func (xcto XContentTypeOptions) Exists() bool {
	if xcto.String() == "" {
		return false
	}

	return true
}

// Header adds the X-Content-Type-Options HTTP security header to the given http.ResponseWriter.
func (xcto XContentTypeOptions) Header(w http.ResponseWriter) {
	if xcto.Exists() {
		w.Header().Set(HeaderXContentTypeOptions, xcto.String())
	}
}
