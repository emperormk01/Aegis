# Aegis

Security audit CLI in Python. Combines three repos into one tool:

- **bugbounty-findings** - S3 enumeration, race condition patterns, methodology and 13 real scan reports (Lofte, texttospeechpro, etc.)
- **kambegoye-scan** - Subdomain enumeration, port scanning, sensitive file checks, stack detection, nuclei-style findings
- **security-research** - OWASP Top 10, frontend and backend vulnerability research, CVE tracking for 2024 to 2026

No Rust. Pure Python with Click, httpx, and Rich.

## Install

```bash
pip install -e .
# or
pip install aegis
```

## Quick start

```bash
# Full audit for a URL or domain
aegis scan https://example.com
aegis scan example.com --json > report.json

# S3 bucket permutation attack
aegis s3 mycompany

# Security headers and cookie flags
aegis headers https://example.com

# Race condition test (TOCTOU)
aegis race https://example.com/api/checkout --method POST --count 50 --data '{"coupon":"TEST"}'

# Recon only
aegis recon example.com
```

## What it checks

**From bugbounty-findings/methodology.md:**

- S3 bucket enumeration across 9 permutations (static, assets, uploads, logs, backup, prod, dev)
- Public ACL and ListBucket detection
- Race condition harness with asyncio and httpx (configurable concurrency, divergent response detection)

**From kambegoye-scan/scan-report.md:**

- Subdomain enumeration (common wordlist and DNS resolve)
- Open ports via TCP connect (21, 25, 80, 110, 143, 443, 3306, etc.)
- Sensitive file exposure (/.env, /.git/config, /config.php, /server-status, /phpmyadmin, etc.)
- Stack detection (Next.js, OpenResty/Nginx, Tailwind)
- Nuclei style info findings

**From security-research:**

- OWASP-aligned header checks (HSTS, CSP, X-Frame-Options, X-Content-Type-Options, Referrer-Policy, Permissions-Policy)
- Cookie flag checks (HttpOnly, Secure, SameSite)
- CVE context for 2024 to 2026 in reports

## Project structure

```
aegis/
  cli.py              # Click group with scan, s3, headers, race, recon
  checks/
    s3.py             # S3 enumeration
    headers.py        # Security headers and sensitive paths
    race.py           # Async race harness
    recon.py          # Subdomains, ports, stack
```

## Development

```bash
python -m venv .venv && source .venv/bin/activate
pip install -e .
aegis scan https://example.com
```

## License

MIT
