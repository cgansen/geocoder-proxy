package main

import (
	"fmt"
	"net/http"
        "io/ioutil"
        "log"
)

type GeocoderProxy struct{}
type GoogleMapsRequest struct {}

func (gp GeocoderProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        var gmr GoogleMapsRequest
        fmt.Println(r.URL)
        resp := gmr.Get("http://example.com" + r.URL.String())
        
        fmt.Println(resp)
}

func (gmr *GoogleMapsRequest) Get(url string) (response string) {
        // TODO: sign request
        resp, err := http.Get(url)
        
        body, err := ioutil.ReadAll(resp.Body)
        resp.Body.Close()
        
        if err != nil {
             log.Fatal(err)
        }        
        
        return string(body[:])
} 

func main() {
        fmt.Println("Listening on port 4444")
        var gp GeocoderProxy
        http.ListenAndServe(":4444", gp)
}
