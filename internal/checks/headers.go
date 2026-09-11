package checks

import (
	"net/http"
	"strings"
	"time"
)

var requiredHeaders = map[string]string{
	"strict-transport-security": "HSTS",
	"content-security-policy":   "CSP",
	"x-frame-options":           "X-Frame-Options",
	"x-content-type-options":    "X-Content-Type-Options",
	"referrer-policy":           "Referrer-Policy",
	"permissions-policy":        "Permissions-Policy",
}

var sensitivePaths = []string{
	"/.env",
	"/.git/config",
	"/.git/HEAD",
	"/config.php",
	"/wp-config.php",
	"/server-status",
	"/phpmyadmin",
	"/admin",
	"/.aws/credentials",
	"/backup.zip",
	"/.DS_Store",
}

type HeaderResult struct {
	URL          string            `json:"url"`
	Status       int               `json:"status"`
	Server       string            `json:"server"`
	Missing      []string          `json:"missing"`
	Present      map[string]string `json:"present"`
	CookieIssues []string          `json:"cookie_issues"`
	Error        string            `json:"error,omitempty"`
}

type SensitiveResult struct {
	Path    string `json:"path"`
	URL     string `json:"url"`
	Status  int    `json:"status"`
	Exposed bool   `json:"exposed"`
	Error   string `json:"error,omitempty"`
}

func CheckHeaders(url string, timeout time.Duration) HeaderResult {
	client := &http.Client{Timeout: timeout}
	resp, err := client.Get(url)
	if err != nil {
		return HeaderResult{URL: url, Error: err.Error()}
	}
	defer resp.Body.Close()
	lower := map[string]string{}
	for k, v := range resp.Header {
		lower[strings.ToLower(k)] = strings.Join(v, ", ")
	}
	present := map[string]string{}
	var missing []string
	for h, label := range requiredHeaders {
		if v, ok := lower[h]; ok {
			present[label] = v
		} else {
			missing = append(missing, label)
		}
	}
	var cookieIssues []string
	if sc := resp.Header.Get("Set-Cookie"); sc != "" {
		low := strings.ToLower(sc)
		if !strings.Contains(low, "httponly") {
			cookieIssues = append(cookieIssues, "Set-Cookie missing HttpOnly")
		}
		if !strings.Contains(low, "secure") && strings.HasPrefix(url, "https") {
			cookieIssues = append(cookieIssues, "Set-Cookie missing Secure")
		}
		if !strings.Contains(low, "samesite") {
			cookieIssues = append(cookieIssues, "Set-Cookie missing SameSite")
		}
	}
	return HeaderResult{
		URL: url, Status: resp.StatusCode, Server: lower["server"],
		Missing: missing, Present: present, CookieIssues: cookieIssues,
	}
}

func CheckSensitive(baseURL string, timeout time.Duration) []SensitiveResult {
	base := strings.TrimSuffix(baseURL, "/")
	var out []SensitiveResult
	client := &http.Client{Timeout: timeout, CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }}
	for _, p := range sensitivePaths {
		u := base + p
		resp, err := client.Get(u)
		if err != nil {
			out = append(out, SensitiveResult{Path: p, URL: u, Error: err.Error()})
			continue
		}
		resp.Body.Close()
		exposed := resp.StatusCode == 200 || resp.StatusCode == 301 || resp.StatusCode == 302
		out = append(out, SensitiveResult{Path: p, URL: u, Status: resp.StatusCode, Exposed: exposed})
	}
	return out
}
