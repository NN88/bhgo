package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/miekg/dns"
)

func print_documentation() {
	usage := fmt.Sprintf(`
		Description: 
		This displays any name servers associated with a domain
		
		Usage: 
		%s url
		- url : can be any url such as megacorpone.com
		
		TIP: 
		remember the domain of .com or any valid top level domain
	
		`, os.Args[0])
	fmt.Println(usage)
	os.Exit(1)
}

func main() {
	// if no arguments given print documentation
	args := os.Args[1:]
	if len(args) != 1 {
		print_documentation()
	}

	// Look up NS records in DNS
	domain := os.Args[1]
	fqdn := dns.Fqdn(domain)
	servers, err := net.LookupNS(fqdn)
	if err != nil {
		log.Fatalf("[!] Error is: %v", err)
	}

	// print out what is found
	if len(servers) < 1 {
		fmt.Println("No records")
		return
	}
	for _, answer := range servers {
		fmt.Println(answer.Host)
	}
}
