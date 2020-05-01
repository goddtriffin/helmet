package helmet

import "net/http"

// HeaderDNSPrefetchControl is the X-DNS-Prefetch-Control HTTP header.
const HeaderDNSPrefetchControl = "X-DNS-Prefetch-Control"

// X-DNS-Prefetch-Control options.
const (
	DNSPrefetchControlOn  DNSPrefetchControl = "on"
	DNSPrefetchControlOff DNSPrefetchControl = "off"
)

// DNSPrefetchControl represents the X-DNS-Prefetch-Control HTTP security header.
type DNSPrefetchControl string

func (dns DNSPrefetchControl) String() string {
	return string(dns)
}

// Exists returns whether the X-DNS-Prefetch-Control has been set.
func (dns DNSPrefetchControl) Exists() bool {
	if dns.String() == "" {
		return false
	}

	return true
}

// Header adds the X-DNS-Prefetch-Control HTTP security header to the given http.ResponseWriter.
func (dns DNSPrefetchControl) Header(w http.ResponseWriter) {
	if dns.Exists() {
		w.Header().Set(HeaderDNSPrefetchControl, dns.String())
	}
}
