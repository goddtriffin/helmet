package helmet

import "testing"

func TestStrictTransportSecurity_New(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name              string
		maxAge            int
		includeSubDomains bool
		preload           bool
	}{
		{name: "Empty", maxAge: 0, includeSubDomains: false, preload: false},
		{name: "Max Age", maxAge: 63072000, includeSubDomains: false, preload: false},
		{name: "Include Sub Domains", maxAge: 0, includeSubDomains: true, preload: false},
		{name: "Preload", maxAge: 0, includeSubDomains: false, preload: true},
		{name: "Max Age, Include Sub Domains", maxAge: 63072000, includeSubDomains: true, preload: false},
		{name: "Max Age, Preload", maxAge: 63072000, includeSubDomains: false, preload: true},
		{name: "Include Sub Domains, Preload", maxAge: 0, includeSubDomains: true, preload: true},
		{name: "Max Age, Include Sub Domains, Preload", maxAge: 63072000, includeSubDomains: true, preload: true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			hsts := NewStrictTransportSecurity(tc.maxAge, tc.includeSubDomains, tc.preload)

			if hsts.MaxAge != tc.maxAge {
				t.Errorf("Incorrect MaxAge\tExpected: %d\tActual: %d\n", tc.maxAge, hsts.MaxAge)
			}

			if hsts.IncludeSubDomains != tc.includeSubDomains {
				t.Errorf("Incorrect IncludeSubDomains\tExpected: %t\tActual: %t\n", tc.includeSubDomains, hsts.IncludeSubDomains)
			}

			if hsts.Preload != tc.preload {
				t.Errorf("Incorrect ReportURI\tExpected: %t\tActual: %t\n", tc.preload, hsts.Preload)
			}

			if hsts.cache != "" {
				t.Errorf("Cache should not be set\tActual: %s\n", hsts.cache)
			}
		})
	}
}

func TestStrictTransportSecurity_Empty(t *testing.T) {
	t.Parallel()

	hsts := EmptyStrictTransportSecurity()

	if hsts.MaxAge != 0 {
		t.Errorf("MaxAge should be zero\tActual: %d\n", hsts.MaxAge)
	}

	if hsts.IncludeSubDomains != false {
		t.Errorf("IncludeSubDomains should be false\tActual: %t\n", hsts.IncludeSubDomains)
	}

	if hsts.Preload != false {
		t.Errorf("Preload should be false\tActual: %t\n", hsts.Preload)
	}

	if hsts.cache != "" {
		t.Errorf("Cache should not be set\tActual: %s\n", hsts.cache)
	}
}
