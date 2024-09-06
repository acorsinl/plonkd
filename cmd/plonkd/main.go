package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func getIP(r *http.Request) string {
	// Try to get the IP from the X-Forwarded-For header (common in reverse proxy setups)
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		ips := strings.Split(ip, ",")
		ip = strings.TrimSpace(ips[0])
		if net.ParseIP(ip) != nil {
			return ip
		}
	}

	// Try to get the IP from the X-Real-Ip header
	ip = r.Header.Get("X-Real-Ip")
	if ip != "" && net.ParseIP(ip) != nil {
		return ip
	}

	// Fallback: get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "Unable to determine IP"
	}

	return ip
}

func handler(w http.ResponseWriter, r *http.Request) {
	ip := getIP(r)
	fmt.Fprintf(w, "%s\n", ip)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
