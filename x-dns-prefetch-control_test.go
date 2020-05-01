package helmet

import "testing"

func TestXDNSPrefetchControl_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                string
		xDNSPrefetchControl XDNSPrefetchControl
		expectedHeader      string
	}{
		{name: "Empty", xDNSPrefetchControl: "", expectedHeader: ""},
		{name: "On", xDNSPrefetchControl: XDNSPrefetchControlOn, expectedHeader: "on"},
		{name: "Off", xDNSPrefetchControl: XDNSPrefetchControlOff, expectedHeader: "off"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			header := tc.xDNSPrefetchControl.String()
			if header != tc.expectedHeader {
				t.Errorf("Expected: %s\tActual: %s\n", tc.expectedHeader, header)
			}
		})
	}
}

func TestXDNSPrefetchControl_Exists(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                string
		xDNSPrefetchControl XDNSPrefetchControl
		expectedExists      bool
	}{
		{name: "Empty", xDNSPrefetchControl: "", expectedExists: false},
		{name: "On", xDNSPrefetchControl: XDNSPrefetchControlOn, expectedExists: true},
		{name: "Off", xDNSPrefetchControl: XDNSPrefetchControlOff, expectedExists: true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			exists := tc.xDNSPrefetchControl.Exists()
			if exists != tc.expectedExists {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedExists, exists)
			}
		})
	}
}
