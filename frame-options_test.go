package helmet

import "testing"

func TestFrameOptions_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		frameOptions   FrameOptions
		expectedHeader string
	}{
		{name: "Empty", frameOptions: "", expectedHeader: ""},
		{name: "None", frameOptions: FrameOptionsDeny, expectedHeader: "DENY"},
		{name: "Master Only", frameOptions: FrameOptionsSameOrigin, expectedHeader: "SAMEORIGIN"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			header := tc.frameOptions.String()
			if header != tc.expectedHeader {
				t.Errorf("Expected: %s\tActual: %s\n", tc.expectedHeader, header)
			}
		})
	}
}

func TestFrameOptions_Exists(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		frameOptions   FrameOptions
		expectedExists bool
	}{
		{name: "Empty", frameOptions: "", expectedExists: false},
		{name: "Deny", frameOptions: FrameOptionsDeny, expectedExists: true},
		{name: "Same Origin", frameOptions: FrameOptionsSameOrigin, expectedExists: true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			exists := tc.frameOptions.Exists()
			if exists != tc.expectedExists {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedExists, exists)
			}
		})
	}
}
