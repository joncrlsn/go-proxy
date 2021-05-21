// helpful resource with additional content:
// http://stackoverflow.com/questions/21270945/how-to-read-the-response-from-a-newsinglehostreverseproxy
package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"fmt"
)

const (
	//BaseUrl = "https://jsonplaceholder.typicode.com/"
	BaseUrl = "https://httpbin.org"
ListeningPort = "8888"
)

func main() {
	fmt.Println("Listening on port", ListeningPort)

	// Pass along all requests on this port
	http.HandleFunc("/", ProxyFunc)
	http.ListenAndServe(":"+ListeningPort, nil)
}

func ProxyFunc(w http.ResponseWriter, req *http.Request) {
	// parse the target url
	u, err := url.Parse(BaseUrl)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	 
	// Create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(u)

	// Update the headers to allow for SSL redirection
	req.URL.Host = u.Host
	req.URL.Scheme = u.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = u.Host

	proxy.ServeHTTP(w, req)
}

