import httpx

REQUIRED = {
    "strict-transport-security": "HSTS",
    "content-security-policy": "CSP",
    "x-frame-options": "X-Frame-Options",
    "x-content-type-options": "X-Content-Type-Options",
    "referrer-policy": "Referrer-Policy",
    "permissions-policy": "Permissions-Policy",
}

SENSITIVE_PATHS = [
    "/.env",
    "/.git/config",
    "/.git/HEAD",
    "/config.php",
    "/wp-config.php",
    "/server-status",
    "/phpmyadmin",
    "/admin",
    "/.aws/credentials",
    "/backup.zip",
    "/.DS_Store",
]


def check_headers(url: str, timeout: float = 10.0) -> dict:
    try:
        r = httpx.get(url, timeout=timeout, follow_redirects=True)
    except Exception as e:
        return {"url": url, "error": str(e)}
    headers = {k.lower(): v for k, v in r.headers.items()}
    missing = []
    present = {}
    for h, label in REQUIRED.items():
        if h in headers:
            present[label] = headers[h]
        else:
            missing.append(label)
    # cookie flags
    set_cookie = r.headers.get("set-cookie", "")
    cookie_issues = []
    if set_cookie:
        low = set_cookie.lower()
        if "httponly" not in low:
            cookie_issues.append("Set-Cookie missing HttpOnly")
        if "secure" not in low and url.startswith("https"):
            cookie_issues.append("Set-Cookie missing Secure")
        if "samesite" not in low:
            cookie_issues.append("Set-Cookie missing SameSite")
    return {
        "url": url,
        "status": r.status_code,
        "server": headers.get("server", ""),
        "missing": missing,
        "present": present,
        "cookie_issues": cookie_issues,
    }


def check_sensitive(base_url: str, timeout: float = 8.0) -> list[dict]:
    base = base_url.rstrip("/")
    out = []
    for p in SENSITIVE_PATHS:
        url = f"{base}{p}"
        try:
            r = httpx.get(url, timeout=timeout, follow_redirects=False)
            out.append({"path": p, "url": url, "status": r.status_code, "exposed": r.status_code in (200, 301, 302)})
        except Exception as e:
            out.append({"path": p, "url": url, "status": None, "error": str(e), "exposed": False})
    return out
