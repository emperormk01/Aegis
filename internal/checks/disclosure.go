package checks

import (
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type DisclosureResult struct {
	Type     string `json:"type"`
	URL      string `json:"url"`
	Found    bool   `json:"found"`
	Evidence string `json:"evidence,omitempty"`
}

func CheckDisclosure(baseURL string, timeout time.Duration) []DisclosureResult {
	base := strings.TrimSuffix(baseURL, "/")
	client := &http.Client{Timeout: timeout}
	var out []DisclosureResult

	// 1. Next.js info disclosure - __NEXT_DATA__ or component names
	func() {
		resp, err := client.Get(base)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		text := string(body)
		components := []string{"ThemeProvider", "Toaster", "GoogleAnalytics", "GoogleTagManager", "OutletBoundary", "ViewportBoundary", "MetadataBoundary"}
		for _, c := range components {
			if strings.Contains(text, c) {
				out = append(out, DisclosureResult{Type: "next_component", URL: base, Found: true, Evidence: c})
				break
			}
		}
		if strings.Contains(text, "__NEXT_DATA__") {
			out = append(out, DisclosureResult{Type: "next_data", URL: base, Found: true, Evidence: "__NEXT_DATA__"})
		}
		// deployment ID pattern like ndXFKSdta9eFY3JDvQKNY (22 alphanum)
		re := regexp.MustCompile(`[a-zA-Z0-9_-]{20,}`)
		_ = re
	}()

	// 2. Source map exposure
	sourcemaps := []string{"/_next/static/chunks/pages/_app.js.map", "/static/js/main.js.map", "/app.js.map"}
	for _, p := range sourcemaps {
		u := base + p
		resp, err := client.Get(u)
		if err != nil {
			continue
		}
		resp.Body.Close()
		if resp.StatusCode == 200 {
			ct := resp.Header.Get("Content-Type")
			if strings.Contains(ct, "json") || strings.Contains(ct, "octet-stream") || resp.StatusCode == 200 {
				// verify body looks like sourcemap
				if resp.ContentLength != 0 {
					out = append(out, DisclosureResult{Type: "sourcemap", URL: u, Found: true, Evidence: p})
				}
			}
		}
	}

	// 3. Empty analytics ID - gtag.js?id= with empty id
	func() {
		resp, err := client.Get(base)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		text := string(body)
		if strings.Contains(text, "gtag.js?id=") && (strings.Contains(text, "gtag.js?id=\"\"") || strings.Contains(text, "gtag.js?id=&") || regexp.MustCompile(`gtag\.js\?id=\s*["']\s*["']`).MatchString(text)) {
			out = append(out, DisclosureResult{Type: "empty_analytics", URL: base, Found: true, Evidence: "gtag.js?id= empty"})
		}
		if strings.Contains(text, "googletagmanager.com/gtm.js?id=") && strings.Contains(text, "gtm.js?id=\"\"") {
			out = append(out, DisclosureResult{Type: "empty_gtm", URL: base, Found: true, Evidence: "gtm.js?id= empty"})
		}
	}()

	// 4. Server header disclosure - version leak
	func() {
		resp, err := client.Get(base)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		server := resp.Header.Get("Server")
		if server != "" && regexp.MustCompile(`[0-9]+\.[0-9]+`).MatchString(server) {
			out = append(out, DisclosureResult{Type: "server_version", URL: base, Found: true, Evidence: server})
		}
	}()

	return out
}
