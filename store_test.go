package authkit

import (
	"net/http"
	//"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestNewCookie(t *testing.T) {
	config := &Config{
		Path:     "/",
		MaxAge:   1000,
		HttpOnly: true,
	}
	cookie := NewCookie("name", "value", config)
	// We copy Expires because time generated by seconds
	// and already increased
	date := cookie.Expires
	testCookie := &http.Cookie{
		Name:     "name",
		Value:    "value",
		Domain:   config.Domain,
		Path:     config.Path,
		MaxAge:   config.MaxAge,
		HttpOnly: config.HttpOnly,
		Secure:   config.Secure,
		Expires:  date,
	}
	if !reflect.DeepEqual(cookie, testCookie) {
		t.Fatal("failed create new cookie", cookie, testCookie)
	}

	config = &Config{
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
	}
	cookie = NewCookie("name", "value", config)
	date = cookie.Expires
	testCookie = &http.Cookie{
		Name:     "name",
		Value:    "value",
		Domain:   config.Domain,
		Path:     config.Path,
		MaxAge:   config.MaxAge,
		HttpOnly: config.HttpOnly,
		Secure:   config.Secure,
		Expires:  date,
	}
	if !reflect.DeepEqual(cookie, testCookie) {
		t.Fatal("failed create new cookie with MaxAge = 0")
	}

	config = &Config{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	cookie = NewCookie("name", "value", config)
	date = cookie.Expires
	testCookie = &http.Cookie{
		Name:     "name",
		Value:    "value",
		Domain:   config.Domain,
		Path:     config.Path,
		MaxAge:   config.MaxAge,
		HttpOnly: config.HttpOnly,
		Secure:   config.Secure,
		Expires:  date,
	}
	if !reflect.DeepEqual(cookie, testCookie) {
		t.Fatal("failed create new cookie with MaxAge = -1")
	}
	if date != time.Unix(1, 0) {
		t.Fatal("failed create new cookie - wrong expired date")
	}
}
