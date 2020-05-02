package helmet

import (
	"testing"
)

func TestReferrerPolicy_NewReferrerPolicy(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		directives []ReferrerPolicyDirective
	}{
		{name: "Zero Directives", directives: []ReferrerPolicyDirective{}},
		{name: "Single Directive", directives: []ReferrerPolicyDirective{DirectiveNoReferrer}},
		{name: "Multiple Directives", directives: []ReferrerPolicyDirective{DirectiveNoReferrer, DirectiveStrictOriginWhenCrossOrigin}},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rp := NewReferrerPolicy(tc.directives...)

			if len(rp.policies) != len(tc.directives) {
				t.Errorf("Length doesn't match\tExpected: %d\tActual: %d\n", len(tc.directives), len(rp.policies))
			}

			for i, directive := range rp.policies {
				if directive != tc.directives[i] {
					t.Errorf("Missing directive\tIndex: %d\tExpected: %s\n", i, directive)
				}
			}

			if rp.cache != "" {
				t.Errorf("Cache should not be set\tCache: %s\n", rp.cache)
			}
		})
	}
}

func TestReferrerPolicy_EmptyReferrerPolicy(t *testing.T) {
	t.Parallel()

	rp := EmptyReferrerPolicy()

	if rp.policies == nil {
		t.Errorf("Policies should not be nil\n")
	}

	if len(rp.policies) != 0 {
		t.Errorf("There should be zero directives\n")
	}

	if rp.cache != "" {
		t.Errorf("Cache should not be set\tActual: %s\n", rp.cache)
	}
}

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

func TestReferrer_Empty(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		referrerPolicy *ReferrerPolicy
		expectedEmpty  bool
	}{
		{name: "Empty", referrerPolicy: EmptyReferrerPolicy(), expectedEmpty: true},
		{
			name:           "Single Directive",
			referrerPolicy: NewReferrerPolicy(DirectiveNoReferrer),
			expectedEmpty:  false,
		},
		{
			name:           "Multiple Directives",
			referrerPolicy: NewReferrerPolicy(DirectiveNoReferrer, DirectiveStrictOriginWhenCrossOrigin),
			expectedEmpty:  false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			exists := tc.referrerPolicy.Empty()
			if exists != tc.expectedEmpty {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedEmpty, exists)
			}
		})
	}
}
