package main

import (
    "flag"
    "fmt"
    "os"
    "port"
    "strconv"
    "strings"
)

func main() {
    // Define command line flags for IP address, port range, and scan type
    ip := flag.String("ip", "127.0.0.1", "IP address to scan")
    ports := flag.String("p", "1-65535", "Range of ports to scan (e.g. 22-80)")
    scanType := flag.String("t", "tcp", "Type of scan (tcp or udp)")
    help := flag.Bool("h", false, "Show help menu")
    flag.Parse()

    // Show help menu and exit if -h flag is passed
    if *help {
        flag.PrintDefaults()
        os.Exit(0)
    }
    // Split the ports range on "-" and convert to integers
    portRange := strings.Split(*ports, "-")
    startPort, _ := strconv.Atoi(portRange[0])
    endPort, _ := strconv.Atoi(portRange[1])

    // Validate input and exit if invalid
    if startPort > endPort {
        fmt.Println("Error: Starting port must be less than ending port.")
        os.Exit(1)
    }
    if *scanType != "tcp" && *scanType != "udp" {
        fmt.Println("Error: Invalid scan type. Must be 'tcp' or 'udp'.")
        os.Exit(1)
    }

    // Run the appropriate scan based on the type flag
    openPorts := []int{}
    if *scanType == "tcp" {
        openPorts = port.ScanTCP(*ip, startPort, endPort)
    } else {
        openPorts = port.ScanUDP(*ip, startPort, endPort)
    }

    // Print the results
    fmt.Printf("Open %s ports on %s: ", *scanType, *ip)
    for _, port := range openPorts {
        fmt.Printf("%d ", port)
    }
    fmt.Println()
}
