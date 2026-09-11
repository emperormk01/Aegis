package checks

import (
	"io"
	"net/http"
	"strings"
	"time"
)

var takeoverSignals = []string{
	"DEPLOYMENT_NOT_FOUND",
	"There is no deployment for this URL",
	"NoSuchBucket",
	"Unrecognized domain",
	"Repository not found",
	"404 - Not Found",
	"Heroku | No such app",
}

type TakeoverResult struct {
	Host      string `json:"host"`
	URL       string `json:"url"`
	Status    int    `json:"status"`
	Vulnerable bool  `json:"vulnerable"`
	Signal    string `json:"signal,omitempty"`
}

func CheckTakeover(subdomains []SubResult, timeout time.Duration) []TakeoverResult {
	var out []TakeoverResult
	client := &http.Client{Timeout: timeout}
	for _, s := range subdomains {
		// only check unresolved or 404-like? we check all resolved hosts that might be dangling
		url := "https://" + s.Host
		resp, err := client.Get(url)
		if err != nil {
			// DNS failure means not vulnerable via HTTP, but could be dangling
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		text := string(body)
		for _, sig := range takeoverSignals {
			if strings.Contains(text, sig) {
				out = append(out, TakeoverResult{Host: s.Host, URL: url, Status: resp.StatusCode, Vulnerable: true, Signal: sig})
				break
			}
		}
	}
	return out
}
