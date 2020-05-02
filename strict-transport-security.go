package helmet

import (
	"fmt"
	"net/http"
	"strings"
)

// HeaderStrictTransportSecurity is the Strict-Transport-Security HTTP security header.
const HeaderStrictTransportSecurity = "Strict-Transport-Security"

// List of all Strict-Transport-Security directives.
const (
	DirectiveIncludeSubDomains HSTSDirective = "includeSubDomains"
	DirectivePreload           HSTSDirective = "preload"
)

// HSTSDirectiveMaxAge is the Strict-Transport-Security MaxAge directive.
func HSTSDirectiveMaxAge(maxAge int) HSTSDirective {
	if maxAge <= 0 {
		return ""
	}
	return HSTSDirective(fmt.Sprintf("max-age=%d", maxAge))
}

type (
	// HSTSDirective represents a Strict-Transport-Security directive.
	HSTSDirective string

	// StrictTransportSecurity represents the Strict-Transport-Security HTTP security header.
	StrictTransportSecurity struct {
		// The time, in seconds, that the browser should remember that a site is only to be accessed using HTTPS.
		MaxAge int

		// If this optional parameter is specified, this rule applies to all of the site's subdomains as well.
		IncludeSubDomains bool

		// After successfully submitting your domain to Google maintained HSTS preload service, browsers will never connect to your domain using an insecure connection.
		Preload bool

		cache string
	}
)

// NewStrictTransportSecurity creates a new Strict-Transport-Security.
func NewStrictTransportSecurity(maxAge int, includeSubDomains bool, preload bool) *StrictTransportSecurity {
	return &StrictTransportSecurity{
		MaxAge:            maxAge,
		IncludeSubDomains: includeSubDomains,
		Preload:           preload,
	}
}

// EmptyStrictTransportSecurity creates a blank slate Strict-Transport-Security.
func EmptyStrictTransportSecurity() *StrictTransportSecurity {
	return NewStrictTransportSecurity(0, false, false)
}

func (hsts *StrictTransportSecurity) String() string {
	if len(hsts.cache) != 0 {
		return hsts.cache
	}

	// max age is not optional
	if hsts.MaxAge <= 0 {
		hsts.cache = ""
		return hsts.cache
	}

	builder := []string{
		string(HSTSDirectiveMaxAge(hsts.MaxAge)),
	}

	if hsts.IncludeSubDomains {
		builder = append(builder, string(DirectiveIncludeSubDomains))
	}

	if hsts.Preload {
		builder = append(builder, string(DirectivePreload))
	}

	hsts.cache = strings.Join(builder, "; ")
	return hsts.cache
}

// Empty returns whether the Strict-Transport-Security is empty.
func (hsts *StrictTransportSecurity) Empty() bool {
	// includeSubDomains and preload are optional
	return hsts.MaxAge == 0
}

// Header adds the Strict-Transport-Security HTTP security header to the given http.ResponseWriter.
func (hsts *StrictTransportSecurity) Header(w http.ResponseWriter) {
	if !hsts.Empty() {
		w.Header().Set(HeaderStrictTransportSecurity, hsts.String())
	}
}
