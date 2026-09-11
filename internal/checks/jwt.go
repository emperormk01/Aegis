package checks

// JWT checks from bug-bounty skill. Passive only: discover JWTs in
// cookies and page bodies, decode the header, flag weak configs
// (alg:none, HS256 without verification context). No active cracking.

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type JWTInfo struct {
	Source   string `json:"source"`
	Alg      string `json:"alg"`
	Typ      string `json:"typ"`
	Note     string `json:"note"`
	Severity string `json:"severity"`
}

type JWTResult struct {
	Tokens []JWTInfo `json:"tokens"`
}

var jwtPattern = regexp.MustCompile(`eyJ[A-Za-z0-9\-_]+\.[A-Za-z0-9\-_]+\.[A-Za-z0-9\-_]*`)

func decodeSegment(seg string) map[string]interface{} {
	if m := len(seg) % 4; m != 0 {
		seg += strings.Repeat("=", 4-m)
	}
	raw, err := base64.URLEncoding.DecodeString(seg)
	if err != nil {
		return nil
	}
	var out map[string]interface{}
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil
	}
	return out
}

func CheckJWT(pageURL string, timeout time.Duration) JWTResult {
	client := &http.Client{Timeout: timeout}
	var out JWTResult
	resp, err := client.Get(pageURL)
	if err != nil {
		return out
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 2<<20))
	text := string(body)

	// cookies
	for _, c := range resp.Cookies() {
		for _, m := range jwtPattern.FindAllString(c.Value, 2) {
			if info := analyzeJWT(m, "cookie:"+c.Name); info != nil {
				out.Tokens = append(out.Tokens, *info)
			}
		}
	}
	// page body
	for _, m := range jwtPattern.FindAllString(text, 5) {
		if info := analyzeJWT(m, "page-body"); info != nil {
			dup := false
			for _, t := range out.Tokens {
				if t.Alg == info.Alg && t.Source == info.Source {
					dup = true
					break
				}
			}
			if !dup {
				out.Tokens = append(out.Tokens, *info)
			}
		}
	}
	return out
}

func analyzeJWT(token, source string) *JWTInfo {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return nil
	}
	header := decodeSegment(parts[0])
	if header == nil {
		return nil
	}
	alg, _ := header["alg"].(string)
	typ, _ := header["typ"].(string)
	info := &JWTInfo{Source: source, Alg: alg, Typ: typ}
	switch alg {
	case "none":
		info.Note = "alg:none accepted in token - signature can be stripped"
		info.Severity = "critical"
	case "HS256", "HS384", "HS512":
		info.Note = "HMAC alg - check for weak secret, kid injection, alg confusion"
		info.Severity = "info"
	default:
		info.Note = "asymmetric alg - check jku/x5u SSRF and claims manipulation"
		info.Severity = "info"
	}
	return info
}
