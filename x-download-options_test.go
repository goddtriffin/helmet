package helmet

import "testing"

func TestXDownloadOptions_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name             string
		xDownloadOptions XDownloadOptions
		expectedHeader   string
	}{
		{name: "Empty", xDownloadOptions: "", expectedHeader: ""},
		{name: "No Open", xDownloadOptions: XDownloadOptionsNoOpen, expectedHeader: "noopen"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			header := tc.xDownloadOptions.String()
			if header != tc.expectedHeader {
				t.Errorf("Expected: %s\tActual: %s\n", tc.expectedHeader, header)
			}
		})
	}
}

func TestXDownloadOptions_Empty(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name             string
		xDownloadOptions XDownloadOptions
		expectedEmpty    bool
	}{
		{name: "Empty", xDownloadOptions: "", expectedEmpty: true},
		{name: "No Open", xDownloadOptions: XDownloadOptionsNoOpen, expectedEmpty: false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			exists := tc.xDownloadOptions.Empty()
			if exists != tc.expectedEmpty {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedEmpty, exists)
			}
		})
	}
}
