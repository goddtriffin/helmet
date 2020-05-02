package helmet

import (
	"testing"
)

func TestReferrerPolicy_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		referrerPolicy *ReferrerPolicy
		expectedHeader string
	}{
		{name: "Empty", referrerPolicy: EmptyReferrerPolicy(), expectedHeader: ""},
		{
			name:           "Single Directive",
			referrerPolicy: NewReferrerPolicy(DirectiveNoReferrer),
			expectedHeader: "no-referrer",
		},
		{
			name:           "Multiple Directives",
			referrerPolicy: NewReferrerPolicy(DirectiveNoReferrer, DirectiveStrictOriginWhenCrossOrigin),
			expectedHeader: "no-referrer, strict-origin-when-cross-origin",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			header := tc.referrerPolicy.String()
			if header != tc.expectedHeader {
				t.Errorf("Expected: %s\tActual: %s\n", tc.expectedHeader, header)
			}

			// check that the cache is set
			if tc.referrerPolicy.cache != header {
				t.Errorf("ExpectCT String() cache is not set!\tActual: %s\n", tc.referrerPolicy.cache)
			}

			// utilize said cache
			header = tc.referrerPolicy.String()
			if header != tc.expectedHeader {
				t.Errorf("Expected: %s\tActual: %s\n", tc.expectedHeader, header)
			}
		})
	}
}

func TestReferrer_Exists(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		referrerPolicy *ReferrerPolicy
		expectedExists bool
	}{
		{name: "Empty", referrerPolicy: EmptyReferrerPolicy(), expectedExists: false},
		{
			name:           "Single Directive",
			referrerPolicy: NewReferrerPolicy(DirectiveNoReferrer),
			expectedExists: true,
		},
		{
			name:           "Multiple Directives",
			referrerPolicy: NewReferrerPolicy(DirectiveNoReferrer, DirectiveStrictOriginWhenCrossOrigin),
			expectedExists: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			exists := tc.referrerPolicy.Exists()
			if exists != tc.expectedExists {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedExists, exists)
			}
		})
	}
}
