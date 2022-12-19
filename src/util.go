package src

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

// print error and exit
func PrintErr(err error) {
	log.Printf("[\u001B[1;31m!\u001B[0;0m]- %s",err.Error())
	os.Exit(1)
}

// Print string with green + signal meaning success
func PrintSucc(s string) {
	log.Printf("[\u001B[1;32m+\u001B[0;0m]- %s",s)
}

// Compile X-Forwarded-For header
func CompileXForwardHead(head http.Header, clientIP string) {
	if prior, ok := head["X-Forwarded-For"]; ok {
		clientIP = strings.Join(prior, ", ") + ", " + clientIP
	}
	head.Set("X-Forwarded-For", clientIP)
}

// Tranfer header from dst to src
func TranferHeader(dst, src http.Header) {
	for i, v := range src {
		for _, vlue := range v {
			dst.Add(i,vlue)
		}
	}
}

// Hop-by-hop headers. It is good to be removed when connecting to the backend
// http://www.w3.org/Protocols/rfc2616/rfc2616-sec13.html
var PHeaders = []string{
	"Connection",
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te",
	"Trailers",
	"Transfer-Encoding",
	"Upgrade",
	"Batata",
}

func DelProxyHead(header *http.Header) {
	for _, h := range PHeaders {
		header.Del(h)
	}
}

// Generate unique request ID. It is used mostly for
// debugging purpose.
func GenRequestID(leng int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, leng)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// Add request id in request for the back end
func AddRequestID(head *http.Header, id string) {
	head.Set("X-Ghoxy-ID",id)
}