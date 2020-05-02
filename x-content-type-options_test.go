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

func TestXContentTypeOptions_Empty(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                string
		xContentTypeOptions XContentTypeOptions
		expectedEmpty       bool
	}{
		{name: "Empty", xContentTypeOptions: "", expectedEmpty: true},
		{name: "No Open", xContentTypeOptions: XContentTypeOptionsNoSniff, expectedEmpty: false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			exists := tc.xContentTypeOptions.Empty()
			if exists != tc.expectedEmpty {
				t.Errorf("Expected: %t\tActual: %t\n", tc.expectedEmpty, exists)
			}
		})
	}
}
