package authkit

import (
	"net/http"
	"time"
)

type Store interface {
	New(req *http.Request, name string) (*Session, error)
	Save(w http.ResponseWriter, session *Session) error
}

// NewCookie returns an http.Cookie with the config set. It also sets
// the Expires field calculated based on the MaxAge value, for Internet
// Explorer compatibility.
func NewCookie(name, value string, config *Config) *http.Cookie {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Domain:   config.Domain,
		Path:     config.Path,
		MaxAge:   config.MaxAge,
		HttpOnly: config.HttpOnly,
		Secure:   config.Secure,
	}
	if config.MaxAge > 0 {
		date := time.Duration(config.MaxAge) * time.Second
		cookie.Expires = time.Now().Add(date)
	} else if config.MaxAge < 0 {
		// Set it to the past to expire now.
		cookie.Expires = time.Unix(1, 0)
	}
	return cookie
}
