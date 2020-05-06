package helmet

import "testing"

func TestXXSSProtection_DirectiveXSSFiltering(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name              string
		xssFiltering      bool
		expectedDirective XXSSProtectionDirective
	}{
		{name: "False", xssFiltering: false, expectedDirective: "0"},
		{name: "True", xssFiltering: true, expectedDirective: "1"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			directive := XXSSProtectionDirectiveXSSFiltering(tc.xssFiltering)
			if directive != tc.expectedDirective {
				t.Errorf("Incorrect XSSFiltering\tExpected: %s\tActual: %s\n", tc.expectedDirective, directive)
			}
		})
	}
}

func TestXXSSProtection_DirectiveReportURI(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name              string
		reportURI         string
		expectedDirective XXSSProtectionDirective
	}{
		{name: "Empty", reportURI: "", expectedDirective: ""},
		{name: "Not empty", reportURI: "/report-uri", expectedDirective: "report=/report-uri"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			directive := XXSSProtectionDirectiveReportURI(tc.reportURI)
			if directive != tc.expectedDirective {
				t.Errorf("Incorrect ReportURI\tExpected: %s\tActual: %s\n", tc.expectedDirective, directive)
			}
		})
	}
}

func TestXXSSProtection_NewXXSSProtection(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		xssFiltering bool
		mode         XXSSProtectionDirective
		reportURI    string
	}{
		{name: "Empty", xssFiltering: false, mode: "", reportURI: ""},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ect := NewXXSSProtection(tc.xssFiltering, tc.mode, tc.reportURI)

			if ect.XSSFiltering != tc.xssFiltering {
				t.Errorf("Incorrect XSSFiltering\tExpected: %t\tActual: %t\n", tc.xssFiltering, ect.XSSFiltering)
			}

			if ect.Mode != tc.mode {
				t.Errorf("Incorrect Mode\tExpected: %s\tActual: %s\n", tc.mode, ect.Mode)
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

func TestXXSSProtection_EmptyXXSSProtection(t *testing.T) {
	t.Parallel()

	xssp := EmptyXXSSProtection()

	if xssp.XSSFiltering != false {
		t.Errorf("XSSFiltering should be false\tActual: %t\n", xssp.XSSFiltering)
	}

	if xssp.Mode != "" {
		t.Errorf("Mode should be blank\tActual: %s\n", xssp.Mode)
	}

	if xssp.ReportURI != "" {
		t.Errorf("ReportURI should be empty\tActual: %s\n", xssp.ReportURI)
	}

	if xssp.cache != "" {
		t.Errorf("Cache should not be set\tActual: %s\n", xssp.cache)
	}
}

func TestXXSSProtection_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		xXSSProtection *XXSSProtection
		expectedHeader string
	}{
		{name: "Empty", xXSSProtection: EmptyXXSSProtection(), expectedHeader: "0"},
		{name: "XSS Filtering", xXSSProtection: NewXXSSProtection(true, "", ""), expectedHeader: "1"},
		{
			name:           "XSS Filtering, Mode Block",
			xXSSProtection: NewXXSSProtection(true, DirectiveModeBlock, ""),
			expectedHeader: "1; mode=block",
		},
		{
			name:           "XSS Filtering, Report URI",
			xXSSProtection: NewXXSSProtection(true, "", "/report-uri"),
			expectedHeader: "1; report=/report-uri",
		},
		{
			name:           "XSS Filtering, Mode Block, Report URI",
			xXSSProtection: NewXXSSProtection(true, DirectiveModeBlock, "/report-uri"),
			expectedHeader: "1; mode=block; report=/report-uri",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			header := tc.xXSSProtection.String()
			if header != tc.expectedHeader {
				t.Errorf("Expected: %s\tActual: %s\n", tc.expectedHeader, header)
			}

			// check that the cache is set
			if tc.xXSSProtection.cache != header {
				t.Errorf("ExpectCT String() cache is not set!\tActual: %s\n", tc.xXSSProtection.cache)
			}

			// utilize said cache
			header = tc.xXSSProtection.String()
			if header != tc.expectedHeader {
				t.Errorf("Expected: %s\tActual: %s\n", tc.expectedHeader, header)
			}
		})
	}
}

func TestXXSSProtection_Empty(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		xXSSProtection *XXSSProtection
		expectedEmpty  bool
	}{
		{name: "XSS Filtering False", xXSSProtection: EmptyXXSSProtection(), expectedEmpty: false},
		{name: "XSS Filtering True", xXSSProtection: NewXXSSProtection(true, "", ""), expectedEmpty: false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			exists := tc.xXSSProtection.Empty()
			if exists != tc.expectedEmpty {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedEmpty, exists)
			}
		})
	}
}
