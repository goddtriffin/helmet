package helmet

import (
	"fmt"
	"strings"
	"testing"
)

func TestFeaturePolicy_New(t *testing.T) {
	t.Parallel()

	t.Run("Nil", func(t *testing.T) {
		t.Parallel()

		fp := NewFeaturePolicy(nil)

		if fp.policies == nil {
			t.Errorf("Policies should not be nil\n")
		}

		if len(fp.policies) != 0 {
			t.Errorf("There should be zero policies/sources\n")
		}

		if fp.cache != "" {
			t.Errorf("Cache should not be set\tActual: %s\n", fp.cache)
		}
	})

	testCases := []struct {
		name     string
		policies map[FeaturePolicyDirective][]FeaturePolicyOrigin
	}{
		{
			name: "Single Directive",
			policies: map[FeaturePolicyDirective][]FeaturePolicyOrigin{
				DirectiveMicrophone: {OriginNone},
			},
		},
		{
			name: "Multiple Directives",
			policies: map[FeaturePolicyDirective][]FeaturePolicyOrigin{
				DirectiveMicrophone:  {OriginNone},
				DirectiveGeolocation: {OriginSelf, OriginSrc},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			fp := NewFeaturePolicy(tc.policies)

			if len(fp.policies) != len(tc.policies) {
				t.Errorf("Length doesn't match\tExpected: %d\tActual: %d\n", len(tc.policies), len(fp.policies))
			}

			for policy, origins := range fp.policies {
				expectedOrigins, ok := tc.policies[policy]
				if !ok {
					t.Errorf("Missing policy\tExpected: %s\n", policy)
				}

				for i, origin := range origins {
					if origin != expectedOrigins[i] {
						t.Errorf("Missing origin\tExpected: %s\n", origin)
					}
				}
			}

			if fp.cache != "" {
				t.Errorf("Cache should not be set\tCache: %s\n", fp.cache)
			}
		})
	}
}

func TestFeaturePolicy_Empty(t *testing.T) {
	t.Parallel()

	fp := EmptyFeaturePolicy()

	if fp.policies == nil {
		t.Errorf("Policies should not be nil\n")
	}

	if len(fp.policies) != 0 {
		t.Errorf("There should be zero policies/origins\n")
	}

	if fp.cache != "" {
		t.Errorf("Cache should not be set\tActual: %s\n", fp.cache)
	}
}

func TestFeaturePolicy_Add(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		directive  FeaturePolicyDirective
		origins    []FeaturePolicyOrigin
		expectedOk bool
	}{
		{name: "Empty", directive: "", expectedOk: false},
		{name: "Directive with Origin", directive: DirectiveMicrophone, origins: []FeaturePolicyOrigin{OriginNone}, expectedOk: true},
		{name: "No Origins", directive: DirectiveGeolocation, origins: []FeaturePolicyOrigin{}, expectedOk: false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			fp := EmptyFeaturePolicy()
			fp.Add(tc.directive, tc.origins...)

			// make sure directive is now in CSP policies
			if _, ok := fp.policies[tc.directive]; ok != tc.expectedOk {
				t.Errorf("Directive is missing\tDirective: %s\n", tc.directive)
			}

			// next part requires there to be sources
			if tc.expectedOk == false {
				return
			}

			// make sure the origins were added correctly
			for i, origin := range fp.policies[tc.directive] {
				if origin != tc.origins[i] {
					t.Errorf("Index: %d\tExpected: %s\tActual: %s\n", i, tc.origins[i], origin)
				}
			}
		})
	}
}

func TestFeaturePolicy_Create(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		directive  FeaturePolicyDirective
		expectedOk bool
	}{
		{name: "Empty", directive: "", expectedOk: false},
		{name: "Directive", directive: DirectiveMicrophone, expectedOk: true},
		{name: "Random Directive", directive: "test", expectedOk: true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			fp := EmptyFeaturePolicy()
			fp.create(tc.directive)

			// make sure directive is now in FeaturePolicy policies
			if _, ok := fp.policies[tc.directive]; ok != tc.expectedOk {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedOk, ok)
			}
		})
	}
}

func TestFeaturePolicy_Remove(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name       string
		fp         *FeaturePolicy
		directives []FeaturePolicyDirective
	}{
		{name: "Empty", fp: EmptyFeaturePolicy(), directives: []FeaturePolicyDirective{}},
		{
			name: "Single Directive",
			fp: NewFeaturePolicy(map[FeaturePolicyDirective][]FeaturePolicyOrigin{
				DirectiveMicrophone: {OriginNone},
			}),
			directives: []FeaturePolicyDirective{DirectiveMicrophone},
		},
		{
			name: "Multiple Directives",
			fp: NewFeaturePolicy(map[FeaturePolicyDirective][]FeaturePolicyOrigin{
				DirectiveMicrophone:  {OriginNone},
				DirectiveGeolocation: {OriginSelf, OriginSrc},
			}),
			directives: []FeaturePolicyDirective{DirectiveMicrophone, DirectiveGeolocation},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// make sure directives are in FeaturePolicy
			for _, directive := range tc.directives {
				if _, ok := tc.fp.policies[directive]; !ok {
					t.Errorf("Directive is missing\tDirective: %s\n", directive)
				}
			}

			tc.fp.Remove(tc.directives...)

			// make sure directives are NOT in FeaturePolicy
			for _, directive := range tc.directives {
				if _, ok := tc.fp.policies[directive]; ok {
					t.Errorf("Directive should be removed\tDirective: %s\n", directive)
				}
			}
		})
	}
}

func TestFeaturePolicy_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name             string
		fp               *FeaturePolicy
		expectedPolicies []string
	}{
		{name: "Empty", fp: EmptyFeaturePolicy(), expectedPolicies: []string{}},
		{name: "Nil", fp: NewFeaturePolicy(nil), expectedPolicies: []string{}},
		{
			name: "Single Directive",
			fp: NewFeaturePolicy(map[FeaturePolicyDirective][]FeaturePolicyOrigin{
				DirectiveMicrophone: {OriginNone},
			}),
			expectedPolicies: []string{fmt.Sprintf("%s %s;", DirectiveMicrophone, OriginNone)},
		},
		{
			name: "Multiple Directives",
			fp: NewFeaturePolicy(map[FeaturePolicyDirective][]FeaturePolicyOrigin{
				DirectiveMicrophone:  {OriginNone},
				DirectiveGeolocation: {OriginSelf, OriginSrc},
			}),
			expectedPolicies: []string{
				fmt.Sprintf("%s %s;", DirectiveMicrophone, OriginNone),
				fmt.Sprintf("%s %s %s;", DirectiveGeolocation, OriginSelf, OriginSrc),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// check that the FeaturePolicy contains all policies
			str := tc.fp.String()
			for _, policy := range tc.expectedPolicies {
				if !strings.Contains(str, policy) {
					t.Errorf("Policy is missing\tExpected: %s\tActual: %s\n", policy, str)
				}
			}

			// check that cache is set
			if len(tc.expectedPolicies) > 0 && len(tc.fp.cache) == 0 {
				t.Errorf("FeaturePolicy String() cache is not set\tActual: %s\n", tc.fp.cache)
			}

			// utilize said cache
			str = tc.fp.String()
			for _, policy := range tc.expectedPolicies {
				if !strings.Contains(str, policy) {
					t.Errorf("Cache policy is missing\tExpected: %s\tActual: %s\n", policy, str)
				}
			}
		})
	}
}
