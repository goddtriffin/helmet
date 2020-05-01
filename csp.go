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
	SourceWildcard      CSPSource = "*"
	SourceNone          CSPSource = "'none'"
	SourceSelf          CSPSource = "'self'"
	SourceHTTP          CSPSource = "http:"
	SourceHTTPS         CSPSource = "https:"
	SourceData          CSPSource = "data:"
	SourceMediastream   CSPSource = "mediastream:"
	SourceBlob          CSPSource = "blob:"
	SourceFilesystem    CSPSource = "filesystem:"
	SourceUnsafeEval    CSPSource = "'unsafe-eval'"
	SourceUnsafeHashes  CSPSource = "'unsafe-hashes'"
	SourceUnsafeInline  CSPSource = "'unsafe-inline'"
	SourceStrictDynamic CSPSource = "'strict-dynamic'"
	SourceReportSample  CSPSource = "'report-sample'"
)

// List of all DeprecatedDirectiveReferrer values.
const (
	DeprecatedReferrerNone                  CSPSource = "\"none\""
	DeprecatedReferrerNoReferrer            CSPSource = "\"no-referrer\""
	DeprecatedReferrerNoneWhenDowngrade     CSPSource = "\"none-when-downgrade\""
	DeprecatedReferrerOrigin                CSPSource = "\"origin\""
	DeprecatedReferrerOriginWhenCrossOrigin CSPSource = "\"origin-when-cross-origin\""
	DeprecatedReferrerUnsafeURL             CSPSource = "\"unsafe-url\""
)

// List of all DirectiveSandbox values.
const (
	SandboxAllowDownloadsWithoutUserActivation  CSPSource = "allow-downloads-without-user-activation"
	SandboxAllowForms                           CSPSource = "allow-forms"
	SandboxAllowModals                          CSPSource = "allow-modals"
	SandboxAllowOrientationLock                 CSPSource = "allow-orientation-lock"
	SandboxAllowPointerLock                     CSPSource = "allow-pointer-lock"
	SandboxAllowPopups                          CSPSource = "allow-popups"
	SandboxAllowPopupsToEscapeSandbox           CSPSource = "allow-popups-to-escape-sandbox"
	SandboxAllowPresentation                    CSPSource = "allow-presentation"
	SandboxAllowSameOrigin                      CSPSource = "allow-same-origin"
	SandboxAllowScripts                         CSPSource = "allow-scripts"
	SandboxAllowStorageAccessByUserActivatation CSPSource = "allow-storage-access-by-user-activation"
	SandboxAllowTopNavigation                   CSPSource = "allow-top-navigation"
	SandboxAllowTopNavigationByUserActivation   CSPSource = "allow-top-navigation-by-user-activation"
)

// List of all DirectiveTrustedTypes values.
const (
	TrustedTypesAllowDuplicates CSPSource = "allow-duplicates"
)

type (
	// CSPDirective represents a Content Security Policy directive.
	CSPDirective string

	// CSPSource represents a Content Security Policy source.
	CSPSource string

	// ContentSecurityPolicy represents the Content-Security-Policy HTTP security header.
	ContentSecurityPolicy struct {
		policies map[CSPDirective][]CSPSource

		cache string
	}
)

// NewContentSecurityPolicy creates a new ContentSecurityPolicy.
func NewContentSecurityPolicy(policies map[CSPDirective][]CSPSource) *ContentSecurityPolicy {
	if policies == nil {
		return EmptyContentSecurityPolicy()
	}
	return &ContentSecurityPolicy{policies, ""}
}

// EmptyContentSecurityPolicy creates a blank slate ContentSecurityPolicy.
func EmptyContentSecurityPolicy() *ContentSecurityPolicy {
	return NewContentSecurityPolicy(make(map[CSPDirective][]CSPSource))
}

// Add adds a directive and its sources.
func (csp *ContentSecurityPolicy) Add(directive CSPDirective, sources ...CSPSource) {
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
		csp.policies[directive] = []CSPSource{}
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

	var policies = []string{}
	for directive, sources := range csp.policies {
		if len(sources) == 0 {
			policies = append(policies, fmt.Sprintf("%s", directive))
		} else {
			sourcesAsStrings := []string{}
			for _, source := range sources {
				sourcesAsStrings = append(sourcesAsStrings, string(source))
			}

			policies = append(policies, fmt.Sprintf("%s %s", directive, strings.Join(sourcesAsStrings, " ")))
		}
	}

	csp.cache = strings.Join(policies, "; ")
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
