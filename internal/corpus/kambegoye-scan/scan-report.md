# Bug Bounty Scan Report: kambegoye.com

**Scan Date:** 2026-03-01
**Target:** https://kambegoye.com

---

## Executive Summary

Kambegoye.com is a Niger-based service marketplace (similar to TaskRabbit) that helps users find workers/contractors in Niamey. The site is built with Node.js, Nginx/OpenResty, and Tailwind CSS. Several open ports and services were identified, but no critical vulnerabilities were found.

---

## Subdomain Enumeration

| Subdomain | Status |
|-----------|--------|
| kambegoye.com | Resolved |
| www.kambegoye.com | Resolved |

**Tools Used:** subfinder, assetfinder, amass
**Result:** Minimal attack surface - only 2 subdomains found.

---

## Technology Stack

| Technology | Version |
|------------|---------|
| Web Server | OpenResty / Nginx |
| Backend | Node.js 20.20.0 |
| Frontend | Tailwind CSS |
| CDN | cdn.tailwindcss.com |

---

## Open Ports & Services

| Port | Service | Info |
|------|---------|------|
| 21 | FTP | pure-ftpd detected |
| 25 | SMTP | Exim 4.99.1 |
| 110 | POP3 | - |
| 143 | IMAP | - |
| 3306 | MySQL | MariaDB 5.5.5-10.11.15 |
| 587 | SMTP (TLS) | Exim 4.99.1 |

**Note:** MySQL accepts native password authentication (potential weak config).

---

## Sensitive Files Check

| File | Status |
|------|--------|
| /.env | Not found |
| /.git/config | Not found |
| /config.php | 200 (Placeholder: "It works! NodeJS") |
| /wp-config.php | Not found |
| /admin | 200 (Placeholder) |
| /login | 200 (Placeholder) |
| /phpmyadmin | 200 (Placeholder) |
| /server-status | 200 (Placeholder) |
| /xmlrpc.php | 403 (Forbidden) |

**Findings:**
- Multiple placeholder pages exposed (config.php, admin, login, phpmyadmin, server-status)
- No sensitive configuration files leaked
- server-status could be disabled to prevent information disclosure

---

## Nuclei Scan Results

### Info Findings:
- openresty-detect (OpenResty server identified)
- ftp-detect (FTP service)
- exim-detect (SMTP server version: 4.99.1)
- imap-detect (IMAP service)
- mysql-detect (MySQL/MariaDB)
- pop3-detect (POP3 service)
- smtp-detect (SMTP on ports 25, 587)
- mysql-native-password (MySQL using native password auth)
- missing-sri (Subresource Integrity not configured for Tailwind CDN)

### No Critical/High Findings

---

## Directory Fuzzing (ffuf)

Basic directory enumeration did not reveal any sensitive paths beyond the placeholder pages.

---

## Security Recommendations

1. **Disable placeholder pages** - Remove or password-protect /admin, /login, /phpmyadmin, /server-status if not in use
2. **Disable server-status** - Can expose server information
3. **Configure Subresource Integrity** - Add SRI hashes for Tailwind CDN
4. **Restrict MySQL access** - Ensure port 3306 is not exposed externally; disable native password auth
5. **FTP/SMTP hardening** - Review exposed services if not needed publicly

---

## Conclusion

**Severity: LOW**

No critical or high-severity vulnerabilities were identified. The site appears to be a legitimate business with standard placeholder pages for future development. Recommended actions are primarily hardening and reducing information disclosure.
