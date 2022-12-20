package src

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func (p Proxy) ValidateBasicAuth(head http.Header) bool {
	creds := head.Get("Proxy-Authorization")
	credSplit := strings.Split(creds," ")
	dec,_ := base64.StdEncoding.DecodeString(credSplit[1])
	credSplit = strings.Split(string(dec),":")
	if credSplit[0] == p.User && credSplit[1] == p.Pass {
		return true
	} else {
		return false
	}
}