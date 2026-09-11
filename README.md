# Aegis

Security audit CLI in Go. Combines three repos into one tool:

- **bugbounty-findings** - S3 enumeration, race condition patterns, 13 real scan reports
- **kambegoye-scan** - Subdomain enumeration, port scanning, sensitive file checks, stack detection
- **security-research** - OWASP Top 10, header checks, CVE context

Single static binary, no Python needed.

## Install

```bash
go install github.com/emperormk01/Aegis@latest
# or build locally
go build -o aegis .
```

## Usage

```bash
# Full audit
aegis scan https://example.com
aegis scan example.com --json > report.json

# S3 bucket permutation
aegis s3 mycompany

# Security headers
aegis headers https://example.com

# Race condition test
aegis race https://example.com/api/checkout --method POST --count 50 --data '{"coupon":"TEST"}'

# Recon only
aegis recon example.com
```

## What it checks

- S3 buckets across 9 permutations, detects public ACL and listable contents
- Security headers (HSTS, CSP, X-Frame-Options, etc.) and cookie flags (HttpOnly, Secure, SameSite)
- Sensitive file exposure (/.env, /.git/config, /server-status, etc.)
- Stack detection (Next.js, OpenResty/Nginx, Tailwind)
- Subdomain enumeration via DNS resolve and port scanning via TCP connect
- Race condition harness with parallel requests and divergent response detection

## Project structure

```
cmd/           # Cobra commands: scan, s3, headers, race, recon
internal/checks/
  s3.go
  headers.go
  race.go
  recon.go
```

## License

MIT
