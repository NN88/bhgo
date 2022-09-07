package main

import (
	"fmt"
	"os"
)

func print_documentation() {
	usage := fmt.Sprintf(`
	Description: 
	This script takes in a list of name servers and attempts a zone transfer on each one.
	
	Usage: 
	%s list.txt
	- list.txt : can be any list that has dns name servers
	
	TIP:
	remember to review your list to make sure it is one name server per line

	`, os.Args[0])
	fmt.Println(usage)
	os.Exit(1)
}

// make a list of name servers
// host -t ns megacorpone > list.txt
// pass those as command line arguments
// perform a zone transfer for each name server
// Instructions for use

func main() {
	// if no arguments given print documentation
	args := os.Args[1:]
	if len(args) != 1 {
		print_documentation()
	}

}

// Attempt zone transfer

// var msg dns.Msg
// fqdn := dns.Fqdn("megacorpone.com")
// msg.SetQuestion(fqdn, dns.TypeA)
// in, err := dns.Exchange(&msg, "8.8.8.8:53")
// if err != nil {
// 	panic(err)
// }
// if len(in.Answer) < 1 {
// 	fmt.Println("No records")
// 	return
// }
// for _, answer := range in.Answer {
// 	if a, ok := answer.(*dns.A); ok {
// 		fmt.Println(a.A)
// 	}
// }

// func ZoneTransfer(domain string) Results {
// 	results := NewResultSet()
// 	fqdn := dns.Fqdn(domain)

// 	servers, err := net.LookupNS(domain)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, server := range servers {
// 		msg := new(dns.Msg)
// 		msg.SetAxfr(fqdn)

// 		transfer := new(dns.Transfer)
// 		answerChan, err := transfer.In(msg, net.JoinHostPort(server.Host, "53"))
// 		if err != nil {
// 			log.Println(err)
// 			continue
// 		}

// 		for envelope := range answerChan {
// 			if envelope.Error != nil {
// 				log.Println(envelope.Error)
// 				break
// 			}

// 			for _, rr := range envelope.RR {
// 				switch v := rr.(type) {
// 				case *dns.A:
// 					results.Add(strings.TrimRight(v.Header().Name, "."), v.A.String())
// 				case *dns.AAAA:
// 					results.Add(strings.TrimRight(v.Header().Name, "."), v.AAAA.String())
// 				default:
// 				}
// 			}
// 		}
// 	}

// 	return results.Results()
// }
