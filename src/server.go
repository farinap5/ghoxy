package src

import (
	"io"
	"log"
	"net"
	"net/http"
)



func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Request ID
	rid := GenRequestID(16)

	client := &http.Client{}
	r.RequestURI = ""

	// Validate authentication
	if p.basicAuth {
		if !p.ValidateBasicAuth(r.Header) {
			msg := "\"Proxy Authentication Required\""
			log.Printf("rhost=%s method=%s url=%s statuscode=\"%s\" status=%s msg=%s id=%s", 
			r.RemoteAddr, r.Method, r.URL,"NA","error",msg,rid)
			http.Error(w,msg,http.StatusProxyAuthRequired)
			return
		}
	}

	// Compile X-Forwarded-For
	clientIP, _ ,err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		msg := "\"Address error\""
		log.Printf("rhost=%s method=%s url=%s statuscode=\"%s\" status=%s msg=%s id=%s", 
		r.RemoteAddr, r.Method, r.URL,"NA","error",msg,rid)
		http.Error(w,msg,http.StatusBadRequest)
		return
	}

	CompileXForwardHead(r.Header,clientIP)
	AddRequestID(&r.Header,rid)
	DelProxyHead(&r.Header)

	resp, err := client.Do(r)
	//PrintErr(err)
	if err != nil {
		msg := "\"Server Error\""
		log.Printf("rhost=%s method=%s url=%s statuscode=\"%s\" status=%s msg=%s id=%s", 
		r.RemoteAddr, r.Method, r.URL,"NA","error",msg,rid)
		http.Error(w,msg,http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	log.Printf("rhost=%s method=%s url=%s statuscode=\"%s\" status=%s msg=%s id=%s", 
	r.RemoteAddr, r.Method, r.URL,resp.Status,"OK","NA",rid)

	DelProxyHead(&resp.Header)
	TranferHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w,resp.Body)
}