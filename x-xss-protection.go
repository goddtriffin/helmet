package helmet

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/valyala/fasthttp"
)

// HeaderXXSSProtection is the X-XSS-Protection HTTP security header.
const HeaderXXSSProtection = "X-XSS-Protection"

// DirectiveModeBlock is the X-XSS-Protection mode=block directive.
const DirectiveModeBlock XXSSProtectionDirective = "mode=block"

// XXSSProtectionDirectiveXSSFiltering is the X-XSS-Protection XSSFiltering directive.
func XXSSProtectionDirectiveXSSFiltering(xssFiltering bool) XXSSProtectionDirective {
	if xssFiltering {
		return "1"
	}
	return "0"
}

// XXSSProtectionDirectiveReportURI is the X-XSS-Protection ReportURI directive.
func XXSSProtectionDirectiveReportURI(reportURI string) XXSSProtectionDirective {
	if reportURI == "" {
		return ""
	}
	return XXSSProtectionDirective(fmt.Sprintf(`report=%s`, reportURI))
}

type (
	// XXSSProtectionDirective represents an X-XSS-Protection directive.
	XXSSProtectionDirective string

	// XXSSProtection represents the X-XSS-Protection HTTP security header.
	XXSSProtection struct {
		XSSFiltering bool
		Mode         XXSSProtectionDirective
		ReportURI    string

		cache string
	}
)

// NewXXSSProtection creates a new X-XSS-Protection.
func NewXXSSProtection(xssFiltering bool, mode XXSSProtectionDirective, reportURI string) *XXSSProtection {
	return &XXSSProtection{
		XSSFiltering: xssFiltering,
		Mode:         mode,
		ReportURI:    reportURI,
	}
}

// EmptyXXSSProtection creates a blank slate X-XSS-Protection.
func EmptyXXSSProtection() *XXSSProtection {
	return NewXXSSProtection(false, "", "")
}

func (xssp *XXSSProtection) String() string {
	if len(xssp.cache) != 0 {
		return xssp.cache
	}

	builder := []string{
		string(XXSSProtectionDirectiveXSSFiltering(xssp.XSSFiltering)),
	}

	if xssp.Mode != "" {
		builder = append(builder, string(DirectiveModeBlock))
	}

	if xssp.ReportURI != "" {
		builder = append(builder, string(XXSSProtectionDirectiveReportURI(xssp.ReportURI)))
	}

	xssp.cache = strings.Join(builder, "; ")
	return xssp.cache
}

// Empty returns whether the X-XSS-Protection is empty.
func (xssp *XXSSProtection) Empty() bool {
	// no matter what, the only required info (XSS Filtering) will always be present
	// true and false are the only options, and they are both valid
	return false
}

// Header adds the X-XSS-Protection HTTP security header to the given http.ResponseWriter.
func (xssp *XXSSProtection) Header(w http.ResponseWriter) {
	if !xssp.Empty() {
		w.Header().Set(HeaderXXSSProtection, xssp.String())
	}
}

// HeaderFastHTTP adds the X-XSS-Protection HTTP security header to the given *fasthttp.RequestCtx.
func (xssp *XXSSProtection) HeaderFastHTTP(ctx *fasthttp.RequestCtx) {
	if !xssp.Empty() {
		ctx.Response.Header.Set(HeaderXXSSProtection, xssp.String())
	}
}
