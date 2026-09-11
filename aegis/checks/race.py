import asyncio
import httpx


async def _one(client: httpx.AsyncClient, method: str, url: str, data=None, json=None, headers=None):
    try:
        r = await client.request(method, url, data=data, json=json, headers=headers, timeout=10.0)
        return {"status": r.status_code, "length": len(r.content), "headers": dict(r.headers)}
    except Exception as e:
        return {"error": str(e)}


async def race(method: str, url: str, count: int = 25, data=None, json=None, headers=None, concurrency: int = 25) -> list[dict]:
    limits = httpx.Limits(max_connections=concurrency, max_keepalive_connections=concurrency)
    async with httpx.AsyncClient(limits=limits) as client:
        tasks = [_one(client, method, url, data=data, json=json, headers=headers) for _ in range(count)]
        return await asyncio.gather(*tasks)


def race_sync(*args, **kwargs) -> list[dict]:
    return asyncio.run(race(*args, **kwargs))


def analyze(results: list[dict]) -> dict:
    codes = {}
    for r in results:
        k = r.get("status", "error")
        codes[k] = codes.get(k, 0) + 1
    lengths = sorted({r.get("length") for r in results if "length" in r})
    diverged = len(lengths) > 1
    return {"total": len(results), "by_status": codes, "unique_lengths": lengths, "diverged": diverged}
