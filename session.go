package authkit

import "net/http"

// Config is the set of session cookie properties
type Config struct {
	// Cookie Path scope
	Path string
	// Cookie Domain scope
	Domain string
	// MaxAge = 0 - it's mean field not set and cookie delete
	// MaxAge > 0 - it's mean field present and given in seconds
	MaxAge uint32
	// cookie may only be transferred over HTTPS
	Secure bool
	// Safe cookie from JavaScript access etc.
	HttpOnly bool
}

// Session represents Values which  a named bundle of maintained web state
// and sote session state
type Session struct {
	name   string
	store  *Store
	Config *Config
	Values map[string]interface{}
}

func NewSession(store Store, name string) *Session {
	return &Session{
		store:  store,
		name:   name,
		Values: make(map[string]interface{}),
	}
}

// Name return the session name
func (sess *Session) Name() string {
	return sess.name
}

// Save session to store.
// Identical to calling store.Save(w, session).
func (sess *Session) Save(w http.ResponseWriter) error {
	return sess.store.Save(w, sess)
}

// Store return session store
func (sess *Session) Store() Store {
	return sess.store
}
