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

// Empty returns whether the X-DNS-Prefetch-Control is empty.
func (dns XDNSPrefetchControl) Empty() bool {
	return dns.String() == ""
}

// Header adds the X-DNS-Prefetch-Control HTTP security header to the given http.ResponseWriter.
func (dns XDNSPrefetchControl) Header(w http.ResponseWriter) {
	if !dns.Empty() {
		w.Header().Set(HeaderXDNSPrefetchControl, dns.String())
	}
}
