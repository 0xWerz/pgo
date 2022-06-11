package main

import (
	"flag"
	"fmt"

	"github.com/0xWerz/port-scanner/port"
)

func main() {
	hostname := flag.String("h", "", "Host ip/hostname")
	start := flag.Int("r1", 22, "port range start")
	end := flag.Int("r2", 161, "port range end")
	flag.Parse()
	fmt.Println("starting", *hostname, "port scan from", *start, "to", *end)
	port.GetOpenPorts(*hostname, port.PortRange{Start: *start, End: *end})
}
