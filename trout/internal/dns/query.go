package dns

import (
	"context"
	"net"
	"time"

	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"

	recordService "github.com/eskpil/cod/trout/internal/records"
)

func (s *Service) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	msg := new(dns.Msg)
	msg.SetReply(r)
	msg.Authoritative = true

	for _, question := range r.Question {
		log.Infof("Trying to resolve \"%d\" record for \"%s\"", question.Qtype, question.Name)
		records, err := recordService.Search(ctx, question.Name)

		if err != nil || 0 >= len(records) {
			log.Warningf(" > Unknown record: \"%s\" tried to be resolved but failed redirecting request to \"8.8.8.8\"", question.Name)
			// We dont have the record, lets try sending the request to someone who has it (google).
			msg.Answer = append(msg.Answer, &dns.NS{
				Hdr: dns.RR_Header{
					Name:   question.Name,
					Rrtype: dns.TypeNS,
					Class:  dns.ClassINET,
					Ttl:    1800,
				},
				Ns: "8.8.8.8",
			})

			w.WriteMsg(msg)
			return
		}

		for _, record := range records {
			if record.Type != question.Qtype {
				continue
			}

			switch question.Qtype {
			case dns.TypeA:
				msg.Answer = append(msg.Answer, &dns.A{
					Hdr: dns.RR_Header{Name: question.Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: record.Ttl},
					A:   net.ParseIP(record.Value.(string)),
				})

			case dns.TypeAAAA:
				msg.Answer = append(msg.Answer, &dns.AAAA{
					Hdr:  dns.RR_Header{Name: question.Name, Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: record.Ttl},
					AAAA: net.ParseIP(record.Value.(string)),
				})
			case dns.TypeTXT:
				msg.Answer = append(msg.Answer, &dns.TXT{
					Hdr: dns.RR_Header{Name: question.Name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: record.Ttl},
					Txt: record.Value.([]string),
				})
			}
		}

	}

	w.WriteMsg(msg)
}
