package helmet

import (
	"net/http"

	"github.com/valyala/fasthttp"
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
	XXSSProtection                *XXSSProtection
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
		XXSSProtection:                NewXXSSProtection(true, DirectiveModeBlock, ""),
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
		XXSSProtection:          EmptyXXSSProtection(),
	}
}

// Secure is the net/http middleware handler.
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
		h.XXSSProtection.Header(w)

		next.ServeHTTP(w, r)
	})
}

// SecureFastHTTP is the fasthttp middleware handler.
func (h *Helmet) SecureFastHTTP(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		h.ContentSecurityPolicy.HeaderFastHTTP(ctx)
		h.XContentTypeOptions.HeaderFastHTTP(ctx)
		h.XDNSPrefetchControl.HeaderFastHTTP(ctx)
		h.XDownloadOptions.HeaderFastHTTP(ctx)
		h.ExpectCT.HeaderFastHTTP(ctx)
		h.FeaturePolicy.HeaderFastHTTP(ctx)
		h.XFrameOptions.HeaderFastHTTP(ctx)
		h.XPermittedCrossDomainPolicies.HeaderFastHTTP(ctx)
		h.XPoweredBy.HeaderFastHTTP(ctx)
		h.ReferrerPolicy.HeaderFastHTTP(ctx)
		h.StrictTransportSecurity.HeaderFastHTTP(ctx)
		h.XXSSProtection.HeaderFastHTTP(ctx)

		next(ctx)
	}
}
