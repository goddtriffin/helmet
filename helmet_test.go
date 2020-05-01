package helmet

import (
	"testing"
)

func TestHelmet_Secure_default(t *testing.T) {
	t.Parallel()

	rr, r := newRecorderRequest(t)

	// default Helmet
	helmet := Default()
	addXPoweredByHelmetMiddleware(helmet.Secure(mockNext)).ServeHTTP(rr, r)
	resp := rr.Result()

	t.Run(HeaderXPoweredBy, func(t *testing.T) {
		t.Parallel()

		header := resp.Header.Get(HeaderXPoweredBy)
		if header != "" {
			t.Errorf("X-Powered-By header should be removed\tActual: %s\n", header)
		}
	})

	testCases := []struct {
		name   string
		header string
	}{
		{HeaderContentSecurityPolicy, ""},
		{HeaderXDNSPrefetchControl, XDNSPrefetchControlOff.String()},
		{HeaderExpectCT, ""},
		{HeaderFeaturePolicy, ""},
		{HeaderFrameOptions, FrameOptionsSameOrigin.String()},
		{HeaderPermittedCrossDomainPolicies, ""},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			header := resp.Header.Get(tc.name)
			if header != tc.header {
				t.Errorf("Expected: %s\tActual: %s\n", tc.header, header)
			}
		})
	}

	testMockNext(t, resp)
}

func TestHelmet_Secure_empty(t *testing.T) {
	t.Parallel()

	rr, r := newRecorderRequest(t)

	// blank slate Helmet
	helmet := Empty()
	addXPoweredByHelmetMiddleware(helmet.Secure(mockNext)).ServeHTTP(rr, r)
	resp := rr.Result()

	t.Run("X-Powered-By", func(t *testing.T) {
		t.Parallel()

		header := resp.Header.Get(HeaderXPoweredBy)
		if header != "Helmet" {
			t.Errorf("X-Powered-By is wrong\tExpected: %s\tActual: %s\n", "Helmet", header)
		}
	})

	testCases := []struct {
		header string
	}{
		{HeaderContentSecurityPolicy},
		{HeaderXDNSPrefetchControl},
		{HeaderExpectCT},
		{HeaderFeaturePolicy},
		{HeaderFrameOptions},
		{HeaderPermittedCrossDomainPolicies},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.header, func(t *testing.T) {
			t.Parallel()

			header := resp.Header.Get(tc.header)
			if header != "" {
				t.Errorf("Header exists when it shouldn't: %s\n", header)
			}
		})
	}

	testMockNext(t, resp)
}

func TestHelmet_Secure_custom(t *testing.T) {
	t.Parallel()

	rr, r := newRecorderRequest(t)

	// fill Helmet with custom parameters
	helmet := Empty()
	helmet.ContentSecurityPolicy = NewContentSecurityPolicy(map[CSPDirective][]CSPSource{
		DirectiveDefaultSrc: {SourceNone},
	})
	helmet.XDNSPrefetchControl = XDNSPrefetchControlOn
	helmet.ExpectCT = NewExpectCT(30, true, "/report-uri")
	helmet.FeaturePolicy = NewFeaturePolicy(map[FeaturePolicyDirective][]FeaturePolicyOrigin{
		DirectiveGeolocation: {OriginSelf, OriginSrc},
	})
	helmet.FrameOptions = FrameOptionsDeny
	helmet.PermittedCrossDomainPolicies = PermittedCrossDomainPoliciesAll
	helmet.XPoweredBy = NewXPoweredBy(false, "PHP 4.2.0")

	addXPoweredByHelmetMiddleware(helmet.Secure(mockNext)).ServeHTTP(rr, r)
	resp := rr.Result()

	testCases := []struct {
		name   string
		header string
	}{
		{HeaderContentSecurityPolicy, "default-src 'none'"},
		{HeaderXDNSPrefetchControl, XDNSPrefetchControlOn.String()},
		{HeaderExpectCT, `max-age=30, enforce, report-uri="/report-uri"`},
		{HeaderFeaturePolicy, "geolocation 'self' 'src'"},
		{HeaderFrameOptions, "DENY"},
		{HeaderPermittedCrossDomainPolicies, PermittedCrossDomainPoliciesAll.String()},
		{HeaderXPoweredBy, "PHP 4.2.0"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			header := resp.Header.Get(tc.name)
			if header != tc.header {
				t.Errorf("Expected: %s\tActual: %s\n", tc.header, header)
			}
		})
	}

	testMockNext(t, resp)
}
