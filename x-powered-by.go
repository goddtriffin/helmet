package helmet

import "net/http"

// HeaderXPoweredBy is the X-Powered-By HTTP security header.
const HeaderXPoweredBy = "X-Powered-By"

// XPoweredBy represents the X-Powered-By HTTP security header.
type XPoweredBy struct {
	Hide        bool
	Replacement string
}

// NewXPoweredBy creates a new XPoweredBy.
func NewXPoweredBy(hide bool, replacement string) *XPoweredBy {
	return &XPoweredBy{
		Hide:        hide,
		Replacement: replacement,
	}
}

// EmptyXPoweredBy creates a blank slate XPoweredBy.
func EmptyXPoweredBy() *XPoweredBy {
	return NewXPoweredBy(false, "")
}

// Empty returns whether the X-Powered-By is empty.
func (xpb XPoweredBy) Empty() bool {
	return !xpb.Hide && xpb.Replacement == ""
}

// Header adds the X-Powered-By HTTP security header to the given http.ResponseWriter.
func (xpb XPoweredBy) Header(w http.ResponseWriter) {
	if xpb.Empty() {
		return
	}

	if xpb.Hide {
		w.Header().Del(HeaderXPoweredBy)
		return
	}

	if xpb.Replacement != "" {
		w.Header().Set(HeaderXPoweredBy, xpb.Replacement)
	}
}
