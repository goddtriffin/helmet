package helmet

import "net/http"

// HeaderPermittedCrossDomainPolicies is the X-Permitted-Cross-Domain-Policies HTTP security header.
const HeaderPermittedCrossDomainPolicies = "X-Permitted-Cross-Domain-Policies"

// X-Permitted-Cross-Domain-Policies options.
const (
	PermittedCrossDomainPoliciesNone          PermittedCrossDomainPolicies = "none"
	PermittedCrossDomainPoliciesMasterOnly    PermittedCrossDomainPolicies = "master-only"
	PermittedCrossDomainPoliciesByContentType PermittedCrossDomainPolicies = "by-content-type"
	PermittedCrossDomainPoliciesByFTPFilename PermittedCrossDomainPolicies = "by-ftp-filename"
	PermittedCrossDomainPoliciesAll           PermittedCrossDomainPolicies = "all"
)

// PermittedCrossDomainPolicies represents the X-Permitted-Cross-Domain-Policies HTTP security header.
type PermittedCrossDomainPolicies string

func (pcdp PermittedCrossDomainPolicies) String() string {
	return string(pcdp)
}

// Exists returns whether the X-Permitted-Cross-Domain-Policies has been set.
func (pcdp PermittedCrossDomainPolicies) Exists() bool {
	if pcdp.String() == "" {
		return false
	}

	return true
}

// AddHeader adds the X-DNS-Prefetch-Control HTTP security header to the given http.ResponseWriter.
func (pcdp PermittedCrossDomainPolicies) AddHeader(w http.ResponseWriter) {
	if pcdp.Exists() {
		w.Header().Set(HeaderPermittedCrossDomainPolicies, pcdp.String())
	}
}
