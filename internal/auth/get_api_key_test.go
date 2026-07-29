package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// This is our "table" of test cases
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError bool
	}{
		{
			name: "Valid API Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey 123456789"},
			},
			expectedKey:   "123456789",
			expectedError: false,
		},
		{
			name: "Missing Header entirely",
			headers: http.Header{},
			expectedKey:   "",
			expectedError: true,
		},
		{
			name: "Malformed Header (Wrong prefix)",
			headers: http.Header{
				"Authorization": []string{"Bearer 123456789"},
			},
			expectedKey:   "",
			expectedError: true,
		},
	}

	// This loop runs every test case in the table above
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// We call the actual function from your app's code
			key, err := GetAPIKey(tc.headers)

			// Check if the error status matches what we expected
			if (err != nil) != tc.expectedError {
				t.Fatalf("Test '%s' failed: expected error presence to be %v, got %v", tc.name, tc.expectedError, err)
			}

			// Check if the returned key matches what we expected
			if key != tc.expectedKey {
				t.Fatalf("Test '%s' failed: expected key %v, got %v", tc.name, tc.expectedKey, key)
			}
		})
	}
}