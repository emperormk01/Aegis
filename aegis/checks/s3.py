import httpx

PERMUTATIONS = [
    "{base}",
    "{base}-static",
    "{base}-assets",
    "{base}-uploads",
    "{base}-logs",
    "{base}-backup-2026",
    "{base}-prod",
    "{base}-dev",
    "{base}-staging",
]


def check_bucket(bucket: str, timeout: float = 8.0) -> dict:
    url = f"https://{bucket}.s3.amazonaws.com/"
    try:
        r = httpx.get(url, timeout=timeout, follow_redirects=True)
    except Exception as e:
        return {"bucket": bucket, "status": "error", "detail": str(e)}
    if r.status_code == 200:
        body = r.text.lower()
        can_list = "<listbucketresult" in body or "<contents>" in body
        return {"bucket": bucket, "status": "exists", "http": 200, "listable": can_list, "url": url}
    if r.status_code == 403:
        return {"bucket": bucket, "status": "exists", "http": 403, "listable": False, "url": url}
    if r.status_code == 404:
        return {"bucket": bucket, "status": "not_found", "http": 404, "url": url}
    return {"bucket": bucket, "status": "unknown", "http": r.status_code, "url": url}


def enumerate(base: str, timeout: float = 8.0) -> list[dict]:
    results = []
    for pat in PERMUTATIONS:
        bucket = pat.format(base=base)
        results.append(check_bucket(bucket, timeout=timeout))
    return results
