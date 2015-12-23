package httpauth

import (
	"encoding/base64"
	"net/http"
	"testing"
)

func TestBasicAuthAuthenticate(t *testing.T) {
	// Provide a minimal test implementation.
	authOpts := AuthOptions{
		Realm:       "Restricted",
		AuthFunc: func(u,p string) bool { return true },
	}

	b := &basicAuth{
		opts: authOpts,
	}

	r := &http.Request{}
	r.Method = "GET"

	// Provide auth data, but no Authorization header
	if b.authenticate(r) != false {
		t.Fatal("No Authorization header supplied.")
	}

	// Initialise the map for HTTP headers
	r.Header = http.Header(make(map[string][]string))

	// Set a malformed/bad header
	r.Header.Set("Authorization", "    Basic")
	if b.authenticate(r) != false {
		t.Fatal("Malformed Authorization header supplied.")
	}

	// Test wrong formated credentials
	auth := base64.StdEncoding.EncodeToString([]byte("santaisnotreal"))
	r.Header.Set("Authorization", "Basic "+auth)
	if b.authenticate(r) != false {
		t.Fatal("Failed on wrong credentials")
	}

	// Test correct credentials
	auth = base64.StdEncoding.EncodeToString([]byte("for:bar"))
	r.Header.Set("Authorization", "Basic "+auth)
	if b.authenticate(r) != true {
		t.Fatal("Failed on correct credentials")
	}
}

func TestBasicAuthAutenticateWithouUserAndPass(t *testing.T) {
	b := basicAuth{opts: AuthOptions{}}

	if b.authenticate(nil) {
		t.Fatal("Should not authenticate if user or pass are not set on opts")
	}
}
