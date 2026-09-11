package checks

import (
	"io"
	"net/http"
	"strings"
	"time"
)

type NucleiFinding struct {
	Template string `json:"template"`
	URL      string `json:"url"`
	Severity string `json:"severity"`
	Matched  bool   `json:"matched"`
	Evidence string `json:"evidence,omitempty"`
}

type template struct {
	name     string
	severity string
	check    func(baseURL string, client *http.Client) (bool, string)
}

var templates = []template{
	{
		name: "openresty-detect", severity: "info",
		check: func(baseURL string, client *http.Client) (bool, string) {
			resp, err := client.Get(baseURL)
			if err != nil {
				return false, ""
			}
			defer resp.Body.Close()
			srv := strings.ToLower(resp.Header.Get("Server"))
			if strings.Contains(srv, "openresty") {
				return true, resp.Header.Get("Server")
			}
			return false, ""
		},
	},
	{
		name: "missing-sri", severity: "low",
		check: func(baseURL string, client *http.Client) (bool, string) {
			resp, err := client.Get(baseURL)
			if err != nil {
				return false, ""
			}
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			text := strings.ToLower(string(body))
			if strings.Contains(text, "cdn.tailwindcss.com") && !strings.Contains(text, "integrity=") {
				return true, "Tailwind CDN without SRI"
			}
			return false, ""
		},
	},
	{
		name: "exposed-server-status", severity: "low",
		check: func(baseURL string, client *http.Client) (bool, string) {
			u := strings.TrimSuffix(baseURL, "/") + "/server-status"
			resp, err := client.Get(u)
			if err != nil {
				return false, ""
			}
			defer resp.Body.Close()
			if resp.StatusCode == 200 {
				return true, "/server-status exposed"
			}
			return false, ""
		},
	},
	{
		name: "graphql-introspection", severity: "medium",
		check: func(baseURL string, client *http.Client) (bool, string) {
			for _, p := range []string{"/graphql", "/api/graphql"} {
				u := strings.TrimSuffix(baseURL, "/") + p
				body := `{"query":"{ __schema { types { name } } }"}`
				resp, err := client.Post(u, "application/json", strings.NewReader(body))
				if err != nil {
					continue
				}
				resp.Body.Close()
				if resp.StatusCode == 200 {
					return true, p + " introspection enabled"
				}
			}
			return false, ""
		},
	},
	{
		name: "cors-misconfig", severity: "medium",
		check: func(baseURL string, client *http.Client) (bool, string) {
			req, _ := http.NewRequest("GET", baseURL, nil)
			req.Header.Set("Origin", "https://evil.com")
			resp, err := client.Do(req)
			if err != nil {
				return false, ""
			}
			defer resp.Body.Close()
			acao := resp.Header.Get("Access-Control-Allow-Origin")
			if acao == "https://evil.com" || acao == "*" {
				return true, "CORS reflects arbitrary Origin: " + acao
			}
			return false, ""
		},
	},
}

func RunNuclei(baseURL string, timeout time.Duration) []NucleiFinding {
	client := &http.Client{Timeout: timeout}
	var out []NucleiFinding
	for _, t := range templates {
		matched, ev := t.check(baseURL, client)
		if matched {
			out = append(out, NucleiFinding{Template: t.name, URL: baseURL, Severity: t.severity, Matched: true, Evidence: ev})
		}
	}
	return out
}
