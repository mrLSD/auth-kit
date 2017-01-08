package authkit

import "net/http"

type Store interface {
	New(req *http.Request, name string) (*Session, error)
	Save(w http.ResponseWriter) error
}

func NewCookie(name, value string, config *Config) *http.Cookie {
	&http.Cookie{
		Name:     name,
		Value:    value,
		Domain:   config.Domain,
		Path:     config.Path,
		MaxAge:   config.MaxAge,
		HttpOnly: config.HTTPOnly,
		Secure:   config.Secure,
	}
}
