package main

import (
	"fmt"
	"net/http"
        "net/url"
        "io/ioutil"
        "log"
)

const GMapsHost = "http://maps.googleapis.com/maps/api/geocode/"
const GMapsResponseFormat = "json"        

type GeocoderProxy struct{}

type GoogleMapsRequest struct {}

func (gp GeocoderProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        var gmr GoogleMapsRequest
        fmt.Println(r.URL)

        gmaps_url := url.URL{Scheme: "https", Host: "maps.googleapis.com", Path: "/maps/api/geocode/json" }         
        q := gmaps_url.Query()
        q.Set("address", r.FormValue("address"))
        gmaps_url.RawQuery = q.Encode()
        
        gmr.Get(gmaps_url.String())
        // fmt.Println(gmaps_url.String())
}

func (gmr *GoogleMapsRequest) Get(url string) (response string) {
        // TODO: sign request
        fmt.Println(url)
        
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
