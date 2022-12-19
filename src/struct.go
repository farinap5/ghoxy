package src

type ServerConf struct {
	Addr string
	tls  bool
	Cert string
	Pem  string
}

type Proxy struct {}