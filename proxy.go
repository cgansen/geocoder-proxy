package main

import (
	"fmt"
	"net/http"
        "net/url"
        "io/ioutil"
        "log"
        "crypto/hmac"
        "crypto/sha1"
        "encoding/base64"        
)

type GeocoderProxy struct{}
type GoogleMapsRequest struct {}

func (gp GeocoderProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        var gmr GoogleMapsRequest                
        resp := gmr.Get(r)
        w.Write([]byte(resp))
}

func (gmr *GoogleMapsRequest) Get(orig_req *http.Request) (response string) {        
        gmaps_url := url.URL{Scheme: "https", Host: "maps.googleapis.com", Path: orig_req.URL.Path, RawQuery: orig_req.URL.RawQuery }
        
        signature := SignRequest(orig_req.RequestURI)
        fmt.Println("generated signature: " + signature)
        
        gmaps_url.RawQuery += ("&signature=" +  signature)
        
        fmt.Println("complete url: " + gmaps_url.String())
        
        to_fetch := gmaps_url.String()
        resp, err := http.Get(to_fetch)
        body, err := ioutil.ReadAll(resp.Body)
        resp.Body.Close()        
        if err != nil {
             log.Fatal(err)
        }                
        return string(body[:])
} 

func SignRequest(unsigned_url string) (signed_url string) {
        var output []byte
                
        // setup the signing hash
        key, _ := base64.URLEncoding.DecodeString("sAKq6oNk-th0b96RRWxOctAt9ic=")
        shash := hmac.New(sha1.New, key)

        fmt.Println("string_to_sign: " + unsigned_url)        

        shash.Write([]byte(unsigned_url))

        return base64.URLEncoding.EncodeToString(shash.Sum(output))
}

func main() {
        fmt.Println("Listening on port 4444")
        var gp GeocoderProxy
        http.ListenAndServe(":4444", gp)
}
