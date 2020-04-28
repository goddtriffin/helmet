package helmet

import (
	"net/http"
)

// HeaderPermittedCrossDomainPolicies is the Permitted Cross Domain Policies HTTP header.
const HeaderPermittedCrossDomainPolicies = "X-Permitted-Cross-Domain-Policies"

// Permitted Cross Domain Policies options.
const (
	PermittedCrossDomainPoliciesNone          = "none"
	PermittedCrossDomainPoliciesMasterOnly    = "master-only"
	PermittedCrossDomainPoliciesByContentType = "by-content-type"
	PermittedCrossDomainPoliciesByFTPFilename = "by-ftp-filename"
	PermittedCrossDomainPoliciesAll           = "all"
)

// Helmet is a HTTP security middleware for Go(lang) inspired by HelmetJS for Express.js.
type Helmet struct {
	ContentSecurityPolicy        *ContentSecurityPolicy
	DNSPrefetchControl           DNSPrefetchControl
	ExpectCT                     *ExpectCT
	PermittedCrossDomainPolicies string
}

// New creates a new Helmet.
func New() *Helmet {
	return &Helmet{
		ContentSecurityPolicy: EmptyCSP(),
		ExpectCT:              EmptyExpectCT(),
	}
}

// Default creates a new Helmet with default settings.
func Default() *Helmet {
	return &Helmet{
		ContentSecurityPolicy:        EmptyCSP(),
		DNSPrefetchControl:           DNSPrefetchControlOff,
		ExpectCT:                     EmptyExpectCT(),
		PermittedCrossDomainPolicies: PermittedCrossDomainPoliciesNone,
	}
}

// Secure is the middleware handler.
func (h *Helmet) Secure(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ContentSecurityPolicy.AddHeader(w)
		h.DNSPrefetchControl.AddHeader(w)
		h.ExpectCT.AddHeader(w)

		if len(h.PermittedCrossDomainPolicies) != 0 {
			w.Header().Set(HeaderPermittedCrossDomainPolicies, h.PermittedCrossDomainPolicies)
		}

		// w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		// w.Header().Set("X-Content-Type-Options", "nosniff")
		// w.Header().Set("X-Frame-Options", "SAMEORIGIN")
		// w.Header().Set("X-XSS-Protection", "1; mode=block")

		next.ServeHTTP(w, r)
	})
}
