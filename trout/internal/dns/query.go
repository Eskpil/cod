package dns

import (
	"context"
	"fmt"
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
			req := new(dns.Msg)

			req.Id = dns.Id()
			req.RecursionDesired = true
			req.Question = make([]dns.Question, 1)
			req.Question[0] = dns.Question{dns.Fqdn(question.Name), question.Qtype, dns.ClassINET}

			c := new(dns.Client)
			res, _, err := c.Exchange(req, "1.1.1.1:53")

			fmt.Println(res.Answer)

			if err != nil {
				fmt.Println("error")
				// Our last resort is redirecting the request to google dns
				// directly through a NS record if asking google fails.
				msg.Answer = append(msg.Answer, &dns.NS{
					Hdr: dns.RR_Header{
						Name:   question.Name,
						Rrtype: dns.TypeNS,
						Class:  dns.ClassINET,
						Ttl:    1800,
					},
					Ns: "8.8.8.8",
				})

				continue
			}

			for _, answer := range res.Answer {
				switch answer.Header().Rrtype {
				case dns.TypeA:
					if a, ok := answer.(*dns.A); ok {
						msg.Answer = append(msg.Answer, &dns.A{
							Hdr: dns.RR_Header{Name: a.Header().Name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: a.Header().Ttl},
							A:   a.A,
						})
					}
				case dns.TypeTXT:
					log.Info(" > Doing TXT record")
					if txt, ok := answer.(*dns.TXT); ok {
						msg.Answer = append(msg.Answer, &dns.TXT{
							Hdr: dns.RR_Header{Name: txt.Header().Name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: txt.Header().Ttl},
							Txt: txt.Txt,
						})
					}
				case dns.TypeAAAA:
					if aaaa, ok := answer.(*dns.AAAA); ok {
						msg.Answer = append(msg.Answer, &dns.AAAA{
							Hdr:  dns.RR_Header{Name: aaaa.Header().Name, Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: aaaa.Header().Ttl},
							AAAA: aaaa.AAAA,
						})
					}
				}

			}
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
