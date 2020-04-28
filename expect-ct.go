package helmet

import (
	"fmt"
	"net/http"
	"strings"
)

// HeaderExpectCT is the Expect-CT HTTP header.
const HeaderExpectCT = "Expect-CT"

// ExpectCT is the Expect-CT HTTP security header.
type ExpectCT struct {
	Enforce   bool   // controls whether the browser should enforce the policy or treat it as report-only mode
	MaxAge    int    // number of seconds that the browser should cache and apply the received policy for
	ReportURI string // specifies where the browser should send reports if it does not receive valid CT information

	cache string
}

// NewExpectCT creates a new ExpectCT.
func NewExpectCT(maxAge int, enforce bool, reportURI string) *ExpectCT {
	return &ExpectCT{
		Enforce:   enforce,
		MaxAge:    maxAge,
		ReportURI: reportURI,
	}
}

// EmptyExpectCT creates a blank slate ExpectCT.
func EmptyExpectCT() *ExpectCT {
	return NewExpectCT(0, false, "")
}

func (ect *ExpectCT) String() string {
	if len(ect.cache) != 0 {
		return ect.cache
	}

	// max age is not optional
	if ect.MaxAge == 0 {
		ect.cache = ""
		return ect.cache
	}

	builder := []string{fmt.Sprintf("max-age=%d", ect.MaxAge)}

	if ect.Enforce {
		builder = append(builder, "enforce")
	}

	if ect.ReportURI != "" {
		builder = append(builder, fmt.Sprintf("report-uri=\"%s\"", ect.ReportURI))
	}

	ect.cache = strings.Join(builder, ", ")
	return ect.cache
}

// Exists returns whether the Expect CT has been set.
func (ect *ExpectCT) Exists() bool {
	if ect.MaxAge == 0 {
		// enfore and report-uri are optional
		return false
	}

	return true
}

// AddHeader adds the Expect-CT HTTP header to the given ResponseWriter.
func (ect *ExpectCT) AddHeader(w http.ResponseWriter) {
	if ect.Exists() {
		w.Header().Set(HeaderExpectCT, ect.String())
	}
}
