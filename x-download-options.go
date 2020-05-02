package helmet

import "net/http"

// HeaderXDownloadOptions is the X-Download-Options HTTP header.
const HeaderXDownloadOptions = "X-Download-Options"

// XDownloadOptionsNoOpen represents the X-Download-Options No Open option.
const XDownloadOptionsNoOpen XDownloadOptions = "noopen"

// XDownloadOptions represents the X-Download-Options HTTP security header.
type XDownloadOptions string

func (xdo XDownloadOptions) String() string {
	return string(xdo)
}

// Empty returns whether the X-Download-Options is empty.
func (xdo XDownloadOptions) Empty() bool {
	return xdo.String() == ""
}

// Header adds the X-Download-Options HTTP security header to the given http.ResponseWriter.
func (xdo XDownloadOptions) Header(w http.ResponseWriter) {
	if !xdo.Empty() {
		w.Header().Set(HeaderXDownloadOptions, xdo.String())
	}
}
