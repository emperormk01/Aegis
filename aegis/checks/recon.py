import socket
import httpx
import re

COMMON_SUBS = ["www", "api", "admin", "app", "dev", "staging", "test", "m", "beta", "mail"]
COMMON_PORTS = [21, 22, 25, 80, 110, 143, 443, 3306, 5432, 6379, 8080, 8443, 3000, 587]


def resolve(host: str) -> bool:
    try:
        socket.gethostbyname(host)
        return True
    except Exception:
        return False


def enumerate_subdomains(domain: str, wordlist: list[str] | None = None) -> list[dict]:
    words = wordlist or COMMON_SUBS
    out = []
    for w in words:
        host = f"{w}.{domain}"
        out.append({"host": host, "resolved": resolve(host)})
    return out


def scan_ports(host: str, ports: list[int] | None = None, timeout: float = 1.0) -> list[dict]:
    ports = ports or COMMON_PORTS
    out = []
    for p in ports:
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.settimeout(timeout)
        try:
            s.connect((host, p))
            s.close()
            out.append({"port": p, "open": True})
        except Exception:
            out.append({"port": p, "open": False})
        finally:
            try:
                s.close()
            except Exception:
                pass
    return out


def detect_stack(url: str, timeout: float = 8.0) -> dict:
    try:
        r = httpx.get(url, timeout=timeout, follow_redirects=True)
    except Exception as e:
        return {"url": url, "error": str(e)}
    headers = {k.lower(): v for k, v in r.headers.items()}
    body = r.text[:8000].lower()
    tech = []
    if "next" in body or headers.get("x-powered-by", "").lower() == "next.js":
        tech.append("Next.js")
    if "openresty" in headers.get("server", "").lower() or "openresty" in body:
        tech.append("OpenResty/Nginx")
    if "cdn.tailwindcss.com" in body or "tailwind" in body:
        tech.append("Tailwind CSS")
    return {"url": url, "server": headers.get("server", ""), "tech": tech, "status": r.status_code}
