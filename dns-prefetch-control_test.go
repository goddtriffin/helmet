package helmet

import "testing"

func TestDNSPrefetchControl_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name               string
		dnsPrefetchControl DNSPrefetchControl
		expectedHeader     string
	}{
		{name: "Empty", dnsPrefetchControl: "", expectedHeader: ""},
		{name: "On", dnsPrefetchControl: DNSPrefetchControlOn, expectedHeader: "on"},
		{name: "Off", dnsPrefetchControl: DNSPrefetchControlOff, expectedHeader: "off"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			header := tc.dnsPrefetchControl.String()
			if header != tc.expectedHeader {
				t.Errorf("Expected: %s\tActual: %s\n", tc.expectedHeader, header)
			}
		})
	}
}

func TestDNSPrefetchControl_Exists(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name               string
		dnsPrefetchControl DNSPrefetchControl
		expectedExists     bool
	}{
		{name: "Empty", dnsPrefetchControl: "", expectedExists: false},
		{name: "On", dnsPrefetchControl: DNSPrefetchControlOn, expectedExists: true},
		{name: "Off", dnsPrefetchControl: DNSPrefetchControlOff, expectedExists: true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			exists := tc.dnsPrefetchControl.Exists()
			if exists != tc.expectedExists {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedExists, exists)
			}
		})
	}
}