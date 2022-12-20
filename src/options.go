package src

import "github.com/cheynewallace/tabby"

func Help() {
	t := tabby.New()
	t.AddHeader("COMMAND", "DESCRIPTION")
	t.AddLine("-a", "Address: 0.0.0.0:8080")
	t.AddLine("-crt", "TLS certificate")
	t.AddLine("-key", "TLS key")
	t.AddLine("-u", "User for basic auth")
	t.AddLine("-p", "Password for basic auth")
	print("\n")
	t.Print()
	print("\n")
}