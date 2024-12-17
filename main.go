package main

import (
	"log"

	"github.com/miekg/dns"

	"github.com/sujay1844/swiftie-dns/swiftiedns"
)

func main() {
	swiftiedns.InitDB("./data/ts_lyrics.csv")

	port := ":8053"

	dns.HandleFunc(".", swiftiedns.HandleDNSRequest)

	server := &dns.Server{
		Addr: port,
		Net:  "udp",
	}

	log.Printf("Starting DNS server on %s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	defer server.Shutdown()
}
