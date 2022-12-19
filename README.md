# Ghoxy

A very simple HTTP proxy written in Golang.

It can be used for pivoting during pentest activities.

It does not supports back end with HTTPS.

**Help Menu**

```
go run main.go -h

COMMAND  DESCRIPTION
-------  -----------
-a       Address: 0.0.0.0:8080
-crt     TLS certificate
-key     TLS key
```

Starting listener and testing:

```
go run main.go -a 0.0.0.0:8080
2022/12/18 23:49:10 [+]- Starting Listener HTTP
```

Connecting with `curl`

```
curl -x "http://0.0.0.0:8080" "http://0.0.0.0:8000/"
```

Log message

```
$ go run main.go -a 0.0.0.0:8080
2022/12/18 23:52:09 [+]- Starting Listener HTTP
2022/12/18 23:52:15 rhost=127.0.0.1:52006 method=GET url=http://0.0.0.0:8000/ statuscode="200 OK" status=OK msg=NA id=VbhV4vC6AWX40IVU
2022/12/18 23:55:22 rhost=127.0.0.1:52010 method=GET url=http://0.0.0.0:8000/xxx statuscode="404 File not found" status=OK msg=NA id=WSP3NcHciWvqZTa3
2022/12/18 23:55:27 rhost=127.0.0.1:52014 method=HEAD url=http://0.0.0.0:8000/xxx statuscode="404 File not found" status=OK msg=NA id=N06RxRTZHWUsaD7H
```

To execute enabling TLS we just need to pass the key and cert through command line:

```
go run main.go -crt testssl.crt -key testssl.key
```

Testing HTTPS mode with `curl`

```
curl --proxy-insecure -x "https://0.0.0.0:8080" "http://0.0.0.0:80/"
curl --proxy-cacert rootCA.crt -x "https://0.0.0.0:8080" "https://0.0.0.0:80/"
```