package checks

// Fu-JS technique from bug-bounty skill: JS files contain hidden
// endpoints, API keys, hardcoded secrets. Fetch same-origin scripts
// and regex for endpoints + secret patterns.

import (
	"io"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

type JSSecretFinding struct {
	File     string   `json:"file"`
	Type     string   `json:"type"`
	Matches  []string `json:"matches"`
	Severity string   `json:"severity"`
}

var secretPatterns = []struct {
	name     string
	re       *regexp.Regexp
	severity string
}{
	{"aws-key", regexp.MustCompile(`AKIA[0-9A-Z]{16}`), "high"},
	{"google-api-key", regexp.MustCompile(`AIza[0-9A-Za-z\-_]{35}`), "high"},
	{"firebase", regexp.MustCompile(`[a-z0-9-]+\.firebaseio\.com`), "medium"},
	{"jwt", regexp.MustCompile(`eyJ[A-Za-z0-9\-_]+\.eyJ[A-Za-z0-9\-_]+`), "medium"},
	{"private-key", regexp.MustCompile(`-----BEGIN (RSA |EC |DSA )?PRIVATE KEY-----`), "critical"},
	{"todo-secret", regexp.MustCompile(`(?i)(TODO|FIXME|XXX|HACK).{0,40}(password|secret|key|token)`), "low"},
	{"hardcoded-secret", regexp.MustCompile(`(?i)(api[_-]?key|secret|bearer)\s*[:=]\s*['"][A-Za-z0-9\-_]{8,}['"]`), "medium"},
}

var endpointPattern = regexp.MustCompile(`(?m)['"](/api/[a-zA-Z0-9/_\-{}]+|/v[0-9]+/[a-zA-Z0-9/_\-]+)['"]`)
var scriptSrcPattern = regexp.MustCompile(`(?i)<script[^>]+src=["']([^"']+)["']`)

func fetchText(client *http.Client, u string, maxBytes int64) (string, bool) {
	resp, err := client.Get(u)
	if err != nil {
		return "", false
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", false
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, maxBytes))
	if err != nil {
		return "", false
	}
	return string(body), true
}

func CheckJSSecrets(pageURL string, timeout time.Duration) []JSSecretFinding {
	client := &http.Client{Timeout: timeout}
	base, err := url.Parse(pageURL)
	if err != nil {
		return nil
	}
	html, ok := fetchText(client, pageURL, 2<<20)
	if !ok {
		return nil
	}
	// collect same-origin script URLs
	seen := map[string]bool{}
	var scripts []string
	for _, m := range scriptSrcPattern.FindAllStringSubmatch(html, 12) {
		raw := m[1]
		u, err := url.Parse(raw)
		if err != nil {
			continue
		}
		abs := base.ResolveReference(u)
		if abs.Host != "" && abs.Host != base.Host {
			continue // same-origin only
		}
		if seen[abs.String()] {
			continue
		}
		seen[abs.String()] = true
		scripts = append(scripts, abs.String())
	}
	var out []JSSecretFinding
	for _, jsURL := range scripts {
		js, ok := fetchText(client, jsURL, 1<<20)
		if !ok {
			continue
		}
		for _, p := range secretPatterns {
			found := map[string]bool{}
			var matches []string
			for _, m := range p.re.FindAllString(js, 5) {
				if !found[m] {
					found[m] = true
					matches = append(matches, m)
				}
			}
			if len(matches) > 0 {
				out = append(out, JSSecretFinding{File: jsURL, Type: p.name, Matches: matches, Severity: p.severity})
			}
		}
		// endpoint disclosure (informational)
		eps := map[string]bool{}
		var endpoints []string
		for _, m := range endpointPattern.FindAllStringSubmatch(js, 20) {
			if !eps[m[1]] {
				eps[m[1]] = true
				endpoints = append(endpoints, m[1])
			}
		}
		if len(endpoints) > 0 {
			out = append(out, JSSecretFinding{File: jsURL, Type: "endpoints", Matches: endpoints, Severity: "info"})
		}
	}
	return out
}
