package main

import (
    "fmt"
    "net"
)

func main() {
    protocol := "icmp"
    netaddr, _ := net.ResolveIPAddr("ip4", "127.0.0.1")
    for{
        conn, _ := net.ListenIP("ip4:"+protocol, netaddr)
        
        buf := make([]byte, 1024)
        numRead, _, _ := conn.ReadFrom(buf)
        s := string(buf[:numRead])
        fmt.Println(s)
    }
}

