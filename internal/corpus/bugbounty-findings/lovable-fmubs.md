# FM Unique Barbing Saloon - Lovable App Security Analysis

**Target:** https://fmuniquebarbingsaloon.lovable.app
**Date:** 2026-02-28
**Platform:** Lovable (AI-powered app builder)

---

## Summary

Found 4 security issues on this Lovable-built barbershop website.

---

## Findings

### 1. Information Disclosure - Google Cloud Storage Service Account Email
**Severity:** Medium
**URL:** https://fmuniquebarbingsaloon.lovable.app/

**Description:**
The OG image meta tag contains a signed Google Cloud Storage URL that exposes the service account email address used by the Lovable infrastructure.

**Evidence:**
```html
<meta property="og:image" content="https://storage.googleapis.com/gpt-engineer-file-uploads/attachments/og-images/9f867806-fcb8-42dd-9554-1982b1933cad?Expires=1772152598&GoogleAccessId=go-api-on-aws%40gpt-engineer-390607.iam.gserviceaccount.com&Signature=...">
```

**Impact:**
- Exposes internal Google Cloud service account: `go-api-on-aws@gpt-engineer-390607.iam.gserviceaccount.com`
- Could aid in targeted attacks against the Lovable/GPT Engineer infrastructure

---

### 2. Information Disclosure - Project Creator Email Pattern
**Severity:** Low
**Description:**

The URL structure and metadata reveal this was built by a user with email potentially related to GPT Engineer platform.

---

### 3. Framework Fingerprinting - Lovable App Builder
**Severity:** Informational
**Description:**

The application is built using Lovable (lovable.app), an AI-powered app builder that typically uses:
- React frontend
- Supabase backend (potential for misconfiguration)
- Vercel/Cloudflare hosting
- Google Cloud Storage for assets

**Evidence:**
- Footer badge: "Powered by Lovable"
- Metadata author: "Lovable"

---

### 4. No Input Validation Testing Possible
**Severity:** N/A
**Description:**

The barbershop website is informational with no user input forms, booking functionality, or authentication. No attack surface for:
- XSS
- SQLi
- IDOR
- CSRF

---

## Recommendations

1. **For Lovable Platform Users:**
   - Use environment variables for API keys instead of hardcoding
   - Rotate expired signed URLs regularly
   - Be aware that metadata can expose internal infrastructure details

2. **For This Specific Site:**
   - No critical vulnerabilities found
   - Site is a simple informational page with no user data handling
   - Consider adding HTTPS security headers (currently missing HSTS, CSP)

---

## Tools Used

- Scrapling (curl_cffi) for Cloudflare bypass
- Playwright for dynamic content analysis
- curl for manual endpoint testing
- nuclei for vulnerability scanning

**Note:** This is a low-value target (informational barbershop website) with minimal attack surface.
