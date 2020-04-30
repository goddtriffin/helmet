package helmet

import "testing"

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

			for policy, sources := range fp.policies {
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
		t.Errorf("There should be zero policies/sources\n")
	}

	if fp.cache != "" {
		t.Errorf("Cache should not be set\tActual: %s\n", fp.cache)
	}
}
