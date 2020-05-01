package helmet

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mockNext = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
})

func addXPoweredByHelmetMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(HeaderXPoweredBy, "Helmet")

		next.ServeHTTP(w, r)
	})
}

func newRecorderRequest(t *testing.T) (*httptest.ResponseRecorder, *http.Request) {
	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	return rr, r
}

func testMockNext(t *testing.T, resp *http.Response) {
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
