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

func TestXDNSPrefetchControl_Empty(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                string
		xDNSPrefetchControl XDNSPrefetchControl
		expectedEmpty       bool
	}{
		{name: "Empty", xDNSPrefetchControl: "", expectedEmpty: true},
		{name: "On", xDNSPrefetchControl: XDNSPrefetchControlOn, expectedEmpty: false},
		{name: "Off", xDNSPrefetchControl: XDNSPrefetchControlOff, expectedEmpty: false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			exists := tc.xDNSPrefetchControl.Empty()
			if exists != tc.expectedEmpty {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedEmpty, exists)
			}
		})
	}
}
