package src

import (
	"flag"
	"net/http"
)

func ServerConfInit() ServerConf {
	var addr = flag.String("a","0.0.0.0:8080","Address")
	var cert = flag.String("crt","none","Certificate")
	var pem  = flag.String("key","none","TLS Key")
	flag.Usage = Help
	flag.Parse() 

	x := ServerConf {
		Addr: *addr,
		Cert: *cert,
		Pem: *pem,
	}

	if *cert != "none" && *pem != "none" {
		x.tls = true
	} else {
		x.tls = false
	}
	return x
}

func (c ServerConf) InitServer() {
	handler := &Proxy{}
	
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