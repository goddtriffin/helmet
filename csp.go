package helmet

import (
	"fmt"
	"net/http"
	"strings"
)

// HeaderContentSecurityPolicy is the Content Security Policy HTTP header.
const HeaderContentSecurityPolicy = "Content-Security-Policy"

// List of all Content Security Policy directives.
const (
	DirectiveBaseURI                 = "base-uri"
	DirectiveBlockAllMixedContent    = "block-all-mixed-content"
	DirectiveChildSrc                = "child-src"
	DirectiveConnectSrc              = "connect-src"
	DirectiveDefaultSrc              = "default-src"
	DirectiveFontSrc                 = "font-src"
	DirectiveFormAction              = "form-action"
	DirectiveFrameAncestors          = "frame-ancestors"
	DirectiveFrameSrc                = "frame-src"
	DirectiveImgSrc                  = "img-src"
	DirectiveManifestSrc             = "manifest-src"
	DirectiveMediaSrc                = "media-src"
	DirectiveNavigateTo              = "navigate-to"
	DirectiveObjectSrc               = "object-src"
	DirectivePluginTypes             = "plugin-types"
	DirectivePrefetchSrc             = "prefetch-src"
	DirectiveReportTo                = "report-to"
	DirectiveSandbox                 = "sandbox"
	DirectiveScriptSrc               = "script-src"
	DirectiveScriptSrcAttr           = "script-src-attr"
	DirectiveScriptSrcElem           = "script-src-elem"
	DirectiveStyleSrc                = "style-src"
	DirectiveStyleSrcAttr            = "style-src-attr"
	DirectiveStyleSrcElem            = "style-src-elem"
	DirectiveTrustedTypes            = "trusted-types"
	DirectiveUpgradeInsecureRequests = "upgrade-insecure-requests"
	DirectiveWorkerSrc               = "worker-src"

	// deprecated
	DeprecatedDirectiveReferrer      = "referrer"   // use 'Referrer-Policy' header instead
	DeprecatedDirectiveReportURI     = "report-uri" // use 'report-to' directive instead
	DeprecatedDirectiveRequireSriFor = "require-sri-for"
)

// List of all Content Security Policy sources.
const (
	SourceWildcard      = "*"
	SourceNone          = "'none'"
	SourceSelf          = "'self'"
	SourceHTTP          = "http:"
	SourceHTTPS         = "https:"
	SourceData          = "data:"
	SourceMediastream   = "mediastream:"
	SourceBlob          = "blob:"
	SourceFilesystem    = "filesystem:"
	SourceUnsafeEval    = "'unsafe-eval'"
	SourceUnsafeHashes  = "'unsafe-hashes'"
	SourceUnsafeInline  = "'unsafe-inline'"
	SourceStrictDynamic = "'strict-dynamic'"
	SourceReportSample  = "'report-sample'"
)

// List of all DeprecatedDirectiveReferrer values.
const (
	DeprecatedReferrerNone                  = "\"none\""
	DeprecatedReferrerNoReferrer            = "\"no-referrer\""
	DeprecatedReferrerNoneWhenDowngrade     = "\"none-when-downgrade\""
	DeprecatedReferrerOrigin                = "\"origin\""
	DeprecatedReferrerOriginWhenCrossOrigin = "\"origin-when-cross-origin\""
	DeprecatedReferrerUnsafeURL             = "\"unsafe-url\""
)

// List of all DirectiveSandbox values.
const (
	SandboxAllowDownloadsWithoutUserActivation  = "allow-downloads-without-user-activation"
	SandboxAllowForms                           = "allow-forms"
	SandboxAllowModals                          = "allow-modals"
	SandboxAllowOrientationLock                 = "allow-orientation-lock"
	SandboxAllowPointerLock                     = "allow-pointer-lock"
	SandboxAllowPopups                          = "allow-popups"
	SandboxAllowPopupsToEscapeSandbox           = "allow-popups-to-escape-sandbox"
	SandboxAllowPresentation                    = "allow-presentation"
	SandboxAllowSameOrigin                      = "allow-same-origin"
	SandboxAllowScripts                         = "allow-scripts"
	SandboxAllowStorageAccessByUserActivatation = "allow-storage-access-by-user-activation"
	SandboxAllowTopNavigation                   = "allow-top-navigation"
	SandboxAllowTopNavigationByUserActivation   = "allow-top-navigation-by-user-activation"
)

// List of all DirectiveTrustedTypes values.
const (
	TrustedTypesAllowDuplicates = "allow-duplicates"
)

// ContentSecurityPolicy is the Content-Security-Policy HTTP security header.
type ContentSecurityPolicy struct {
	policies map[string][]string

	cache string
}

// NewCSP creates a new ContentSecurityPolicy.
func NewCSP(policies map[string][]string) *ContentSecurityPolicy {
	csp := &ContentSecurityPolicy{}

	if policies != nil {
		csp.policies = policies
	} else {
		csp.policies = make(map[string][]string)
	}

	return csp
}

// EmptyCSP creates a blank slate ContentSecurityPolicy.
func EmptyCSP() *ContentSecurityPolicy {
	return NewCSP(make(map[string][]string))
}

// Add adds a directive and its sources.
func (csp *ContentSecurityPolicy) Add(directive string, sources ...string) {
	if len(directive) == 0 {
		return
	}

	csp.cache = ""

	csp.create(directive)
	for _, source := range sources {
		csp.policies[directive] = append(csp.policies[directive], source)
	}
}

func (csp *ContentSecurityPolicy) create(directive string) {
	if len(directive) == 0 {
		return
	}

	csp.cache = ""

	if _, ok := csp.policies[directive]; !ok {
		csp.policies[directive] = []string{}
	}
}

// String generates the Content-Security-Policy.
func (csp *ContentSecurityPolicy) String() string {
	if csp.cache != "" {
		return csp.cache
	}

	var builder string
	for directive, sources := range csp.policies {
		if len(builder) != 0 {
			builder += " "
		}

		if len(sources) == 0 {
			builder += fmt.Sprintf("%s;", directive)
			continue
		}

		builder += fmt.Sprintf("%s %s;", directive, strings.Join(sources, " "))
	}

	csp.cache = builder
	return csp.cache
}

// Exists returns whether the Content Security Policy contains any policies.
func (csp *ContentSecurityPolicy) Exists() bool {
	if len(csp.policies) == 0 {
		return false
	}

	return true
}

// AddHeader adds the Content Security Policy HTTP header to the given ResponseWriter.
func (csp *ContentSecurityPolicy) AddHeader(w http.ResponseWriter) {
	if csp.Exists() {
		w.Header().Set(HeaderContentSecurityPolicy, csp.String())
	}
}
