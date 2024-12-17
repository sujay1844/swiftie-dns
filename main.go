package main

import (
	"log"

	"github.com/miekg/dns"
)

func main() {
	port := ":8053"

	dns.HandleFunc(".", handleDNSRequest)

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
