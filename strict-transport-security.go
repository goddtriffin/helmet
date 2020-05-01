package helmet

// HeaderStrictTransportSecurity is the Strict-Transport-Security HTTP security header.
const HeaderStrictTransportSecurity = "Strict-Transport-Security"

// StrictTransportSecurity represents the Strict-Transport-Security HTTP security header.
type StrictTransportSecurity struct {
	// The time, in seconds, that the browser should remember that a site is only to be accessed using HTTPS.
	MaxAge int

	// If this optional parameter is specified, this rule applies to all of the site's subdomains as well.
	IncludeSubDomains bool

	// After successfully submitting your domain to Google maintained HSTS preload service, browsers will never connect to your domain using an insecure connection.
	Preload bool

	cache string
}

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
