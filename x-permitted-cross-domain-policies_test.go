package helmet

import "testing"

func TestXPermittedCrossDomainPolicies_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                          string
		xPermittedCrossDomainPolicies XPermittedCrossDomainPolicies
		expectedHeader                string
	}{
		{name: "Empty", xPermittedCrossDomainPolicies: "", expectedHeader: ""},
		{name: "None", xPermittedCrossDomainPolicies: PermittedCrossDomainPoliciesNone, expectedHeader: "none"},
		{name: "Master Only", xPermittedCrossDomainPolicies: PermittedCrossDomainPoliciesMasterOnly, expectedHeader: "master-only"},
		{name: "By Content Type", xPermittedCrossDomainPolicies: PermittedCrossDomainPoliciesByContentType, expectedHeader: "by-content-type"},
		{name: "By FTP Filename", xPermittedCrossDomainPolicies: PermittedCrossDomainPoliciesByFTPFilename, expectedHeader: "by-ftp-filename"},
		{name: "All", xPermittedCrossDomainPolicies: PermittedCrossDomainPoliciesAll, expectedHeader: "all"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			header := tc.xPermittedCrossDomainPolicies.String()
			if header != tc.expectedHeader {
				t.Errorf("Expected: %s\tActual: %s\n", tc.expectedHeader, header)
			}
		})
	}
}

func TestXPermittedCrossDomainPolicies_Exists(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                          string
		xPermittedCrossDomainPolicies XPermittedCrossDomainPolicies
		expectedExists                bool
	}{
		{name: "Empty", xPermittedCrossDomainPolicies: "", expectedExists: false},
		{name: "None", xPermittedCrossDomainPolicies: PermittedCrossDomainPoliciesNone, expectedExists: true},
		{name: "Master Only", xPermittedCrossDomainPolicies: PermittedCrossDomainPoliciesMasterOnly, expectedExists: true},
		{name: "By Content Type", xPermittedCrossDomainPolicies: PermittedCrossDomainPoliciesByContentType, expectedExists: true},
		{name: "By FTP Filename", xPermittedCrossDomainPolicies: PermittedCrossDomainPoliciesByFTPFilename, expectedExists: true},
		{name: "All", xPermittedCrossDomainPolicies: PermittedCrossDomainPoliciesAll, expectedExists: true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			exists := tc.xPermittedCrossDomainPolicies.Exists()
			if exists != tc.expectedExists {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedExists, exists)
			}
		})
	}
}
