package authkit

type Config struct {
	Path     string
	Domain   string
	MaxAge   uint32
	Secure   bool
	HttpOnly bool
}

type Session struct {
	name   string
	store  *Store
	Config *Config
	Values map[string]interface{}
}
