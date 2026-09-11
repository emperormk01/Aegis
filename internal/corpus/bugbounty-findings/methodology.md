# Bug Bounty Methodology

## 1. S3 Bucket Enumeration

**Classic Check:**
```bash
aws s3 ls s3://target-bucket-name --no-sign-request
```

**Permutation Attack:**
Loop through common bucket names:
- target-static
- target-assets
- target-uploads
- target-logs
- target-backup-2026
- target-prod
- target-dev

**Key Misconfigurations:**
- Public ACL + ListBucket + GetObject = free data exfil
- Write access = upload malicious files

**Tools:**
- lazys3
- s3scanner
- AWS CLI

---

## 2. Race Condition Exploitation

**What to test:**
- Create account twice → negative balance
- Apply coupon + checkout simultaneously
- Password reset + change email race
- 2FA disable + login race
- Wallet balance overflow
- Time-of-check to time-of-use (TOCTOU)

**Python + asyncio example:**
```python
import asyncio, httpx

async def race():
    async with httpx.AsyncClient() as c:
        tasks = [c.post(url, data=payload) for _ in range(50)]
        return await asyncio.gather(*tasks, return_exceptions=True)

asyncio.run(race())
```

**Pro Tools:**
- racepwn (my fork of turbo-intruder)
- Turbo Intruder with 0.01s delay clusters
- Python asyncio + httpx for custom races

**Billion Dollar Bug Pattern:**
Multiple requests hitting payment/balance endpoints simultaneously can bypass:
- Rate limiting
- Validation logic
- Balance checks

**Key targets:**
- Financial transactions
- Coupon/stripe usage
- Account creation limits
- OTP verification
- Session management