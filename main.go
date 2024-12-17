package main

import (
	"log"
	"strings"

	"github.com/miekg/dns"
)

func getResponses(name string) ([]string, error) {
	return []string{"Hello " + name, "Bye " + name}, nil
}

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)
	msg.Authoritative = true

	for _, question := range r.Question {
		log.Printf("Received query for: %s, Type: %d", question.Name, question.Qtype)

		name := strings.TrimSuffix(question.Name, ".")

		responses, err := getResponses(name)
		if err != nil {
			log.Printf("Failed to get responses: %v", err)
			continue
		}
		for _, responseText := range responses {
			rr, err := dns.NewRR(question.Name + " 3600 IN TXT \"" + responseText + "\"")
			if err != nil {
				log.Printf("Failed to create TXT record: %v", err)
				continue
			}
			// Add the TXT record to the response message
			msg.Answer = append(msg.Answer, rr)
		}
	}

	err := w.WriteMsg(&msg)
	if err != nil {
		log.Printf("Failed to send response: %v", err)
	}
}

func main() {
	port := ":4353"

	dns.HandleFunc(".", handleDNSRequest)

	server := &dns.Server{
		Addr: port,
		Net:  "udp",
	}

	log.Printf("Starting DNS server on %s", port)

	// Start the DNS server
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	defer server.Shutdown()
}
