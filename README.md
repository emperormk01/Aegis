# Aegis

Security audit kit in Go, not just a tool. Combines three repos into one kit:

- **bugbounty-findings** - 13 reports, S3 enumeration, race conditions, subdomain takeover, deployment ID and source map disclosure, empty analytics
- **kambegoye-scan** - Subdomain enumeration, port scanning, sensitive file checks, stack detection, nuclei-style findings
- **security-research** - OWASP Top 10, frontend and backend vuln research, CVE tracking 2024 to 2026

Single static binary, full corpus embedded. Aegis is a kit - S3, headers, race, JS secrets, JWT, takeover, nuclei, fuzz, plus searchable corpus - not a single-purpose tool.

## Install - multiple options

### Option 1: npm (with binary)

```bash
npm install -g aegis-kit
# or without install
npx aegis-kit --help
# or
bunx aegis-kit --help
```

### Option 2: GitHub Releases (single binary)

```bash
# One-liner (linux/mac)
curl -fsSL https://raw.githubusercontent.com/emperormk01/Aegis/main/install.sh 2>/dev/null | bash || go install github.com/emperormk01/Aegis@latest

# Direct download (pick platform)
curl -L https://github.com/emperormk01/Aegis/releases/latest/download/aegis-linux-x64 -o /usr/local/bin/aegis && chmod +x /usr/local/bin/aegis
# also: aegis-linux-arm64, aegis-darwin-x64, aegis-darwin-arm64
```

Releases build only on `v*` tags. Commits without tags do nothing.

### Option 3: From source (Go)

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
- sqli-error, xss-reflected, robots-txt, security-txt, graphql-field-suggestion (from bug-bounty skill)

**JS secrets (Fu-JS technique):**
- `aegis js <url>` scans same-origin scripts for AWS keys, Google keys, Firebase, JWTs, private keys, hardcoded secrets, plus API endpoint disclosure

**JWT (passive):**
- `aegis jwt <url>` discovers JWTs in cookies and page bodies, decodes header, flags alg:none (critical) and HS256 notes

**Deeper scan (ffuf-style):**
- Fuzzes 26 common paths (admin, login, backup, .env.bak, swagger, actuator, etc.)

**Corpus:**
- 21 files embedded via go:embed, searchable with `aegis explain <keyword>` (includes 485KB bug-bounty skill + report guide + recon script)

## Project structure

```
cmd/              # Cobra commands: scan, s3, headers, race, recon, explain, corpus, js, jwt
internal/
  checks/         # s3, headers, race, recon, takeover, disclosure, nuclei, fuzz, jssecrets, jwt
  corpus/         # Embedded bugbounty + kambegoye + security-research + bugbounty-skill reports
```

## License

MIT
