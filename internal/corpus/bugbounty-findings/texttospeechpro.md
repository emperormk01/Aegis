# TextToSpeechPro Vulnerability Scan

**Target:** https://www.texttospeechpro.com/voices?page=2
**Date:** 2026-02-27
**Tool:** Nuclei v3.7.0

## Results

### Subdomains Found
- smtp.texttospeechpro.com
- smtp.texttospeechpro.net
- smtp.texttospeechpro.org
- texttospeechpro.com

### Endpoints Discovered
- /admin
- /api
- /debug
- /login
- /uploads
- /assets
- /js
- /static
- /images
- /backup
- /private
- /config

### Vulnerability Scan Results
- **Critical:** 0
- **High:** 0
- **Medium:** 0
- **Low:** 0
- **Total:** 0

### Technology Stack
- Host: Vercel
- Framework: Next.js (likely)
- Security: Good default protection

### Notes
The site appears to be well-secured with proper redirects and no exposed vulnerabilities. All discovered endpoints return 307 redirects, indicating proper access controls.