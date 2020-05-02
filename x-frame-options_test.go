package helmet

import "testing"

func TestXFrameOptions_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		xFrameOptions  XFrameOptions
		expectedHeader string
	}{
		{name: "Empty", xFrameOptions: "", expectedHeader: ""},
		{name: "None", xFrameOptions: XFrameOptionsDeny, expectedHeader: "DENY"},
		{name: "Master Only", xFrameOptions: XFrameOptionsSameOrigin, expectedHeader: "SAMEORIGIN"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			header := tc.xFrameOptions.String()
			if header != tc.expectedHeader {
				t.Errorf("Expected: %s\tActual: %s\n", tc.expectedHeader, header)
			}
		})
	}
}

func TestFrameOptions_Empty(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		xFrameOptions XFrameOptions
		expectedEmpty bool
	}{
		{name: "Empty", xFrameOptions: "", expectedEmpty: true},
		{name: "Deny", xFrameOptions: XFrameOptionsDeny, expectedEmpty: false},
		{name: "Same Origin", xFrameOptions: XFrameOptionsSameOrigin, expectedEmpty: false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			exists := tc.xFrameOptions.Empty()
			if exists != tc.expectedEmpty {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedEmpty, exists)
			}
		})
	}
}
