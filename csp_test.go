package helmet

import (
	"fmt"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		directive  string
		sources    []string
		expectedOk bool
	}{
		{name: DirectiveDefaultSrc, directive: DirectiveDefaultSrc, sources: []string{SourceNone}, expectedOk: true},
		{name: "Empty", directive: "", expectedOk: false},
		{name: "No Sources", directive: DirectiveDefaultSrc, sources: []string{}, expectedOk: true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			csp := NewContentSecurityPolicy(nil)
			csp.Add(tc.directive, tc.sources...)

			// make sure directive is now in CSP policies
			if _, ok := csp.policies[tc.directive]; ok != tc.expectedOk {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedOk, ok)
			}

			// next part requires there to be sources
			if tc.expectedOk == false {
				return
			}

			// make sure the sources were added correctly
			for i, source := range csp.policies[tc.directive] {
				if source != tc.sources[i] {
					t.Errorf("Index: %d\tExpected: %s\tActual: %s\n", i, tc.sources[i], source)
				}
			}
		})
	}
}

func TestCreate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		directive  string
		expectedOk bool
	}{
		{name: DirectiveDefaultSrc, directive: DirectiveDefaultSrc, expectedOk: true},
		{name: "Empty", directive: "", expectedOk: false},
		{name: "Random Directive", directive: "test", expectedOk: true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			csp := NewContentSecurityPolicy(nil)
			csp.create(tc.directive)

			// make sure directive is now in CSP policies
			if _, ok := csp.policies[tc.directive]; ok != tc.expectedOk {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedOk, ok)
			}
		})
	}
}

func TestString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name             string
		csp              *ContentSecurityPolicy
		expectedPolicies []string
	}{
		{
			name: DirectiveDefaultSrc,
			csp: NewContentSecurityPolicy(map[string][]string{
				DirectiveDefaultSrc: {SourceNone},
			}),
			expectedPolicies: []string{fmt.Sprintf("%s %s;", DirectiveDefaultSrc, SourceNone)},
		},
		{
			name:             "Empty",
			csp:              NewContentSecurityPolicy(make(map[string][]string)),
			expectedPolicies: []string{},
		},
		{
			name:             "Nil",
			csp:              NewContentSecurityPolicy(nil),
			expectedPolicies: []string{},
		},
		{
			name: "No Sources",
			csp: NewContentSecurityPolicy(map[string][]string{
				DirectiveUpgradeInsecureRequests: {},
			}),
			expectedPolicies: []string{fmt.Sprintf("%s;", DirectiveUpgradeInsecureRequests)},
		},
		{
			name: "Multiple Directives",
			csp: NewContentSecurityPolicy(map[string][]string{
				DirectiveDefaultSrc:              {SourceNone},
				DirectiveUpgradeInsecureRequests: {},
			}),
			expectedPolicies: []string{
				fmt.Sprintf("%s %s;", DirectiveDefaultSrc, SourceNone),
				fmt.Sprintf("%s;", DirectiveUpgradeInsecureRequests),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// check that the CSP contains all policies
			str := tc.csp.String()
			for _, policy := range tc.expectedPolicies {
				if !strings.Contains(str, policy) {
					t.Errorf("CSP doesn't contain policy\tExpected: %s\tActual: %s\n", policy, str)
				}
			}

			// check that cache is set
			if len(tc.expectedPolicies) > 0 && len(tc.csp.cache) == 0 {
				t.Errorf("CSP String() cache is not set!\tActual: %s\n", tc.csp.cache)
			}

			// utilize said cache
			str = tc.csp.String()
			for _, policy := range tc.expectedPolicies {
				if !strings.Contains(str, policy) {
					t.Errorf("CSP doesn't contain policy\tExpected: %s\tActual: %s\n", policy, str)
				}
			}
		})
	}
}

func TestExists(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		csp            *ContentSecurityPolicy
		expectedExists bool
	}{
		{
			name: DirectiveDefaultSrc,
			csp: NewContentSecurityPolicy(map[string][]string{
				DirectiveDefaultSrc: {SourceNone},
			}),
			expectedExists: true,
		},
		{
			name:           "Empty",
			csp:            NewContentSecurityPolicy(make(map[string][]string)),
			expectedExists: false,
		},
		{
			name:           "Nil",
			csp:            NewContentSecurityPolicy(nil),
			expectedExists: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			exists := tc.csp.Exists()
			if exists != tc.expectedExists {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedExists, exists)
			}
		})
	}
}
