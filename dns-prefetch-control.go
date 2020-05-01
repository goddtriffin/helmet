package helmet

import "net/http"

// HeaderXDNSPrefetchControl is the X-DNS-Prefetch-Control HTTP header.
const HeaderXDNSPrefetchControl = "X-DNS-Prefetch-Control"

// X-DNS-Prefetch-Control options.
const (
	XDNSPrefetchControlOn  XDNSPrefetchControl = "on"
	XDNSPrefetchControlOff XDNSPrefetchControl = "off"
)

// XDNSPrefetchControl represents the X-DNS-Prefetch-Control HTTP security header.
type XDNSPrefetchControl string

func (dns XDNSPrefetchControl) String() string {
	return string(dns)
}

// Exists returns whether the X-DNS-Prefetch-Control has been set.
func (dns XDNSPrefetchControl) Exists() bool {
	if dns.String() == "" {
		return false
	}

	return true
}

// Header adds the X-DNS-Prefetch-Control HTTP security header to the given http.ResponseWriter.
func (dns XDNSPrefetchControl) Header(w http.ResponseWriter) {
	if dns.Exists() {
		w.Header().Set(HeaderXDNSPrefetchControl, dns.String())
	}
}
