package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// GETAPIKeyTest
func TestGETAPIKeyHasToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "ApiKey token")

	apiKey, err := GetAPIKey(req.Header)
	want := "token"

	if !reflect.DeepEqual(want, apiKey) {
		t.Fatalf("expected: %v, got: %v", want, apiKey)
	}
}

func TestGETAPIKeyDoesNotHaveToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Something token")

	_, headerError := GetAPIKey(req.Header)
	want := errors.New("malformed authorization header")

	if !assert.Equal(t, want, headerError) {
		t.Fatalf("expected: %v, got: %v", want, headerError)
	}
}
