package helmet

import (
	"net/http"
)

// Helmet is a HTTP security middleware for Go(lang) inspired by HelmetJS for Express.js.
type Helmet struct {
	ContentSecurityPolicy         *ContentSecurityPolicy
	XContentTypeOptions           XContentTypeOptions
	XDNSPrefetchControl           XDNSPrefetchControl
	XDownloadOptions              XDownloadOptions
	ExpectCT                      *ExpectCT
	FeaturePolicy                 *FeaturePolicy
	XFrameOptions                 XFrameOptions
	XPermittedCrossDomainPolicies XPermittedCrossDomainPolicies
	XPoweredBy                    *XPoweredBy
	ReferrerPolicy                *ReferrerPolicy
	StrictTransportSecurity       *StrictTransportSecurity
}

// Default creates a new Helmet with default settings.
func Default() *Helmet {
	return &Helmet{
		ContentSecurityPolicy:         EmptyContentSecurityPolicy(),
		XContentTypeOptions:           XContentTypeOptionsNoSniff,
		XDNSPrefetchControl:           XDNSPrefetchControlOff,
		XDownloadOptions:              XDownloadOptionsNoOpen,
		ExpectCT:                      EmptyExpectCT(),
		FeaturePolicy:                 EmptyFeaturePolicy(),
		XFrameOptions:                 XFrameOptionsSameOrigin,
		XPermittedCrossDomainPolicies: "",
		XPoweredBy:                    NewXPoweredBy(true, ""),
		ReferrerPolicy:                EmptyReferrerPolicy(),
		StrictTransportSecurity:       NewStrictTransportSecurity(5184000, true, false),
	}
}

// Empty creates a new Helmet.
func Empty() *Helmet {
	return &Helmet{
		ContentSecurityPolicy:   EmptyContentSecurityPolicy(),
		ExpectCT:                EmptyExpectCT(),
		FeaturePolicy:           EmptyFeaturePolicy(),
		XPoweredBy:              EmptyXPoweredBy(),
		ReferrerPolicy:          EmptyReferrerPolicy(),
		StrictTransportSecurity: EmptyStrictTransportSecurity(),
	}
}

// Secure is the middleware handler.
func (h *Helmet) Secure(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ContentSecurityPolicy.Header(w)
		h.XContentTypeOptions.Header(w)
		h.XDNSPrefetchControl.Header(w)
		h.XDownloadOptions.Header(w)
		h.ExpectCT.Header(w)
		h.FeaturePolicy.Header(w)
		h.XFrameOptions.Header(w)
		h.XPermittedCrossDomainPolicies.Header(w)
		h.XPoweredBy.Header(w)
		h.ReferrerPolicy.Header(w)
		h.StrictTransportSecurity.Header(w)

		next.ServeHTTP(w, r)
	})
}
