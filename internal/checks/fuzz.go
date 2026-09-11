package checks

import (
	"net/http"
	"strings"
	"time"
)

var fuzzPaths = []string{
	"admin", "login", "api", "backup", "config", ".env.bak", "old", "test",
	"debug", "staging", "dev", "internal", "private", "secret", "v1", "v2",
	"graphql", "swagger", "openapi.json", "actuator", "health", "metrics",
}

type FuzzResult struct {
	Path   string `json:"path"`
	URL    string `json:"url"`
	Status int    `json:"status"`
	Found  bool   `json:"found"`
}

func Fuzz(baseURL string, timeout time.Duration, extra []string) []FuzzResult {
	base := strings.TrimSuffix(baseURL, "/")
	paths := fuzzPaths
	if len(extra) > 0 {
		paths = append(paths, extra...)
	}
	client := &http.Client{Timeout: timeout, CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }}
	var out []FuzzResult
	for _, p := range paths {
		u := base + "/" + strings.TrimPrefix(p, "/")
		resp, err := client.Get(u)
		if err != nil {
			continue
		}
		resp.Body.Close()
		found := resp.StatusCode == 200 || resp.StatusCode == 301 || resp.StatusCode == 302 || resp.StatusCode == 401 || resp.StatusCode == 403
		if found {
			out = append(out, FuzzResult{Path: p, URL: u, Status: resp.StatusCode, Found: true})
		}
	}
	return out
}
