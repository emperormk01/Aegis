# lofte.live - Next.js Error Page Information Disclosure

## Deployment Details Exposed

### 1. Deployment ID
- **ID:** `ndXFKSdta9eFY3JDvQKNY`
- **Location:** Found in HTML comments and data routes
- **URL Pattern:** `/_next/data/ndXFKSdta9eFY3JDvQKNY/index.json`

### 2. Framework Version
- **Next.js:** Version 14+ (inferred from turbopack usage)
- **Build Tool:** Turbopack (development) or Vercel Build
- **Evidence:** `turbopack-8dd4983c9af74e21.js`

### 3. Internal Paths Exposed

#### Static Asset Paths:
```
/_next/static/chunks/7d921709e2221006.css
/_next/static/chunks/6c08ddb64da1711c.css
/_next/static/chunks/9a6796bb1e351dff.js
/_next/static/chunks/eabd545ef1f861ad.js
/_next/static/chunks/38ab34174ddef122.js
/_next/static/chunks/48ec67728a41b3e5.js
/_next/static/chunks/f5a3e0d2bee8cf00.js
/_next/static/chunks/aee6b08dc66dee8b.js
/_next/static/chunks/2d636f1de35f1b37.js
/_next/static/chunks/5817e784bcaf875c.js
/_next/static/chunks/6bcd65f79dd9e216.js
/_next/static/chunks/a6dad97d9634a72d.js
```

#### Internal Component Names (from __next_f data):
- `ThemeProvider` (64,482 bytes)
- `GoogleAnalytics` (19,398 bytes)
- `GoogleTagManager` (7,564 bytes)
- `Toaster` (35,820 bytes)
- `MaxWidthWrapper` (12,874 bytes)
- `OutletBoundary` (59,399 bytes)
- `ViewportBoundary` (59,399 bytes)
- `MetadataBoundary` (59,399 bytes)
- `IconMark` (27,923 bytes)

### 4. Tech Stack Revealed

#### Frontend:
- Next.js 14+
- Turbopack
- React (inferred from __next_f React hydration)
- Tailwind CSS
- Framer Motion (likely - animation classes)
- Inter & Playfair Display fonts

#### Analytics (Configured but Empty):
- Google Analytics 4: Configured but `gaId=""` (empty)
- Google Tag Manager: Configured but `gtmId=""` (empty)

#### Third-Party:
- Google Fonts (fonts.googleapis.com)
- Google Tag Manager (googletagmanager.com)

### 5. Favicon Hash
- `favicon.0b3bf435.ico`

## Security Impact

| Finding | Severity | Description |
|---------|----------|-------------|
| Deployment ID Exposed | Medium | Unique Vercel deployment ID visible in HTML |
| Internal Paths Leaked | Low | Source map-like paths reveal internal structure |
| Component Names Exposed | Low | Reveals tech stack and internal architecture |
| Empty Analytics IDs | Info | GA4/GTM configured but not functional |

## Remediation

1. **Disable Error Page Debugging in Production:**
   ```javascript
   // next.config.js
   module.exports = {
     productionBrowserSourceMaps: false,
     reactStrictMode: true,
   }
   ```

2. **Remove Deployment ID from URLs:**
   - Use Vercel Analytics for clean URLs
   - Or configure a custom domain

3. **Hide Stack Traces:**
   - Ensure `NODE_ENV=production`
   - Configure custom error pages without implementation details

4. **Complete Analytics Setup:**
   - Add valid GA4 measurement ID
   - Add valid GTM container ID

## References

- Original 404 page: https://lofte.live/404
- Data route: https://lofte.live/_next/data/ndXFKSdta9eFY3JDvQKNY/index.json
