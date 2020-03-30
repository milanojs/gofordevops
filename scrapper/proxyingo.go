package main

import (
	"math/rand"
	"net/http"
	"net/url"
)

// Public proxies from https://hidemyna.me
// These proxies are subject to frequent change.
// Please replace them if necessary.
var proxies []string = []string{
	"http://207.154.231.208:8080",
	"http://138.68.230.88:8080",
	"http://162.243.107.45:8080",
	"http://14.207.170.194:8213",
	"http://41.164.247.186:53281",
	"http://36.89.106.247:43184",
	"http://186.215.97.253:80",
}

func GetProxy(_ *http.Request) (*url.URL, error) {
	randomIndex := rand.Int31n(int32(len(proxies)) - int32(1))
	randomProxy := proxies[randomIndex]
	return url.Parse(randomProxy)
}

func main() {
	http.DefaultTransport.(*http.Transport).Proxy = GetProxy
	// Continue with your HTTP requests ...
}
