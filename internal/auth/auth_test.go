package auth_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {

	type test struct {
		input  http.Header
		apikey string
		err    error
	}

	tests := []test{
		{http.Header{}, "", auth.ErrNoAuthHeaderIncluded},
		{http.Header{"Authorization": []string{"Bearer ApiKey"}}, "", errors.New("malformed authorization header")},
		{http.Header{"Authorization": []string{"ApiKey"}}, "", errors.New("malformed authorization header")},
	}

	for _, tc := range tests {

		apikey, err := auth.GetAPIKey(tc.input)

		if apikey != tc.apikey {
			t.Fatalf("expected: %v, got: %v", tc.apikey, apikey)
		}

		if err != tc.err {
			t.Fatalf("expected: %v, got: %v", tc.err, err)
		}

	}
}
