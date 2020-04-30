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
	DirectiveBaseURI                 Directive = "base-uri"
	DirectiveBlockAllMixedContent    Directive = "block-all-mixed-content"
	DirectiveChildSrc                Directive = "child-src"
	DirectiveConnectSrc              Directive = "connect-src"
	DirectiveDefaultSrc              Directive = "default-src"
	DirectiveFontSrc                 Directive = "font-src"
	DirectiveFormAction              Directive = "form-action"
	DirectiveFrameAncestors          Directive = "frame-ancestors"
	DirectiveFrameSrc                Directive = "frame-src"
	DirectiveImgSrc                  Directive = "img-src"
	DirectiveManifestSrc             Directive = "manifest-src"
	DirectiveMediaSrc                Directive = "media-src"
	DirectiveNavigateTo              Directive = "navigate-to"
	DirectiveObjectSrc               Directive = "object-src"
	DirectivePluginTypes             Directive = "plugin-types"
	DirectivePrefetchSrc             Directive = "prefetch-src"
	DirectiveReportTo                Directive = "report-to"
	DirectiveSandbox                 Directive = "sandbox"
	DirectiveScriptSrc               Directive = "script-src"
	DirectiveScriptSrcAttr           Directive = "script-src-attr"
	DirectiveScriptSrcElem           Directive = "script-src-elem"
	DirectiveStyleSrc                Directive = "style-src"
	DirectiveStyleSrcAttr            Directive = "style-src-attr"
	DirectiveStyleSrcElem            Directive = "style-src-elem"
	DirectiveTrustedTypes            Directive = "trusted-types"
	DirectiveUpgradeInsecureRequests Directive = "upgrade-insecure-requests"
	DirectiveWorkerSrc               Directive = "worker-src"

	// deprecated
	DeprecatedDirectiveReferrer      Directive = "referrer"   // use 'Referrer-Policy' HTTP header instead
	DeprecatedDirectiveReportURI     Directive = "report-uri" // use 'report-to' CSP directive instead
	DeprecatedDirectiveRequireSriFor Directive = "require-sri-for"
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

// Directive represents a Content Security Policy directive.
type Directive string

// ContentSecurityPolicy represents the Content-Security-Policy HTTP security header.
type ContentSecurityPolicy struct {
	policies map[Directive][]string

	cache string
}

// NewContentSecurityPolicy creates a new ContentSecurityPolicy.
func NewContentSecurityPolicy(policies map[Directive][]string) *ContentSecurityPolicy {
	if policies == nil {
		return EmptyContentSecurityPolicy()
	}
	return &ContentSecurityPolicy{policies, ""}
}

// EmptyContentSecurityPolicy creates a blank slate ContentSecurityPolicy.
func EmptyContentSecurityPolicy() *ContentSecurityPolicy {
	return NewContentSecurityPolicy(make(map[Directive][]string))
}

// Add adds a directive and its sources.
func (csp *ContentSecurityPolicy) Add(directive Directive, sources ...string) {
	if len(directive) == 0 {
		return
	}
	csp.cache = ""

	csp.create(directive)
	for _, source := range sources {
		csp.policies[directive] = append(csp.policies[directive], source)
	}
}

func (csp *ContentSecurityPolicy) create(directive Directive) {
	if len(directive) == 0 {
		return
	}
	csp.cache = ""

	if _, ok := csp.policies[directive]; !ok {
		csp.policies[directive] = []string{}
	}
}

// Remove removes a directive and its sources.
func (csp *ContentSecurityPolicy) Remove(directives ...Directive) {
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
