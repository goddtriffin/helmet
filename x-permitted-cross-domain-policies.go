package helmet

import "net/http"

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

// Exists returns whether the X-Permitted-Cross-Domain-Policies has been set.
func (cdp XPermittedCrossDomainPolicies) Exists() bool {
	if cdp.String() == "" {
		return false
	}

	return true
}

// Header adds the X-DNS-Prefetch-Control HTTP security header to the given http.ResponseWriter.
func (cdp XPermittedCrossDomainPolicies) Header(w http.ResponseWriter) {
	if cdp.Exists() {
		w.Header().Set(HeaderXPermittedCrossDomainPolicies, cdp.String())
	}
}
