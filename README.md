# Aegis

Security audit CLI in Go. Combines three repos into one tool:

- **bugbounty-findings** - 13 reports, S3 enumeration, race conditions, subdomain takeover, deployment ID and source map disclosure, empty analytics
- **kambegoye-scan** - Subdomain enumeration, port scanning, sensitive file checks, stack detection, nuclei-style findings
- **security-research** - OWASP Top 10, frontend and backend vuln research, CVE tracking 2024 to 2026

Single static binary, full corpus embedded.

## Install

```bash
go install github.com/emperormk01/Aegis@latest
# or build locally
go build -o aegis .
```

## Usage

```bash
# Full audit (now includes disclosure, nuclei, fuzz, takeover)
aegis scan https://example.com
aegis scan example.com --json > report.json

# Focused checks
aegis s3 mycompany
aegis headers https://example.com
aegis race https://example.com/api/checkout --method POST --count 50 --data '{"coupon":"TEST"}'
aegis recon example.com

# Corpus - search the ingested reports
aegis explain takeover
aegis explain "race condition"
aegis explain s3
aegis corpus
```

## What it checks

**Widened (from bugbounty reports):**
- Subdomain takeover (Vercel DEPLOYMENT_NOT_FOUND, NoSuchBucket, etc.)
- Next.js disclosure (__NEXT_DATA__, component names, deployment IDs)
- Source map exposure (/_next/static/chunks/pages/_app.js.map etc.)
- Empty analytics and GTM IDs
- Server version leak

**Nuclei templates:**
- openresty-detect, missing-sri, exposed-server-status, graphql-introspection, cors-misconfig

**Deeper scan (ffuf-style):**
- Fuzzes 26 common paths (admin, login, backup, .env.bak, swagger, actuator, etc.)

**Corpus:**
- 18 files embedded via go:embed, searchable with `aegis explain <keyword>`

## Project structure

```
cmd/              # Cobra commands: scan, s3, headers, race, recon, explain, corpus
internal/
  checks/         # s3, headers, race, recon, takeover, disclosure, nuclei, fuzz
  corpus/         # Embedded bugbounty + kambegoye + security-research reports
```

## License

MIT
