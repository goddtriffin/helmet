package helmet

import (
	"fmt"
	"strings"
	"testing"
)

func TestContentSecurityPolicy_New(t *testing.T) {
	t.Parallel()

	t.Run("Nil", func(t *testing.T) {
		t.Parallel()

		csp := NewContentSecurityPolicy(nil)

		if csp.policies == nil {
			t.Errorf("Policies should not be nil\n")
		}

		if len(csp.policies) != 0 {
			t.Errorf("There should be zero policies/sources\n")
		}

		if csp.cache != "" {
			t.Errorf("Cache should not be set\tActual: %s\n", csp.cache)
		}
	})

	testCases := []struct {
		name     string
		policies map[CSPDirective][]CSPSource
	}{
		{
			name: "Single Directive",
			policies: map[CSPDirective][]CSPSource{
				DirectiveDefaultSrc: {SourceNone},
			},
		},
		{
			name: "Multiple Directives",
			policies: map[CSPDirective][]CSPSource{
				DirectiveDefaultSrc: {SourceNone},
				DirectiveScriptSrc:  {SourceSelf, SourceUnsafeInline},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			csp := NewContentSecurityPolicy(tc.policies)

			if len(csp.policies) != len(tc.policies) {
				t.Errorf("Length doesn't match\tExpected: %d\tActual: %d\n", len(tc.policies), len(csp.policies))
			}

			for policy, sources := range csp.policies {
				expectedSources, ok := tc.policies[policy]
				if !ok {
					t.Errorf("Missing policy\tExpected: %s\n", policy)
				}

				for i, source := range sources {
					if source != expectedSources[i] {
						t.Errorf("Missing source\tExpected: %s\n", source)
					}
				}
			}

			if csp.cache != "" {
				t.Errorf("Cache should not be set\tCache: %s\n", csp.cache)
			}
		})
	}
}

func TestContentSecurityPolicy_Empty(t *testing.T) {
	t.Parallel()

	csp := EmptyContentSecurityPolicy()

	if csp.policies == nil {
		t.Errorf("Policies should not be nil\n")
	}

	if len(csp.policies) != 0 {
		t.Errorf("There should be zero policies/sources\n")
	}

	if csp.cache != "" {
		t.Errorf("Cache should not be set\tActual: %s\n", csp.cache)
	}
}

func TestCSP_Add(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		directive  CSPDirective
		sources    []CSPSource
		expectedOk bool
	}{
		{name: "Empty", directive: "", expectedOk: false},
		{name: "Default Directive", directive: DirectiveDefaultSrc, sources: []CSPSource{SourceNone}, expectedOk: true},
		{name: "No Sources", directive: DirectiveDefaultSrc, sources: []CSPSource{}, expectedOk: true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			csp := EmptyContentSecurityPolicy()
			csp.Add(tc.directive, tc.sources...)

			// make sure directive is now in CSP policies
			if _, ok := csp.policies[tc.directive]; ok != tc.expectedOk {
				t.Errorf("Directive is missing\tDirective: %s\n", tc.directive)
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

func TestCSP_Create(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		directive  CSPDirective
		expectedOk bool
	}{
		{name: "Empty", directive: "", expectedOk: false},
		{name: "Default Directive", directive: DirectiveDefaultSrc, expectedOk: true},
		{name: "Random Directive", directive: "test", expectedOk: true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			csp := EmptyContentSecurityPolicy()
			csp.create(tc.directive)

			// make sure directive is now in CSP policies
			if _, ok := csp.policies[tc.directive]; ok != tc.expectedOk {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedOk, ok)
			}
		})
	}
}

func TestCSP_Remove(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		csp        *ContentSecurityPolicy
		directives []CSPDirective
	}{
		{name: "Empty", csp: EmptyContentSecurityPolicy(), directives: []CSPDirective{}},
		{
			name: "Single Directive",
			csp: NewContentSecurityPolicy(map[CSPDirective][]CSPSource{
				DirectiveDefaultSrc: {SourceNone},
			}),
			directives: []CSPDirective{DirectiveDefaultSrc},
		},
		{
			name: "Multiple Directives",
			csp: NewContentSecurityPolicy(map[CSPDirective][]CSPSource{
				DirectiveDefaultSrc: {SourceNone},
				DirectiveScriptSrc:  {SourceSelf, SourceUnsafeInline},
			}),
			directives: []CSPDirective{DirectiveDefaultSrc, DirectiveScriptSrc},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// make sure directives are in CSP
			for _, directive := range tc.directives {
				if _, ok := tc.csp.policies[directive]; !ok {
					t.Errorf("Directive is missing\tDirective: %s\n", directive)
				}
			}

			tc.csp.Remove(tc.directives...)

			// make sure directives are NOT in CSP
			for _, directive := range tc.directives {
				if _, ok := tc.csp.policies[directive]; ok {
					t.Errorf("Directive should be removed\tDirective: %s\n", directive)
				}
			}
		})
	}
}

func TestCSP_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name             string
		csp              *ContentSecurityPolicy
		expectedPolicies []string
	}{
		{name: "Empty", csp: EmptyContentSecurityPolicy(), expectedPolicies: []string{}},
		{name: "Nil", csp: NewContentSecurityPolicy(nil), expectedPolicies: []string{}},
		{
			name: "Single Directive",
			csp: NewContentSecurityPolicy(map[CSPDirective][]CSPSource{
				DirectiveDefaultSrc: {SourceNone},
			}),
			expectedPolicies: []string{fmt.Sprintf("%s %s", DirectiveDefaultSrc, SourceNone)},
		},
		{
			name: "Single Directive, No Sources",
			csp: NewContentSecurityPolicy(map[CSPDirective][]CSPSource{
				DirectiveUpgradeInsecureRequests: {},
			}),
			expectedPolicies: []string{fmt.Sprintf("%s", DirectiveUpgradeInsecureRequests)},
		},
		{
			name: "Multiple Directives",
			csp: NewContentSecurityPolicy(map[CSPDirective][]CSPSource{
				DirectiveDefaultSrc:              {SourceNone},
				DirectiveUpgradeInsecureRequests: {},
			}),
			expectedPolicies: []string{
				fmt.Sprintf("%s %s", DirectiveDefaultSrc, SourceNone),
				fmt.Sprintf("%s", DirectiveUpgradeInsecureRequests),
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

			// check for the correct amount of semicolons
			semicolonCount := strings.Count(str, ";")
			expectedSemicolonCount := len(tc.expectedPolicies) - 1
			if expectedSemicolonCount < 0 {
				expectedSemicolonCount = 0
			}
			if semicolonCount != expectedSemicolonCount {
				t.Errorf("Incorrect amount of semicolons\tExpected: %d\tActual: %d\n", expectedSemicolonCount, semicolonCount)
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

func TestCSP_Exists(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		csp            *ContentSecurityPolicy
		expectedExists bool
	}{
		{name: "Empty", csp: EmptyContentSecurityPolicy(), expectedExists: false},
		{name: "Nil", csp: NewContentSecurityPolicy(nil), expectedExists: false},
		{
			name: "Single Directive",
			csp: NewContentSecurityPolicy(map[CSPDirective][]CSPSource{
				DirectiveDefaultSrc: {SourceNone},
			}),
			expectedExists: true,
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
