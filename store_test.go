package authkit

import (
	"net/http"
	//"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewCookie(t *testing.T) {
	config := &Config{
		Path:     "/",
		MaxAge:   1000,
		HttpOnly: true,
	}
	cookie := NewCookie("name", "value", config)
	testCookie := &http.Cookie{
		Name:     "name",
		Value:    "value",
		Domain:   config.Domain,
		Path:     config.Path,
		MaxAge:   config.MaxAge,
		HttpOnly: config.HttpOnly,
		Secure:   config.Secure,
	}
	if !reflect.DeepEqual(cookie, testCookie) {
		t.Fatal("failed create new cookie")
	}
}
