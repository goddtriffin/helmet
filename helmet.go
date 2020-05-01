package helmet

import (
	"net/http"
)

// Helmet is a HTTP security middleware for Go(lang) inspired by HelmetJS for Express.js.
type Helmet struct {
	ContentSecurityPolicy        *ContentSecurityPolicy
	DNSPrefetchControl           DNSPrefetchControl
	ExpectCT                     *ExpectCT
	FeaturePolicy                *FeaturePolicy
	FrameOptions                 FrameOptions
	PermittedCrossDomainPolicies PermittedCrossDomainPolicies
}

// Default creates a new Helmet with default settings.
func Default() *Helmet {
	return &Helmet{
		ContentSecurityPolicy:        EmptyContentSecurityPolicy(),
		DNSPrefetchControl:           DNSPrefetchControlOff,
		ExpectCT:                     EmptyExpectCT(),
		FeaturePolicy:                EmptyFeaturePolicy(),
		FrameOptions:                 FrameOptionsSameOrigin,
		PermittedCrossDomainPolicies: "",
	}
}

// Empty creates a new Helmet.
func Empty() *Helmet {
	return &Helmet{
		ContentSecurityPolicy: EmptyContentSecurityPolicy(),
		ExpectCT:              EmptyExpectCT(),
		FeaturePolicy:         EmptyFeaturePolicy(),
	}
}

// Secure is the middleware handler.
func (h *Helmet) Secure(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ContentSecurityPolicy.Header(w)
		h.DNSPrefetchControl.Header(w)
		h.ExpectCT.Header(w)
		h.FeaturePolicy.Header(w)
		h.FrameOptions.Header(w)
		h.PermittedCrossDomainPolicies.Header(w)

		// w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		// w.Header().Set("X-Content-Type-Options", "nosniff")
		// w.Header().Set("X-XSS-Protection", "1; mode=block")

		next.ServeHTTP(w, r)
	})
}
