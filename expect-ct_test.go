package helmet

import (
	"testing"
)

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
