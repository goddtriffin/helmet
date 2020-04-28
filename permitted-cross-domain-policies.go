package helmet

import "net/http"

// HeaderPermittedCrossDomainPolicies is the Permitted Cross Domain Policies HTTP header.
const HeaderPermittedCrossDomainPolicies = "X-Permitted-Cross-Domain-Policies"

// Permitted Cross Domain Policies options.
const (
	PermittedCrossDomainPoliciesNone          PermittedCrossDomainPolicies = "none"
	PermittedCrossDomainPoliciesMasterOnly    PermittedCrossDomainPolicies = "master-only"
	PermittedCrossDomainPoliciesByContentType PermittedCrossDomainPolicies = "by-content-type"
	PermittedCrossDomainPoliciesByFTPFilename PermittedCrossDomainPolicies = "by-ftp-filename"
	PermittedCrossDomainPoliciesAll           PermittedCrossDomainPolicies = "all"
)

// PermittedCrossDomainPolicies represents the Permitted Cross Domain Policies HTTP header.
type PermittedCrossDomainPolicies string

func (pcdp PermittedCrossDomainPolicies) String() string {
	return string(pcdp)
}

// Exists returns whether the DNSPrefetchControl has been set.
func (pcdp PermittedCrossDomainPolicies) Exists() bool {
	if pcdp.String() == "" {
		return false
	}

	return true
}

// AddHeader adds the X-DNS-Prefetch-Control HTTP header to the given ResponseWriter.
func (pcdp PermittedCrossDomainPolicies) AddHeader(w http.ResponseWriter) {
	if pcdp.Exists() {
		w.Header().Set(HeaderPermittedCrossDomainPolicies, pcdp.String())
	}
}
