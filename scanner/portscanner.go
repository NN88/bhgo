package main

/********************************************************
# simple port scanner

# help
./port_scanner -h

# Potential improvements
1) split the results channel to only receive open ports
2) take in specific port numbers
3) UDP flag (UDP is slow by default because of no response)

* You could do banner grabbing and all of that, but honestly
this script should just be very simple and fast.
********************************************************/

import (
	"flag"
	"fmt"
	"net"
	"sort"
	"strconv"
	"time"
)

// set vars
var ports = make(chan int, 100)
var results = make(chan int)
var openports []int

// functions
func print_consts(ip string, how_many_ports int) {
	fmt.Printf("[+] IP: %s \n", ip)
	fmt.Printf("[+] Consequetive ports: %d \n", how_many_ports)
	fmt.Printf("\n")
}

func worker(set_ip string, ports, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("%s:%d", set_ip, p)
		conn, err := net.Dial("tcp", address)

		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func main() {
	defer elapsed("script")()

	var set_ip string
	flag.StringVar(&set_ip, "ip", "0.0.0.0", "an IP4 address or hostname")
	var set_ports string
	flag.StringVar(&set_ports, "how_many_ports", "1000", "a max port number to scan up to consequetively")
	flag.Parse()

	// keeping this a string and converting because you can pass by value
	// when you use flag.Int or flag.String it passed by reference
	how_many_ports, _ := strconv.Atoi(set_ports)
	print_consts(set_ip, how_many_ports)

	fmt.Println("[+] Setting workers")
	for i := 0; i < cap(ports); i++ {
		go worker(set_ip, ports, results)
	}

	fmt.Println("[+] Setting ports")
	go func() {
		for i := 1; i <= how_many_ports; i++ {
			ports <- i
		}
	}()

	fmt.Println("[+] Getting open ports")
	for i := 0; i < how_many_ports; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	fmt.Println("[+] Closing channels")
	close(ports)
	close(results)

	fmt.Println("[+] Sorting")
	sort.Ints(openports)
	fmt.Printf("\n")
	for _, port := range openports {
		fmt.Printf("%d open \n", port)
	}
	fmt.Printf("\n")
}
