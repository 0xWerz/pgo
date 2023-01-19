package port

import (
    "fmt"
    "net"
    "time"
)

// ScanTCP scans for open TCP ports on the given IP address and port range
func ScanTCP(ip string, startPort int, endPort int) []int {
    openPorts := []int{}

    for port := startPort; port <= endPort; port++ {
        address := fmt.Sprintf("%s:%d", ip, port)

        conn, err := net.DialTimeout("tcp", address, time.Second*10)
        if err == nil {
            openPorts = append(openPorts, port)
            conn.Close()
        }
    }

    return openPorts
}

// ScanUDP scans for open UDP ports on the given IP address and port range
func ScanUDP(ip string, startPort int, endPort int) []int {
    openPorts := []int{}

    for port := startPort; port <= endPort; port++ {
        address := fmt.Sprintf("%s:%d", ip, port)

        conn, err := net.Dial("udp", address)
        if err == nil {
            openPorts = append(openPorts, port)
            conn.Close()
			}
    return openPorts
}