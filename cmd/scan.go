package cmd

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/emperormk01/Aegis/internal/checks"
	"github.com/spf13/cobra"
)

var scanJSON bool

var scanCmd = &cobra.Command{
	Use:   "scan [target]",
	Short: "Full audit: headers, sensitive files, stack, subdomains, ports",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := args[0]
		url := target
		if !strings.HasPrefix(target, "http") {
			url = "https://" + target
		}
		domain := target
		domain = strings.TrimPrefix(domain, "https://")
		domain = strings.TrimPrefix(domain, "http://")
		domain = strings.Split(domain, "/")[0]

		hdr := checks.CheckHeaders(url, 10*time.Second)
		sens := checks.CheckSensitive(url, 8*time.Second)
		stack := checks.DetectStack(url, 8*time.Second)
		subs := checks.EnumerateSubdomains(domain, nil)
		ports := checks.ScanPorts(domain, nil, 1*time.Second)
		takeover := checks.CheckTakeover(subs, 8*time.Second)
		disclosure := checks.CheckDisclosure(url, 8*time.Second)
		nuclei := checks.RunNuclei(url, 8*time.Second)
		fuzz := checks.Fuzz(url, 8*time.Second, nil)
		jssecrets := checks.CheckJSSecrets(url, 10*time.Second)
		jwt := checks.CheckJWT(url, 8*time.Second)

		if scanJSON {
			out := map[string]interface{}{
				"target": target, "headers": hdr, "sensitive": sens,
				"stack": stack, "subdomains": subs, "ports": ports,
				"takeover": takeover, "disclosure": disclosure, "nuclei": nuclei, "fuzz": fuzz,
				"jssecrets": jssecrets, "jwt": jwt,
			}
			b, _ := json.MarshalIndent(out, "", "  ")
			fmt.Println(string(b))
			return
		}

		fmt.Printf("\n=== Aegis scan: %s ===\n\n", target)
		fmt.Println("--- Security headers ---")
		for _, label := range []string{"HSTS", "CSP", "X-Frame-Options", "X-Content-Type-Options", "Referrer-Policy", "Permissions-Policy"} {
			present := true
			for _, m := range hdr.Missing {
				if m == label {
					present = false
				}
			}
			if present {
				fmt.Printf("  [OK] %s\n", label)
			} else {
				fmt.Printf("  [MISS] %s\n", label)
			}
		}
		for _, iss := range hdr.CookieIssues {
			fmt.Printf("  [WARN] %s\n", iss)
		}
		fmt.Println("\n--- Sensitive files ---")
		for _, r := range sens {
			tag := "hidden"
			if r.Exposed {
				tag = "EXPOSED"
			}
			fmt.Printf("  %s -> %d %s\n", r.Path, r.Status, tag)
		}
		fmt.Println("\n--- Disclosure ---")
		if len(disclosure) == 0 {
			fmt.Println("  [OK] no disclosure found")
		}
		for _, d := range disclosure {
			fmt.Printf("  [%s] %s -> %s\n", d.Type, d.URL, d.Evidence)
		}
		fmt.Println("\n--- Nuclei ---")
		if len(nuclei) == 0 {
			fmt.Println("  [OK] no template matched")
		}
		for _, n := range nuclei {
			fmt.Printf("  [%s] %s: %s\n", n.Severity, n.Template, n.Evidence)
		}
		fmt.Println("\n--- Fuzz (ffuf-style) ---")
		if len(fuzz) == 0 {
			fmt.Println("  [OK] no interesting paths")
		}
		for _, f := range fuzz {
			fmt.Printf("  %s -> %d\n", f.Path, f.Status)
		}
		fmt.Println("\n--- JS secrets (Fu-JS) ---")
		if len(jssecrets) == 0 {
			fmt.Println("  [OK] no secrets or endpoints in JS")
		}
		for _, j := range jssecrets {
			fmt.Printf("  [%s] %s in %s\n", j.Severity, j.Type, j.File)
			for _, m := range j.Matches {
				if len(m) > 80 {
					m = m[:80] + "..."
				}
				fmt.Printf("    - %s\n", m)
			}
		}
		fmt.Println("\n--- JWT ---")
		if len(jwt.Tokens) == 0 {
			fmt.Println("  [OK] no JWTs discovered")
		}
		for _, t := range jwt.Tokens {
			fmt.Printf("  [%s] alg=%s via %s: %s\n", t.Severity, t.Alg, t.Source, t.Note)
		}
		fmt.Println("\n--- Takeover ---")
		if len(takeover) == 0 {
			fmt.Println("  [OK] no takeover signals")
		}
		for _, t := range takeover {
			fmt.Printf("  [CRITICAL] %s -> %s (%s)\n", t.Host, t.Signal, t.URL)
		}
		fmt.Println("\n--- Stack ---")
		fmt.Printf("  Server: %s\n  Tech: %s\n", stack.Server, strings.Join(stack.Tech, ", "))
		fmt.Println("\n--- Subdomains ---")
		for _, s := range subs {
			status := "no"
			if s.Resolved {
				status = "yes"
			}
			fmt.Printf("  %s -> %s\n", s.Host, status)
		}
		fmt.Println("\n--- Ports ---")
		for _, p := range ports {
			status := "closed"
			if p.Open {
				status = "OPEN"
			}
			fmt.Printf("  %d -> %s\n", p.Port, status)
		}
	},
}

func init() {
	scanCmd.Flags().BoolVar(&scanJSON, "json", false, "Output JSON")
	rootCmd.AddCommand(scanCmd)
}
