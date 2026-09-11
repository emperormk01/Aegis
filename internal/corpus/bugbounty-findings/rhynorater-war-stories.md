# Bug Bounty War Stories - Key Learnings from Rhynorater (DEF CON 32)

## Overview
- Speaker: Justin Gardner (Rhynorater)
- 11 critical bugs documented over 2-3 years
- Total bounties: $225K-$400K
- Categories: 3 Easy, 2 Medium, 4 Hard, 2 Very Hard

## Key Vulnerabilities Found

### 1. EngineX 403 Bypass to PII Leak (Easy)
**Bug:** Double URL-encoded path traversal bypass
- Company moved internal software to public use (threat model changed)
- Found `/api/internal/getallusers` blocked at reverse proxy level
- Double URL-encoded `%s` passed through EngineX, decoded to `s` on backend
- Result: 4.5M users PII leaked
- Bounty: $15K-$20K

**Takeaways:**
- Identify odd 403 pages (reverse proxy vs application-level)
- Know bypass tricks: URL encoding, path normalization
- Change in threat model = new attack surface

### 2. Arbitrary Account Takeover via Docs (Easy)
**Bug:** Hardcoded credentials in API docs
- OTP login app but docs showed password-based auth
- Found credentials in JS files
- Could submit any password for any user
- Bounty: $40K-$60K

**Takeaways:**
- read_file the documentation
- Check JS files for credentials
- Think about how auth actually works vs docs

### 3. Numeric IDOR - Password Hash/Reset Token Leak (Easy)
**Bug:** IDOR in API leaked password reset tokens
- Spent 40 hours understanding the app first
- Numeric user ID allowed access to password hashes
- Bounty: $40K-$60K

**Takeaways:**
- Get deep into applications
- Invest time upfront to understand the attack surface

### 4. Blind XSS via SMS to Account Takeover (Medium)
**Bug:** XSS in CRM dealer panel
- Car dealer CRM showed customer leads
- Found XSS in lead info
- Stored XSS executed in admin panel
- Bounty: $12K-$15K

**Takeaways:**
- Check all user input points
- Look for internal panels (dealer, admin, CRM)

### 5. SQL Injection via Protobuff (Hard)
**Bug:** SQLi in protobuf parameters
- Binary format: param name + null byte + value length + value
- Used strct.pack in Python to craft payloads
- ChatGPT helped write the exploit
- Bounty: $30K

**Takeaways:**
- read_file articles/documentation for binary formats
- Dont shy away from complex targets
- Use AI tools for help with unfamiliar code

### 6. IoT Router RCE via Malicious Settings File (Hard)
**Bug:** Path traversal in file write + auto-run exe
- Extracted firmware via hardware hacking (eMMC chip)
- Found file write via strct.pack
- Overwrote exe that ran every 2 seconds
- Bounty: $15K-$30K

**Takeaways:**
- read_file the docs
- Hardware hacking: hot air, flux, eMMC reader
- Use FCC website to identify chips

### 7. Mobile App VoIP Calling - Call Hijacking (Very Hard)
**Bug:** SIP header injection in VoIP app
- Custom TLS pinning (bypassed with Frida)
- SIP protocol injection
- Could inject arbitrary From headers
- Auto-answer calls on victim devices
- Bounty: $20K-$50K

**Takeaways:**
- Dont be afraid of new protocols
- Use transparent proxies (Polar Proxy)
- Binary patching with hex editor for strings

### 8. Third-Party Registration Hijacking (Very Hard)
**Bug:** Register token could register for other users
- Third-party registration feature
- Could map any users address of record to attacker
- Anyone calling victim would ring attackers device
- Bounty: $20K-$50K

## Tools & Techniques Used

### Recon
- subfinder (subdomains)
- amass
- httpx (probing)
- JS file analysis

### Scanning
- nuclei (vulnerabilities)
- nmap

### Fuzzing
- ffuf

### Hardware
- Hot air rework station
- Flux
- eMMC reader
- BGA chip identification

### Mobile
- Frida (root detection bypass)
- Objection (APK patching)
- TLS pinning bypass scripts
- Polar Proxy (transparent proxy)
- SIP debugging

### Binary Exploitation
- strct.pack (Python)
- Hex editing
- Ghidra for reverse engineering

## Key Mindset Shifts

1. **Threat Model Changes** - Internal apps made public = new attack surface
2. **Know Your Bypass Tricks** - URL encoding, path traversal, normalization
3. **Deep Application Understanding** - 40+ hours before finding bugs
4. **Dont Fear Complex Targets** - Use AI, docs, articles
5. **New Protocols** - SIP, gRPC, WebSocket - learn them
6. **Hardware is Valid** - IoT devices have big bounties

## Resources Mentioned
- Critical Thinking Bug Bounty Podcast
- Port Swagger talk on path normalization
- Kaido (advisory platform)
