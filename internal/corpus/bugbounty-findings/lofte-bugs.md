# LOFT3
[truncated]
y.

[truncated]
# VULNERABILITY #1: Subdomain Takeover (CRITICAL)

**Endpoint:** api.lofte.live

**Response:**
```
The deployment could not be found on Vercel.

DEPLOYMENT_NOT_FOUND

iad1::vblqg-1772257370292-21a850008701
```

**Issue:** api.lofte.live points to a deleted/missing Vercel deployment.

**Impact:** Potential subdomain takeover if attacker can claim the Vercel project.

**Fix:** Point api.lofte.live to a valid deployment or remove the DNS record.

---

## VULNERABILITY #2: Empty Analytics IDs (LOW)

**Issue:** Google Analytics and Google Tag Manager are loaded but have empty IDs:
- gtag.js?id= (empty)
- gtm.js?id= (empty)

**Impact:** Analytics not working, tracking broken.

---

## VULNERABILITY #3: Information Disclosure via Error Messages

**Issue:** Detailed Next.js error pages expose:
- Framework version
- Deployment IDs
- Internal paths

---

## TECH STACK

- Next.js 14+ (Turbopack)
- Vercel hosting
- Empty GA4/GTM configuration
- Waitlist: Google Forms (3rd party)