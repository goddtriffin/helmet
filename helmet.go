package helmet

import (
	"net/http"
)

// DNS Prefetch Control options.
const (
	DNSPrefetchControlOn  = "on"
	DNSPrefetchControlOff = "off"
)

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
	DNSPrefetchControl           string
	PermittedCrossDomainPolicies string
}

// New creates a new Helmet.
func New() *Helmet {
	return &Helmet{
		ContentSecurityPolicy:        NewContentSecurityPolicy(nil),
		DNSPrefetchControl:           DNSPrefetchControlOn,
		PermittedCrossDomainPolicies: PermittedCrossDomainPoliciesAll,
	}
}

// Default creates a new Helmet with default settings.
func Default() *Helmet {
	return &Helmet{
		ContentSecurityPolicy:        NewContentSecurityPolicy(nil),
		DNSPrefetchControl:           DNSPrefetchControlOff,
		PermittedCrossDomainPolicies: PermittedCrossDomainPoliciesNone,
	}
}

// Secure is the middleware handler.
func (h *Helmet) Secure(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if h.ContentSecurityPolicy.Exists() {
			w.Header().Set("Content-Security-Policy", h.ContentSecurityPolicy.String())
		}

		w.Header().Set("X-DNS-Prefetch-Control", h.DNSPrefetchControl)
		w.Header().Set("X-Permitted-Cross-Domain-Policies", h.PermittedCrossDomainPolicies)

		// w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		// w.Header().Set("X-Content-Type-Options", "nosniff")
		// w.Header().Set("X-Frame-Options", "SAMEORIGIN")
		// w.Header().Set("X-XSS-Protection", "1; mode=block")

		next.ServeHTTP(w, r)
	})
}
