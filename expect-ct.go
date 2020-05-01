package helmet

import (
	"fmt"
	"net/http"
	"strings"
)

// HeaderExpectCT is the Expect-CT HTTP security header.
const HeaderExpectCT = "Expect-CT"

// DirectiveEnforce is the Expect-CT Enforce directive.
const DirectiveEnforce ExpectCTDirective = "enforce"

// ExpectCTDirectiveMaxAge is the Expect-CT MaxAge directive.
func ExpectCTDirectiveMaxAge(maxAge int) ExpectCTDirective {
	if maxAge <= 0 {
		return ""
	}
	return ExpectCTDirective(fmt.Sprintf("max-age=%d", maxAge))
}

// ExpectCTDirectiveReportURI is the Expect-CT ReportURI directive.
func ExpectCTDirectiveReportURI(reportURI string) ExpectCTDirective {
	if reportURI == "" {
		return ""
	}
	return ExpectCTDirective(fmt.Sprintf(`report-uri="%s"`, reportURI))
}

type (
	// ExpectCTDirective represents a Expect-CT directive.
	ExpectCTDirective string

	// ExpectCT represents the Expect-CT HTTP security header.
	ExpectCT struct {
		MaxAge    int    // number of seconds that the browser should cache and apply the received policy for
		Enforce   bool   // controls whether the browser should enforce the policy or treat it as report-only mode
		ReportURI string // specifies where the browser should send reports if it does not receive valid CT information

		cache string
	}
)

// NewExpectCT creates a new Expect-CT.
func NewExpectCT(maxAge int, enforce bool, reportURI string) *ExpectCT {
	return &ExpectCT{
		Enforce:   enforce,
		MaxAge:    maxAge,
		ReportURI: reportURI,
	}
}

// EmptyExpectCT creates a blank slate Expect-CT.
func EmptyExpectCT() *ExpectCT {
	return NewExpectCT(0, false, "")
}

func (ect *ExpectCT) String() string {
	if len(ect.cache) != 0 {
		return ect.cache
	}

	// max age is not optional
	if ect.MaxAge <= 0 {
		ect.cache = ""
		return ect.cache
	}

	builder := []string{
		string(ExpectCTDirectiveMaxAge(ect.MaxAge)),
	}

	if ect.Enforce {
		builder = append(builder, string(DirectiveEnforce))
	}

	if ect.ReportURI != "" {
		builder = append(builder, string(ExpectCTDirectiveReportURI(ect.ReportURI)))
	}

	ect.cache = strings.Join(builder, ", ")
	return ect.cache
}

// Exists returns whether the Expect-CT has been set.
func (ect *ExpectCT) Exists() bool {
	if ect.MaxAge == 0 {
		// enforce and report-uri are optional
		return false
	}

	return true
}

// Header adds the Expect-CT HTTP security header to the given http.ResponseWriter.
func (ect *ExpectCT) Header(w http.ResponseWriter) {
	if ect.Exists() {
		w.Header().Set(HeaderExpectCT, ect.String())
	}
}
