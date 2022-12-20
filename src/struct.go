package src

type Proxy struct {
	basicAuth bool
	User 	  string
	Pass 	  string

	Addr string
	tls  bool
	Cert string
	Pem  string
}