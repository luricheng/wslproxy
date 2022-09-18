package main

import (
	"flag"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	var formPort, toPort string

	flag.StringVar(&formPort, "f", "1234", "from port")
	flag.StringVar(&toPort, "t", "1234", "to port")
	flag.Parse()

	remote, err := url.Parse("http://127.0.0.1:" + toPort)
	if err != nil {
		panic(err)
	}

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			r.Host = remote.Host
			p.ServeHTTP(w, r)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(":"+formPort, nil)
	if err != nil {
		panic(err)
	}
}
