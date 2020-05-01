package helmet

import "testing"

func TestXContentTypeOptions_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                string
		xContentTypeOptions XContentTypeOptions
		expectedHeader      string
	}{
		{name: "Empty", xContentTypeOptions: "", expectedHeader: ""},
		{name: "No Open", xContentTypeOptions: XContentTypeOptionsNoSniff, expectedHeader: "nosniff"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			header := tc.xContentTypeOptions.String()
			if header != tc.expectedHeader {
				t.Errorf("Expected: %s\tActual: %s\n", tc.expectedHeader, header)
			}
		})
	}
}

func TestXContentTypeOptions_Exists(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                string
		xContentTypeOptions XContentTypeOptions
		expectedExists      bool
	}{
		{name: "Empty", xContentTypeOptions: "", expectedExists: false},
		{name: "No Open", xContentTypeOptions: XContentTypeOptionsNoSniff, expectedExists: true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			exists := tc.xContentTypeOptions.Exists()
			if exists != tc.expectedExists {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedExists, exists)
			}
		})
	}
}
