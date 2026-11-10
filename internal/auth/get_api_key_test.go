package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := map[string]struct {
		headers     http.Header
		expectedKey string
		expectedErr error
	}{
		"happyPath": {http.Header{"Authorization": []string{"ApiKey my-api-key"}}, "my-api-key", nil},
        "noAuthHeader": {http.Header{}, "", ErrNoAuthHeaderIncluded},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			key, err := GetAPIKey(tc.headers)

			if tc.expectedErr != nil && err != tc.expectedErr {
				t.Fatalf("expected error: %v, got: %v", tc.expectedErr, err)
			}

			if tc.expectedErr == nil && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if key != tc.expectedKey {
				t.Fatalf("expected key: %s, got: %s", tc.expectedKey, key)
			}
		})
	}
}
