package helmet

import (
	"net/http"
)

// Helmet is a HTTP security middleware for Go(lang) inspired by HelmetJS for Express.js.
type Helmet struct {
	ContentSecurityPolicy         *ContentSecurityPolicy
	XDNSPrefetchControl           XDNSPrefetchControl
	ExpectCT                      *ExpectCT
	FeaturePolicy                 *FeaturePolicy
	XFrameOptions                 XFrameOptions
	XPermittedCrossDomainPolicies XPermittedCrossDomainPolicies
	XPoweredBy                    *XPoweredBy
}

// Default creates a new Helmet with default settings.
func Default() *Helmet {
	return &Helmet{
		ContentSecurityPolicy:         EmptyContentSecurityPolicy(),
		XDNSPrefetchControl:           XDNSPrefetchControlOff,
		ExpectCT:                      EmptyExpectCT(),
		FeaturePolicy:                 EmptyFeaturePolicy(),
		XFrameOptions:                 XFrameOptionsSameOrigin,
		XPermittedCrossDomainPolicies: "",
		XPoweredBy:                    NewXPoweredBy(true, ""),
	}
}

// Empty creates a new Helmet.
func Empty() *Helmet {
	return &Helmet{
		ContentSecurityPolicy: EmptyContentSecurityPolicy(),
		ExpectCT:              EmptyExpectCT(),
		FeaturePolicy:         EmptyFeaturePolicy(),
		XPoweredBy:            EmptyXPoweredBy(),
	}
}

// Secure is the middleware handler.
func (h *Helmet) Secure(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ContentSecurityPolicy.Header(w)
		h.XDNSPrefetchControl.Header(w)
		h.ExpectCT.Header(w)
		h.FeaturePolicy.Header(w)
		h.XFrameOptions.Header(w)
		h.XPermittedCrossDomainPolicies.Header(w)
		h.XPoweredBy.Header(w)

		next.ServeHTTP(w, r)
	})
}
