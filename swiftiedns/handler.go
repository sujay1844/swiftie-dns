package swiftiedns

import (
	"fmt"
	"log"
	"strings"

	"github.com/miekg/dns"
)

func HandleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
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
			rr, err := dns.NewRR(fmt.Sprintf("%s 3600 IN TXT \"%s\"", question.Name, responseText))
			if err != nil {
				log.Printf("Failed to create TXT record: %v", err)
				continue
			}
			msg.Answer = append(msg.Answer, rr)
		}
	}

	err := w.WriteMsg(&msg)
	if err != nil {
		log.Printf("Failed to send response: %v", err)
	}
}
