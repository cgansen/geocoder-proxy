package main

import (
	"fmt"
	"net/http"
)

type GeocoderProxy struct{}

func (gp GeocoderProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        fmt.Println(r.URL)
}
func main() {
        var gp GeocoderProxy
        http.ListenAndServe(":4444", gp)
}
