package src

import (
	"flag"
	"net/http"
)

func ServerConfInit() Proxy {
	var addr = flag.String("a","0.0.0.0:8080","Address")
	var cert = flag.String("crt","none","Certificate")
	var pem  = flag.String("key","none","TLS Key")

	var user = flag.String("u","none","User")
	var pass = flag.String("p","none","Pass")
	flag.Usage = Help
	flag.Parse()

	bauth := false
	if *user != "none" && *pass != "none" {
		PrintSucc("Basic authentication for user "+*user)
		bauth = true
	}

	x := Proxy {
		Addr: *addr,
		Cert: *cert,
		Pem: *pem,
		User: *user,
		Pass: *pass,
		basicAuth: bauth,
	}

	if *cert != "none" && *pem != "none" {
		x.tls = true
	} else {
		x.tls = false
	}
	return x
}

func (c Proxy) InitServer() {
	handler := &c
	
	server := &http.Server{
		Addr: c.Addr,
		Handler: handler,
	}
	
	if c.tls {
		PrintSucc("Starting Listener HTTPS")
		err := server.ListenAndServeTLS(c.Cert,c.Pem)
		if err != nil {
			PrintErr(err)
		}
	} else {
		PrintSucc("Starting Listener HTTP")
		err := server.ListenAndServe()
		if err != nil {
			PrintErr(err)
		}
	}
}

func Start() {
	c := ServerConfInit()
	c.InitServer()
}