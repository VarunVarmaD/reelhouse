package redis

import (
	"strings"
	"testing"
)

// A malformed URL must never echo the password back in the error, because
// main logs this error and hosting platforms retain logs.
func TestNewClient_InvalidURLDoesNotLeakPassword(t *testing.T) {
	const password = "s3cr3t-do-not-log"
	cases := map[string]string{
		"surrounding quotes": `"rediss://default:` + password + `@example.upstash.io:6379"`,
		"wrong scheme":       "http://default:" + password + "@example.upstash.io:6379",
	}

	for name, url := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := NewClient(url)
			if err == nil {
				t.Fatal("expected an error for malformed URL")
			}
			if strings.Contains(err.Error(), password) {
				t.Fatalf("error leaks password: %q", err.Error())
			}
		})
	}
}
