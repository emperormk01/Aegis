# LoFTe.live - Comprehensive Vulnerability Assessment

## Executive Summary
Deep analysis reveals multiple information disclosure vulnerabilities and security misconfigurations.

---

## Critical Findings

### 1. Deployment ID Exposure (CRITICAL)
- **Finding**: Full Vercel deployment ID exposed in multiple locations
- **Value**: `ndXFKSdta9eFY3JDvQKNY`
- **Impact**: Enables targeted attacks against Vercel infrastructure
- **Location**: HTML comments, API responses

### 2. Internal Component Disclosure (HIGH)
All React component names exposed via Next.js SSR:
- ThemeProvider (state management)
- Toaster (notifications)
- GoogleAnalytics (tracking - ID EMPTY)
- GoogleTagManager (tracking - ID EMPTY)  
- OutletBoundary (routing)
- ViewportBoundary (rendering)
- MetadataBoundary (SEO)
- MaxWidthWrapper (layout)
- IconMark (icons)

**Impact**: Reveals application architecture to attackers

### 3. Full Source Map Paths (HIGH)
13+ internal file paths exposed:
```
/_next/static/chunks/7d921709e2221006.css
/_next/static/chunks/6c08ddb64da1711c.css
/_next/static/chunks/9a6796bb1e351dff.js
/_next/static/chunks/eabd545ef1f861ad.js
/_next/static/chunks/38ab34174ddef122.js
/_next/static/chunks/48ec67728a41b3e5.js
/_next/static/chunks/turbopack-8dd4983c9af74e21.js
/_next/static/chunks/f5a3e0d2bee8cf00.js
/_next/static/chunks/aee6b08dc66dee8b.js
/_next/static/chunks/2d636f1de35f1b37.js
/_next/static/chunks/5817e784bcaf875c.js
/_next/static/chunks/6bcd65f79dd9e216.js
/_next/static/chunks/a6dad97d9634a72d.js
```

### 4. Technology Stack Disclosure (MEDIUM)
- **Framework**: Next.js 14+ with Turbopack
- **Runtime**: Vercel Edge
- **Styling**: Tailwind CSS
- **Fonts**: Inter, Playfair Display (full Google Fonts URLs exposed)

### 5. Empty Analytics Configuration (MEDIUM)
- Google Analytics ID: EMPTY (no tracking)
- GTM ID: EMPTY (no container)
- **Impact**: Privacy compliance issue + tracking not working

### 6. API Endpoint Information (MEDIUM)
- `api.lofte.live` returns deployment errors
- Vercel region: iad1 (US East)
- Error codes expose internal routing

### 7. Security Headers Missing (MEDIUM)
Missing headers detected:
- X-Content-Type-Options
- X-Frame-Options
- X-XSS-Protection
- Content-Security-Policy (partial)

### 8. Third-Party Domains (INFO)
- fonts.googleapis.com
- fonts.gstatic.com  
- www.googletagmanager.com (empty ID)
- favicon.ico with hash

---

## Attack Surface Analysis

### Potential Attack Vectors

1. **Vercel Deployment Attack**
   - Using deployment ID to enumerate other deployments
   - Target specific Vercel infrastructure

2. **Component-Based Exploitation**
   - Known vulnerabilities in exposed components
   - ThemeProvider could have XSS via theme injection

3. **Reconnaissance**
   - Full tech stack enables targeted exploit search
   - Turbopack specific vulnerabilities

---

## Recommendations

1. Remove deployment ID from all responses
2. Disable component name exposure in production
3. Configure proper analytics IDs or remove tracking code
4. Implement security headers
5. Use Next.js built-in component hiding
6. Remove internal paths from error messages

---

## References
- Tested: 2026-02-28
- Tools: curl, Scrapling, nuclei, manual analysis
