package helmet

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

// HeaderXPermittedCrossDomainPolicies is the X-Permitted-Cross-Domain-Policies HTTP security header.
const HeaderXPermittedCrossDomainPolicies = "X-Permitted-Cross-Domain-Policies"

// X-Permitted-Cross-Domain-Policies options.
const (
	PermittedCrossDomainPoliciesNone          XPermittedCrossDomainPolicies = "none"
	PermittedCrossDomainPoliciesMasterOnly    XPermittedCrossDomainPolicies = "master-only"
	PermittedCrossDomainPoliciesByContentType XPermittedCrossDomainPolicies = "by-content-type"
	PermittedCrossDomainPoliciesByFTPFilename XPermittedCrossDomainPolicies = "by-ftp-filename"
	PermittedCrossDomainPoliciesAll           XPermittedCrossDomainPolicies = "all"
)

// XPermittedCrossDomainPolicies represents the X-Permitted-Cross-Domain-Policies HTTP security header.
type XPermittedCrossDomainPolicies string

func (cdp XPermittedCrossDomainPolicies) String() string {
	return string(cdp)
}

// Empty returns whether the X-Permitted-Cross-Domain-Policies is empty.
func (cdp XPermittedCrossDomainPolicies) Empty() bool {
	return cdp.String() == ""
}

// Header adds the X-DNS-Prefetch-Control HTTP security header to the given http.ResponseWriter.
func (cdp XPermittedCrossDomainPolicies) Header(w http.ResponseWriter) {
	if !cdp.Empty() {
		w.Header().Set(HeaderXPermittedCrossDomainPolicies, cdp.String())
	}
}

// HeaderFastHTTP adds the X-DNS-Prefetch-Control HTTP security header to the given *fasthttp.RequestCtx.
func (cdp XPermittedCrossDomainPolicies) HeaderFastHTTP(ctx *fasthttp.RequestCtx) {
	if !cdp.Empty() {
		ctx.Response.Header.Set(HeaderXPermittedCrossDomainPolicies, cdp.String())
	}
}
