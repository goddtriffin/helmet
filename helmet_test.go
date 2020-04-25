package helmet

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecure_new(t *testing.T) {
	t.Parallel()

	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// mock HTTP handler that we can pass to our secureHeaders middleware
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	helmet := New()
	helmet.ContentSecurityPolicy.Add(DirectiveDefaultSrc, SourceNone)

	helmet.Secure(next).ServeHTTP(rr, r)
	resp := rr.Result()

	testCases := []struct {
		name   string
		header string
	}{
		{"Content-Security-Policy", "default-src 'none';"},
		{"X-DNS-Prefetch-Control", "on"},
		{"X-Permitted-Cross-Domain-Policies", "all"},
	}

	// test headers
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

func TestSecure_default(t *testing.T) {
	t.Parallel()

	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// mock HTTP handler that we can pass to our secureHeaders middleware
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	helmet := Default()
	helmet.ContentSecurityPolicy.Add(DirectiveDefaultSrc, SourceNone)

	helmet.Secure(next).ServeHTTP(rr, r)
	resp := rr.Result()

	testCases := []struct {
		name   string
		header string
	}{
		{"Content-Security-Policy", "default-src 'none';"},
		{"X-DNS-Prefetch-Control", "off"},
		{"X-Permitted-Cross-Domain-Policies", "none"},
	}

	// test headers
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
