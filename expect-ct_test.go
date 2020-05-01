package helmet

import (
	"testing"
)

func TestExpectCT_DirectiveMaxAge(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name              string
		maxAge            int
		expectedDirective ExpectCTDirective
	}{
		{name: "0", maxAge: 0, expectedDirective: ""},
		{name: "-1", maxAge: -1, expectedDirective: ""},
		{name: "1", maxAge: 1, expectedDirective: "max-age=1"},
		{name: "1 day", maxAge: 86400, expectedDirective: "max-age=86400"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			directive := ExpectCTDirectiveMaxAge(tc.maxAge)
			if directive != tc.expectedDirective {
				t.Errorf("Incorrect MaxAge\tExpected: %s\tActual: %s\n", tc.expectedDirective, directive)
			}
		})
	}
}

func TestExpectCT_DirectiveReportURI(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name              string
		reportURI         string
		expectedDirective ExpectCTDirective
	}{
		{name: "Empty", reportURI: "", expectedDirective: ""},
		{name: "Not empty", reportURI: "/report-uri", expectedDirective: `report-uri="/report-uri"`},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			directive := ExpectCTDirectiveReportURI(tc.reportURI)
			if directive != tc.expectedDirective {
				t.Errorf("Incorrect ReportURI\tExpected: %s\tActual: %s\n", tc.expectedDirective, directive)
			}
		})
	}
}

func TestExpectCT_New(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		maxAge    int
		enforce   bool
		reportURI string
	}{
		{name: "Empty", maxAge: 0, enforce: false, reportURI: ""},
		{name: "Max Age", maxAge: 30, enforce: false, reportURI: ""},
		{name: "Enforce", maxAge: 0, enforce: true, reportURI: ""},
		{name: "ReportURI", maxAge: 0, enforce: false, reportURI: "/report-uri"},
		{name: "Max Age, Enforce", maxAge: 30, enforce: true, reportURI: ""},
		{name: "Max Age, ReportURI", maxAge: 30, enforce: false, reportURI: "/report-uri"},
		{name: "Enforce, ReportURI", maxAge: 0, enforce: true, reportURI: "/report-uri"},
		{name: "Max Age, Enforce, ReportURI", maxAge: 30, enforce: true, reportURI: "/report-uri"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ect := NewExpectCT(tc.maxAge, tc.enforce, tc.reportURI)

			if ect.MaxAge != tc.maxAge {
				t.Errorf("Incorrect MaxAge\tExpected: %d\tActual: %d\n", tc.maxAge, ect.MaxAge)
			}

			if ect.Enforce != tc.enforce {
				t.Errorf("Incorrect Enforce\tExpected: %t\tActual: %t\n", tc.enforce, ect.Enforce)
			}

			if ect.ReportURI != tc.reportURI {
				t.Errorf("Incorrect ReportURI\tExpected: %s\tActual: %s\n", tc.reportURI, ect.ReportURI)
			}

			if ect.cache != "" {
				t.Errorf("Cache should not be set\tActual: %s\n", ect.cache)
			}
		})
	}
}

func TestExpectCT_Empty(t *testing.T) {
	t.Parallel()

	ect := EmptyExpectCT()

	if ect.MaxAge != 0 {
		t.Errorf("MaxAge should be zero\tActual: %d\n", ect.MaxAge)
	}

	if ect.Enforce != false {
		t.Errorf("Enforce should be false\tActual: %t\n", ect.Enforce)
	}

	if ect.ReportURI != "" {
		t.Errorf("ReportURI should be empty\tActual: %s\n", ect.ReportURI)
	}

	if ect.cache != "" {
		t.Errorf("Cache should not be set\tActual: %s\n", ect.cache)
	}
}

func TestExpectCT_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		expectCT       *ExpectCT
		expectedHeader string
	}{
		{name: "Empty", expectCT: EmptyExpectCT(), expectedHeader: ""},
		{name: "Max Age Zero", expectCT: NewExpectCT(0, false, ""), expectedHeader: ""},
		{name: "Max Age", expectCT: NewExpectCT(30, false, ""), expectedHeader: "max-age=30"},
		{name: "Max Age, Enforce", expectCT: NewExpectCT(30, true, ""), expectedHeader: "max-age=30, enforce"},
		{name: "Max Age, Report URI", expectCT: NewExpectCT(30, false, "/report-uri"), expectedHeader: `max-age=30, report-uri="/report-uri"`},
		{name: "Max Age, Enforce, Report URI", expectCT: NewExpectCT(30, true, "/report-uri"), expectedHeader: `max-age=30, enforce, report-uri="/report-uri"`},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			header := tc.expectCT.String()
			if header != tc.expectedHeader {
				t.Errorf("Expected: %s\tActual: %s\n", tc.expectedHeader, header)
			}

			// check that the cache is set
			if tc.expectCT.cache != header {
				t.Errorf("ExpectCT String() cache is not set!\tActual: %s\n", tc.expectCT.cache)
			}

			// utilize said cache
			header = tc.expectCT.String()
			if header != tc.expectedHeader {
				t.Errorf("Expected: %s\tActual: %s\n", tc.expectedHeader, header)
			}
		})
	}
}

func TestExpectCT_Exists(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		expectCT       *ExpectCT
		expectedExists bool
	}{
		{name: "Empty", expectCT: EmptyExpectCT(), expectedExists: false},
		{name: "Max Age Zero", expectCT: NewExpectCT(0, false, ""), expectedExists: false},
		{name: "Max Age Set", expectCT: NewExpectCT(30, false, ""), expectedExists: true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			exists := tc.expectCT.Exists()
			if exists != tc.expectedExists {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedExists, exists)
			}
		})
	}
}
