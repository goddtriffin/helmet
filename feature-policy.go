package helmet

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
	DirectiveGeolocation                 FeaturePolicyDirective = "geolocation"
	DirectiveGyroscope                   FeaturePolicyDirective = "gyroscope"
	DirectiveLayoutAnimations            FeaturePolicyDirective = "layout-animations"
	DirectiveLegacyImageFormats          FeaturePolicyDirective = "legacy-image-formats"
	DirectiveMagnetometer                FeaturePolicyDirective = "magnetometer"
	DirectiveMicrophone                  FeaturePolicyDirective = "microphone"
	DirectiveMidi                        FeaturePolicyDirective = "midi"
	DirectiveNavigationOverride          FeaturePolicyDirective = "navigation-override"
	DirectiveOversizedImages             FeaturePolicyDirective = "oversized-images"
	DirectivePayment                     FeaturePolicyDirective = "payment"
	DirectivePictureInPicture            FeaturePolicyDirective = "picture-in-picture"
	DirectivePublicKeyCredentials        FeaturePolicyDirective = "publickey-credentials"
	DirectiveSyncXHR                     FeaturePolicyDirective = "sync-xhr"
	DirectiveUSB                         FeaturePolicyDirective = "usb"
	DirectiveWakeLock                    FeaturePolicyDirective = "wake-lock"
	DirectiveXRSpacialTracking           FeaturePolicyDirective = "xr-spatial-tracking"

	// deprecated
	DeprecatedDirectiveVibrate FeaturePolicyDirective = "vibrate"
	DeprecatedDirectiveVR      FeaturePolicyDirective = "vr"
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
)

// FeaturePolicy represents the Feature-Policy HTTP security header.
type FeaturePolicy struct {
	policies map[FeaturePolicyDirective][]FeaturePolicyOrigin

	cache string
}

// NewFeaturePolicy creates a new FeaturePolicy.
func NewFeaturePolicy(policies map[FeaturePolicyDirective][]FeaturePolicyOrigin) *FeaturePolicy {
	if policies == nil {
		return EmptyFeaturePolicy()
	}
	return &FeaturePolicy{policies, ""}
}

// EmptyFeaturePolicy creates a blank slate FeaturePolicy.
func EmptyFeaturePolicy() *FeaturePolicy {
	return NewFeaturePolicy(make(map[FeaturePolicyDirective][]FeaturePolicyOrigin))
}
