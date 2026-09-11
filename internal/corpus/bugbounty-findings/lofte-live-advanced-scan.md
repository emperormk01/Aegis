# lofte.live - Advanced Scan Results

## Previous Findings
- Basic scan: Clean, Next.js + Vercel, no critical vulnerabilities
- Scrapling: Bypassed Cloudflare, found JS files

## Deep Analysis Results

### 1. JS Files Analysis
- No API endpoints or secrets found in JavaScript files
- Standard Next.js chunk files only
- No hardcoded credentials

### 2. Google Forms OAuth Test
**Form URL:** https://forms.gle/gwhB683FptSMNsE39

**Form Details:**
- Title: "LOFTE-3 ELITE DINNER NIGHT waitlist"
- Form ID: 1FAIpQLSc6qWAeRAnGn-iSkCFTgUQXQt9Uk7dWneFPAFg0sbQQI_gg

**Form Fields Collected:**
1. X username (required)
2. State of residence (for accommodation)
3. Joined TG for updates? (Yes/No)
4. Followed founder? (Yes/No)

**Potential Issues:**
- Form collects PII (names, locations, social handles)
- Google Forms tied to Google account (potential for Google account enumeration)
- No obvious IDOR - form submissions are private by default

**Security Notes:**
- Google Forms is generally secure by default
- Responses can only be viewed by form owner
- No OAuth tokens exposed in form source

### 3. Summary
The site is well-secured with:
- Next.js framework (latest)
- Vercel hosting with Cloudflare protection
- No obvious vulnerabilities found
- Professional implementation

**Recommendation:** This appears to be a legitimate, well-secured web3 event platform. No critical security issues identified for bug bounty purposes.

---
Scanned: lofte.live
Date: 2026-02-28
Tools: subfinder, httpx, nuclei, ffuf, Scrapling, curl
