package helmet

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

// FeaturePolicyDirective represents a Feature-Policy directive.
type FeaturePolicyDirective string

// FeaturePolicy represents the Feature-Policy HTTP security header.
type FeaturePolicy struct {
}
