package checks

import (
	"net"
	"net/http"
	"strings"
	"time"
)

var commonSubs = []string{"www", "api", "admin", "app", "dev", "staging", "test", "m", "beta", "mail"}
var commonPorts = []int{21, 22, 25, 80, 110, 143, 443, 3306, 5432, 6379, 8080, 8443, 3000, 587}

type SubResult struct {
	Host     string `json:"host"`
	Resolved bool   `json:"resolved"`
}

type PortResult struct {
	Port int  `json:"port"`
	Open bool `json:"open"`
}

type StackResult struct {
	URL    string   `json:"url"`
	Server string   `json:"server"`
	Tech   []string `json:"tech"`
	Status int      `json:"status"`
	Error  string   `json:"error,omitempty"`
}

func EnumerateSubdomains(domain string, wordlist []string) []SubResult {
	if wordlist == nil {
		wordlist = commonSubs
	}
	var out []SubResult
	for _, w := range wordlist {
		host := w + "." + domain
		_, err := net.LookupHost(host)
		out = append(out, SubResult{Host: host, Resolved: err == nil})
	}
	return out
}

func ScanPorts(host string, ports []int, timeout time.Duration) []PortResult {
	if ports == nil {
		ports = commonPorts
	}
	var out []PortResult
	for _, p := range ports {
		addr := net.JoinHostPort(host, itoa(p))
		conn, err := net.DialTimeout("tcp", addr, timeout)
		if err == nil {
			conn.Close()
			out = append(out, PortResult{Port: p, Open: true})
		} else {
			out = append(out, PortResult{Port: p, Open: false})
		}
	}
	return out
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	b := make([]byte, 0, 6)
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}

func DetectStack(url string, timeout time.Duration) StackResult {
	client := &http.Client{Timeout: timeout}
	resp, err := client.Get(url)
	if err != nil {
		return StackResult{URL: url, Error: err.Error()}
	}
	defer resp.Body.Close()
	buf := make([]byte, 8000)
	n, _ := resp.Body.Read(buf)
	body := strings.ToLower(string(buf[:n]))
	server := resp.Header.Get("Server")
	var tech []string
	if strings.Contains(body, "next") || strings.EqualFold(resp.Header.Get("X-Powered-By"), "Next.js") {
		tech = append(tech, "Next.js")
	}
	if strings.Contains(strings.ToLower(server), "openresty") || strings.Contains(body, "openresty") {
		tech = append(tech, "OpenResty/Nginx")
	}
	if strings.Contains(body, "cdn.tailwindcss.com") || strings.Contains(body, "tailwind") {
		tech = append(tech, "Tailwind CSS")
	}
	return StackResult{URL: url, Server: server, Tech: tech, Status: resp.StatusCode}
}
