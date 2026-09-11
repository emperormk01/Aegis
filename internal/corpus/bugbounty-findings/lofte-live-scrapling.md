# lofte.live - Scrapling Deep Scan Results

**Date:** 2026-02-28  
**Tool:** Scrapling (StealthyFetcher)  
**Target:** https://www.lofte.live/

## Results

### Pages Discovered
- `/` - Main page (LoFT3 | Web3 Events)
- `/previous-events` - Previous events page
- Google Forms: `forms.gle/gwhB683FptSMNsE39` (registration)

### External Links
- Twitter: `x.com/hidreams__`
- Telegram: `T.me/lofte_live`

### JavaScript Files (16 found)
- Next.js chunks: `/_next/static/chunks/*.js`
- Potential attack surface for JS analysis

### Technology Stack
- Next.js (React framework)
- Vercel hosting
- No obvious API endpoints exposed
- No forms found
- No email addresses in source
- No hardcoded secrets found

### Security Notes
- Site uses Cloudflare protection
- Scrapling successfully bypassed with `StealthyFetcher` headless mode
- Next.js app - check for IDOR in event URLs
- Registration via Google Forms - check for Google OAuth misconfigurations

## How Scrapling Bypassed Cloudflare
```python
from scrapling.fetchers import StealthyFetcher
p = StealthyFetcher.fetch('https://www.lofte.live/', headless=True, network_idle=True)
```

Uses:
1. Browser fingerprint impersonation
2. TLS/SSL handshake mimicry  
3. Realistic headers (User-Agent, Accept-Language)
4. JavaScript execution via Playwright
5. Network idle detection for dynamic content
