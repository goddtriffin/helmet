package helmet

import (
	"fmt"
	"net/http"
	"strings"
)

// HeaderFeaturePolicy is the Feature-Policy HTTP security header.
const HeaderFeaturePolicy = "Feature-Policy"

// List of all Feature-Policy directives.
const (
	DirectiveAccelerometer               FeaturePolicyDirective = "accelerometer"
	DirectiveAmbientLightSensor          FeaturePolicyDirective = "ambient-light-sensor"
	DirectiveAutoplay                    FeaturePolicyDirective = "autoplay"
	DirectiveBattery                     FeaturePolicyDirective = "battery"
	DirectiveCamera                      FeaturePolicyDirective = "camera"
	DirectiveDisplayCapture              FeaturePolicyDirective = "display-capture"
	DirectiveDocumentDomain              FeaturePolicyDirective = "document-domain"
	DirectiveEncryptedMedia              FeaturePolicyDirective = "encrypted-media"
	DirectiveExecutionWhileNotRendered   FeaturePolicyDirective = "execution-while-not-rendered"
	DirectiveExecutionWhileOutOfViewport FeaturePolicyDirective = "execution-while-out-of-viewport"
	DirectiveFullscreen                  FeaturePolicyDirective = "fullscreen"
	DirectiveGamepad                     FeaturePolicyDirective = "gamepad"
	DirectiveGeolocation                 FeaturePolicyDirective = "geolocation"
	DirectiveGyroscope                   FeaturePolicyDirective = "gyroscope"
	DirectiveMagnetometer                FeaturePolicyDirective = "magnetometer"
	DirectiveMicrophone                  FeaturePolicyDirective = "microphone"
	DirectiveMidi                        FeaturePolicyDirective = "midi"
	DirectiveNavigationOverride          FeaturePolicyDirective = "navigation-override"
	DirectiveOversizedImages             FeaturePolicyDirective = "oversized-images"
	DirectivePayment                     FeaturePolicyDirective = "payment"
	DirectivePictureInPicture            FeaturePolicyDirective = "picture-in-picture"
	DirectivePublicKeyCredentialsGet     FeaturePolicyDirective = "publickey-credentials-get"
	DirectiveSpeakerSelection            FeaturePolicyDirective = "speaker-selection"
	DirectiveSyncXHR                     FeaturePolicyDirective = "sync-xhr"
	DirectiveUSB                         FeaturePolicyDirective = "usb"
	DirectiveScreenWakeLock              FeaturePolicyDirective = "screen-wake-lock"
	DirectiveWebShare                    FeaturePolicyDirective = "web-share"
	DirectiveXRSpacialTracking           FeaturePolicyDirective = "xr-spatial-tracking"
)

// deprecated
const (
	NonStandardDirectiveLayoutAnimations   FeaturePolicyDirective = "layout-animations"
	NonStandardDirectiveLegacyImageFormats FeaturePolicyDirective = "legacy-image-formats"
	NonStandardDirectiveOversizedImages    FeaturePolicyDirective = "oversized-images"
	NonStandardDirectiveUnoptimizedImages  FeaturePolicyDirective = "unoptimized-images"
	NonStandardDirectiveUnsizedMedia       FeaturePolicyDirective = "unsized-media"
)

// List of all Feature-Policy origins.
const (
	OriginWildcard FeaturePolicyOrigin = "*"
	OriginSelf     FeaturePolicyOrigin = "'self'"
	OriginSrc      FeaturePolicyOrigin = "'src'"
	OriginNone     FeaturePolicyOrigin = "'none'"
)

type (
	// FeaturePolicyDirective represents a Feature-Policy directive.
	FeaturePolicyDirective string

	// FeaturePolicyOrigin represents a Feature-Policy origin.
	FeaturePolicyOrigin string

	// FeaturePolicy represents the Feature-Policy HTTP security header.
	FeaturePolicy struct {
		policies map[FeaturePolicyDirective][]FeaturePolicyOrigin

		cache string
	}
)

// NewFeaturePolicy creates a new Feature-Policy.
func NewFeaturePolicy(policies map[FeaturePolicyDirective][]FeaturePolicyOrigin) *FeaturePolicy {
	if policies == nil {
		return EmptyFeaturePolicy()
	}
	return &FeaturePolicy{policies, ""}
}

// EmptyFeaturePolicy creates a blank slate Feature-Policy.
func EmptyFeaturePolicy() *FeaturePolicy {
	return NewFeaturePolicy(make(map[FeaturePolicyDirective][]FeaturePolicyOrigin))
}

// Add adds a directive and its origins.
func (fp *FeaturePolicy) Add(directive FeaturePolicyDirective, origins ...FeaturePolicyOrigin) {
	if len(directive) == 0 || len(origins) == 0 {
		return
	}
	fp.cache = ""

	fp.create(directive)
	for _, origin := range origins {
		fp.policies[directive] = append(fp.policies[directive], origin)
	}
}

func (fp *FeaturePolicy) create(directive FeaturePolicyDirective) {
	if len(directive) == 0 {
		return
	}
	fp.cache = ""

	if _, ok := fp.policies[directive]; !ok {
		fp.policies[directive] = []FeaturePolicyOrigin{}
	}
}

// Remove removes a directive and its origins.
func (fp *FeaturePolicy) Remove(directives ...FeaturePolicyDirective) {
	if len(directives) == 0 {
		return
	}

	didRemove := false
	for _, directive := range directives {
		if _, ok := fp.policies[directive]; ok {
			didRemove = true
			delete(fp.policies, directive)
		}
	}

	if didRemove {
		fp.cache = ""
	}
}

// String generates the Feature-Policy.
func (fp *FeaturePolicy) String() string {
	if fp.cache != "" {
		return fp.cache
	}

	var policies = []string{}
	for directive, origins := range fp.policies {
		originsAsStrings := []string{}
		for _, origin := range origins {
			originsAsStrings = append(originsAsStrings, string(origin))
		}

		policies = append(policies, fmt.Sprintf("%s %s", directive, strings.Join(originsAsStrings, " ")))
	}

	fp.cache = strings.Join(policies, "; ")
	return fp.cache
}

// Empty returns whether the Feature-Policy is empty.
func (fp *FeaturePolicy) Empty() bool {
	return len(fp.policies) == 0
}

// Header adds the Feature-Policy HTTP security header to the given http.ResponseWriter.
func (fp *FeaturePolicy) Header(w http.ResponseWriter) {
	if !fp.Empty() {
		w.Header().Set(HeaderFeaturePolicy, fp.String())
	}
}
