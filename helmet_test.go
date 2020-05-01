package helmet

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelmet_Secure_default(t *testing.T) {
	t.Parallel()

	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockNext := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// default Helmet
	helmet := Default()
	helmet.Secure(mockNext).ServeHTTP(rr, r)
	resp := rr.Result()

	testCases := []struct {
		name   string
		header string
	}{
		{HeaderContentSecurityPolicy, ""},
		{HeaderDNSPrefetchControl, DNSPrefetchControlOff.String()},
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

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected: %d\tActual: %d\n", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	body := string(buf)
	if body != "OK" {
		t.Errorf("Expected: %s\tActual: %s\n", "OK", body)
	}
}

func TestHelmet_Secure_empty(t *testing.T) {
	t.Parallel()

	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockNext := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// blank slate Helmet
	helmet := Empty()
	helmet.Secure(mockNext).ServeHTTP(rr, r)
	resp := rr.Result()

	testCases := []struct {
		header string
	}{
		{HeaderContentSecurityPolicy},
		{HeaderDNSPrefetchControl},
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

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected: %d\tActual: %d\n", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	body := string(buf)
	if body != "OK" {
		t.Errorf("Expected: %s\tActual: %s\n", "OK", body)
	}
}

func TestHelmet_Secure_custom(t *testing.T) {
	t.Parallel()

	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockNext := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// fill Helmet with custom parameters
	helmet := Empty()
	helmet.ContentSecurityPolicy = NewContentSecurityPolicy(map[CSPDirective][]CSPSource{
		DirectiveDefaultSrc: {SourceNone},
	})
	helmet.DNSPrefetchControl = DNSPrefetchControlOn
	helmet.ExpectCT = NewExpectCT(30, true, "/report-uri")
	helmet.FeaturePolicy = NewFeaturePolicy(map[FeaturePolicyDirective][]FeaturePolicyOrigin{
		DirectiveGeolocation: {OriginSelf, OriginSrc},
	})
	helmet.FrameOptions = FrameOptionsDeny
	helmet.PermittedCrossDomainPolicies = PermittedCrossDomainPoliciesAll

	helmet.Secure(mockNext).ServeHTTP(rr, r)
	resp := rr.Result()

	testCases := []struct {
		name   string
		header string
	}{
		{HeaderContentSecurityPolicy, "default-src 'none'"},
		{HeaderDNSPrefetchControl, DNSPrefetchControlOn.String()},
		{HeaderExpectCT, `max-age=30, enforce, report-uri="/report-uri"`},
		{HeaderFeaturePolicy, "geolocation 'self' 'src'"},
		{HeaderFrameOptions, "DENY"},
		{HeaderPermittedCrossDomainPolicies, PermittedCrossDomainPoliciesAll.String()},
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

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected: %d\tActual: %d\n", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	body := string(buf)
	if body != "OK" {
		t.Errorf("Expected: %s\tActual: %s\n", "OK", body)
	}
}
