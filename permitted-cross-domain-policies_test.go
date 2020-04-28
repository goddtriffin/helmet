package helmet

import "testing"

func TestPermittedCrossDomainPolicies_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                         string
		permittedCrossDomainPolicies PermittedCrossDomainPolicies
		expectedHeader               string
	}{
		{name: "Empty", permittedCrossDomainPolicies: "", expectedHeader: ""},
		{name: "None", permittedCrossDomainPolicies: PermittedCrossDomainPoliciesNone, expectedHeader: "none"},
		{name: "Master Only", permittedCrossDomainPolicies: PermittedCrossDomainPoliciesMasterOnly, expectedHeader: "master-only"},
		{name: "By Content Type", permittedCrossDomainPolicies: PermittedCrossDomainPoliciesByContentType, expectedHeader: "by-content-type"},
		{name: "By FTP Filename", permittedCrossDomainPolicies: PermittedCrossDomainPoliciesByFTPFilename, expectedHeader: "by-ftp-filename"},
		{name: "All", permittedCrossDomainPolicies: PermittedCrossDomainPoliciesAll, expectedHeader: "all"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			header := tc.permittedCrossDomainPolicies.String()
			if header != tc.expectedHeader {
				t.Errorf("Expected: %s\tActual: %s\n", tc.expectedHeader, header)
			}
		})
	}
}

func TestPermittedCrossDomainPolicies_Exists(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                         string
		permittedCrossDomainPolicies PermittedCrossDomainPolicies
		expectedExists               bool
	}{
		{name: "Empty", permittedCrossDomainPolicies: "", expectedExists: false},
		{name: "None", permittedCrossDomainPolicies: PermittedCrossDomainPoliciesNone, expectedExists: true},
		{name: "Master Only", permittedCrossDomainPolicies: PermittedCrossDomainPoliciesMasterOnly, expectedExists: true},
		{name: "By Content Type", permittedCrossDomainPolicies: PermittedCrossDomainPoliciesByContentType, expectedExists: true},
		{name: "By FTP Filename", permittedCrossDomainPolicies: PermittedCrossDomainPoliciesByFTPFilename, expectedExists: true},
		{name: "All", permittedCrossDomainPolicies: PermittedCrossDomainPoliciesAll, expectedExists: true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			exists := tc.permittedCrossDomainPolicies.Exists()
			if exists != tc.expectedExists {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedExists, exists)
			}
		})
	}
}
