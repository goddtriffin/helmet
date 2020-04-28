package helmet

import "net/http"

// HeaderDNSPrefetchControl is the DNS Prefetch Control HTTP header.
const HeaderDNSPrefetchControl = "X-DNS-Prefetch-Control"

// DNS Prefetch Control options.
const (
	DNSPrefetchControlOn  DNSPrefetchControl = "on"
	DNSPrefetchControlOff DNSPrefetchControl = "off"
)

// DNSPrefetchControl represents the X-DNS-Prefetch-Control HTTP security header.
type DNSPrefetchControl string

func (dns DNSPrefetchControl) String() string {
	return string(dns)
}

// Exists returns whether the DNSPrefetchControl has been set.
func (dns DNSPrefetchControl) Exists() bool {
	if dns.String() == "" {
		return false
	}

	return true
}

// AddHeader adds the X-DNS-Prefetch-Control HTTP header to the given ResponseWriter.
func (dns DNSPrefetchControl) AddHeader(w http.ResponseWriter) {
	if dns.Exists() {
		w.Header().Set(HeaderDNSPrefetchControl, dns.String())
	}
}
