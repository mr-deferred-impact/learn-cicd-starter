package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey_EmptyHeader(t *testing.T) {
	header := http.Header{}
	_, err := GetAPIKey(header)

	if err != nil {
		t.Fatalf("expected error, got nil")
	} else {
		fmt.Println("Passed check with empty header! The function panics!")
	}
}
