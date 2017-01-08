package authkit

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"reflect"
)

// TestStore - struct for testing Store interface
type TestStore struct{}

// New - implement Store interface
func (ts TestStore) New(req *http.Request, name string) (*Session, error) {
	return &Session{
		name:name,
	}, nil
}

// Save - implement Store interface
func (ts TestStore) Save(w http.ResponseWriter, session *Session) error {
	return nil
}

// TestNewSession - test creating Session
func TestNewSession(t *testing.T) {
	var store = TestStore{}
	var sess *Session = NewSession(store, "test")
	var newSession *Session = &Session{
		store:  store,
		name:   "test",
		Values: make(map[string]interface{}),
	}
	if !reflect.DeepEqual(sess, newSession) {
		t.Fatal("Failed to init session", sess, newSession)
	}
}

// TestSessionMethods - test Sessions methodd
func TestSessionMethods(t *testing.T) {
	store := TestStore{}
	req, err := http.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		t.Fatal("failed to create request", err)
	}
	w := httptest.NewRecorder()

	session, err := store.New(req, "test")
	if err != nil {
		t.Fatal("failed to create session", err)
	}

	if session.Name() != "test" {
		t.Fatal("failed get session name", err)
	}

	var sess *Session = NewSession(store, "test")
	if !reflect.DeepEqual(sess.Store(), store) {
		t.Fatal("failed get session store", err)
	}

	err = sess.Save(w)
	if err != nil {
			t.Fatal("failed to save session", err)
	}
}
