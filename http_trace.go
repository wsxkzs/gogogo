package main 

import (
	"net/http/httptrace"
	"log"
	"net/http"
	"time"
	"fmt"
	"crypto/tls"
	"os"
)

func timeGet(url string ) {
	req , _ := http.NewRequest("GET", url , nil)

	var  start , connect , dns ,  tlsHandshake time.Time

	trace := &httptrace.ClientTrace{
		DNSStart: func(dsi httptrace.DNSStartInfo) {dns = time.Now()},
		DNSDone:  func(ddi httptrace.DNSDoneInfo){
			fmt.Printf("DNS Done: %v\n", time.Since(dns))
		} ,
		TLSHandshakeStart: func(){tlsHandshake = time.Now()},
		TLSHandshakeDone: func(cs tls.ConnectionState, err error) {
			fmt.Printf("TLS Handshake: %v\n", time.Since(tlsHandshake))
		} ,
		ConnectStart: func(network, addr string) { connect = time.Now() },
        ConnectDone: func(network, addr string, err error) {
            fmt.Printf("Connect time: %v\n", time.Since(connect))
        },

        GotFirstResponseByte: func() {
            fmt.Printf("Time from start to first byte: %v\n", time.Since(start))
        },
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(),trace))
	start = time.Now()
	if _, err := http.DefaultTransport.RoundTrip(req); err !=nil {
		log.Fatal(err)
	}
	fmt.Printf("Total time: %v\n", time.Since(start))
}


func main(){
	urls := os.Args[1:2]
	for _ , url  := range  urls {
	
         timeGet(url)
	}
	

}