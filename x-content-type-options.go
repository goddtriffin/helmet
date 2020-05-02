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

// Empty returns whether the X-Content-Type-Options is empty.
func (xcto XContentTypeOptions) Empty() bool {
	return xcto.String() == ""
}

// Header adds the X-Content-Type-Options HTTP security header to the given http.ResponseWriter.
func (xcto XContentTypeOptions) Header(w http.ResponseWriter) {
	if !xcto.Empty() {
		w.Header().Set(HeaderXContentTypeOptions, xcto.String())
	}
}
