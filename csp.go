package helmet

import (
	"fmt"
	"net/http"
	"strings"
)

// HeaderContentSecurityPolicy is the Content Security Policy HTTP security header.
const HeaderContentSecurityPolicy = "Content-Security-Policy"

// List of all Content Security Policy directives.
const (
	DirectiveBaseURI                 CSPDirective = "base-uri"
	DirectiveBlockAllMixedContent    CSPDirective = "block-all-mixed-content"
	DirectiveChildSrc                CSPDirective = "child-src"
	DirectiveConnectSrc              CSPDirective = "connect-src"
	DirectiveDefaultSrc              CSPDirective = "default-src"
	DirectiveFontSrc                 CSPDirective = "font-src"
	DirectiveFormAction              CSPDirective = "form-action"
	DirectiveFrameAncestors          CSPDirective = "frame-ancestors"
	DirectiveFrameSrc                CSPDirective = "frame-src"
	DirectiveImgSrc                  CSPDirective = "img-src"
	DirectiveManifestSrc             CSPDirective = "manifest-src"
	DirectiveMediaSrc                CSPDirective = "media-src"
	DirectiveNavigateTo              CSPDirective = "navigate-to"
	DirectiveObjectSrc               CSPDirective = "object-src"
	DirectivePluginTypes             CSPDirective = "plugin-types"
	DirectivePrefetchSrc             CSPDirective = "prefetch-src"
	DirectiveReportTo                CSPDirective = "report-to"
	DirectiveSandbox                 CSPDirective = "sandbox"
	DirectiveScriptSrc               CSPDirective = "script-src"
	DirectiveScriptSrcAttr           CSPDirective = "script-src-attr"
	DirectiveScriptSrcElem           CSPDirective = "script-src-elem"
	DirectiveStyleSrc                CSPDirective = "style-src"
	DirectiveStyleSrcAttr            CSPDirective = "style-src-attr"
	DirectiveStyleSrcElem            CSPDirective = "style-src-elem"
	DirectiveTrustedTypes            CSPDirective = "trusted-types"
	DirectiveUpgradeInsecureRequests CSPDirective = "upgrade-insecure-requests"
	DirectiveWorkerSrc               CSPDirective = "worker-src"

	// deprecated
	DeprecatedDirectiveReferrer      CSPDirective = "referrer"   // use 'Referrer-Policy' HTTP header instead
	DeprecatedDirectiveReportURI     CSPDirective = "report-uri" // use 'report-to' CSP directive instead
	DeprecatedDirectiveRequireSriFor CSPDirective = "require-sri-for"
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

// CSPDirective represents a Content Security Policy directive.
type CSPDirective string

// ContentSecurityPolicy represents the Content-Security-Policy HTTP security header.
type ContentSecurityPolicy struct {
	policies map[CSPDirective][]string

	cache string
}

// NewContentSecurityPolicy creates a new ContentSecurityPolicy.
func NewContentSecurityPolicy(policies map[CSPDirective][]string) *ContentSecurityPolicy {
	if policies == nil {
		return EmptyContentSecurityPolicy()
	}
	return &ContentSecurityPolicy{policies, ""}
}

// EmptyContentSecurityPolicy creates a blank slate ContentSecurityPolicy.
func EmptyContentSecurityPolicy() *ContentSecurityPolicy {
	return NewContentSecurityPolicy(make(map[CSPDirective][]string))
}

// Add adds a directive and its sources.
func (csp *ContentSecurityPolicy) Add(directive CSPDirective, sources ...string) {
	if len(directive) == 0 {
		return
	}
	csp.cache = ""

	csp.create(directive)
	for _, source := range sources {
		csp.policies[directive] = append(csp.policies[directive], source)
	}
}

func (csp *ContentSecurityPolicy) create(directive CSPDirective) {
	if len(directive) == 0 {
		return
	}
	csp.cache = ""

	if _, ok := csp.policies[directive]; !ok {
		csp.policies[directive] = []string{}
	}
}

// Remove removes a directive and its sources.
func (csp *ContentSecurityPolicy) Remove(directives ...CSPDirective) {
	if len(directives) == 0 {
		return
	}

	didRemove := false
	for _, directive := range directives {
		if _, ok := csp.policies[directive]; ok {
			didRemove = true
			delete(csp.policies, directive)
		}
	}

	if didRemove {
		csp.cache = ""
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
