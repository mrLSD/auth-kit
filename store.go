package authkit

import "net/http"

type Store interface {
	New(req *http.Request, name string) (*Session, error)
	Save(w http.ResponseWriter, session *Session) error
}

func NewCookie(name, value string, config *Config) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Domain:   config.Domain,
		Path:     config.Path,
		MaxAge:   config.MaxAge,
		HttpOnly: config.HttpOnly,
		Secure:   config.Secure,
	}
}
