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

func TestXPermittedCrossDomainPolicies_Empty(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                          string
		xPermittedCrossDomainPolicies XPermittedCrossDomainPolicies
		expectedEmpty                 bool
	}{
		{name: "Empty", xPermittedCrossDomainPolicies: "", expectedEmpty: true},
		{name: "None", xPermittedCrossDomainPolicies: PermittedCrossDomainPoliciesNone, expectedEmpty: false},
		{name: "Master Only", xPermittedCrossDomainPolicies: PermittedCrossDomainPoliciesMasterOnly, expectedEmpty: false},
		{name: "By Content Type", xPermittedCrossDomainPolicies: PermittedCrossDomainPoliciesByContentType, expectedEmpty: false},
		{name: "By FTP Filename", xPermittedCrossDomainPolicies: PermittedCrossDomainPoliciesByFTPFilename, expectedEmpty: false},
		{name: "All", xPermittedCrossDomainPolicies: PermittedCrossDomainPoliciesAll, expectedEmpty: false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			exists := tc.xPermittedCrossDomainPolicies.Empty()
			if exists != tc.expectedEmpty {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedEmpty, exists)
			}
		})
	}
}
