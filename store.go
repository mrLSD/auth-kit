package authkit

import "net/http"

type Store interface {
	New(req *http.Request, name string) (*Session, error)
}
