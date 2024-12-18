package main

import (
	_ "embed"
	"log"
	"strings"

	"github.com/miekg/dns"

	"github.com/sujay1844/swiftie-dns/swiftiedns"
)

//go:embed data/ts_lyrics.csv
var lyricsFile string

func main() {
	lyricsFileReader := strings.NewReader(lyricsFile)
	songs, err := swiftiedns.InitDB(lyricsFileReader)
	if err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	port := ":8053"

	dns.HandleFunc(".", swiftiedns.HandleDNSRequest(songs))

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
