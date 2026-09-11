#!/bin/bash
# Bug Bounty Hunter - Automated Recon & Scanning Script
set -e
export PATH=$PATH:/root/go/bin
TARGET=""
MODE="full"
usage() {
echo "Usage: $0  [mode]"
echo " domain - Target domain (e.g., example.com)"
echo " mode - recon|scan|fuzz|full (default: full)"
echo ""
echo "Examples:"
echo " $0 example.com # Full workflow"
echo " $0 example.com recon # Passive recon only"
echo " $0 example.com scan # Active scanning"
exit 1
}
if [ $# -lt 1 ]; then
usage
fi
TARGET="$1"
[ $# -ge 2 ] && MODE="$2"
OUTPUT\_DIR="/home/workspace/bugbounty-output/$(date +%Y%m%d\_%H%M%S)\_${TARGET}"
mkdir -p "$OUTPUT\_DIR/subdomains" "$OUTPUT\_DIR/scans" "$OUTPUT\_DIR/fuzz"
echo "=============================================="
echo " Bug Bounty Hunter - $TARGET"
echo " Mode: $MODE"
echo " Output: $OUTPUT\_DIR"
echo "=============================================="
recon() {
echo "[\*] Running subdomain enumeration..."
subfinder -d "$TARGET" -o "$OUTPUT\_DIR/subdomains/subs.txt" 2>/dev/null || true
# Also check for common subdomains
if [ -f "$OUTPUT\_DIR/subdomains/subs.txt" ]; then
echo "[+] Found $(wc -l < $OUTPUT\_DIR/subdomains/subs.txt) subdomains"
else
echo "[-] No subdomains found"
echo "$TARGET" > "$OUTPUT\_DIR/subdomains/subs.txt"
fi
echo "[\*] Probing for alive hosts..."
cat "$OUTPUT\_DIR/subdomains/subs.txt" | httpx -silent -threads 50 -o "$OUTPUT\_DIR/subdomains/alive.txt" 2>/dev/null || true
if [ -f "$OUTPUT\_DIR/subdomains/alive.txt" ]; then
echo "[+] $(wc -l < $OUTPUT\_DIR/subdomains/alive.txt) alive hosts"
else
echo "[-] No alive hosts found"
cp "$OUTPUT\_DIR/subdomains/subs.txt" "$OUTPUT\_DIR/subdomains/alive.txt"
fi
echo "[\*] Saving subdomain list..."
cat "$OUTPUT\_DIR/subdomains/alive.txt" > "$OUTPUT\_DIR/all\_hosts.txt"
echo "$TARGET" >> "$OUTPUT\_DIR/all\_hosts.txt"
}
scan() {
echo "[\*] Running nuclei vulnerability scan..."
if [ ! -f "$OUTPUT\_DIR/all\_hosts.txt" ]; then
echo "[-] No hosts found. Run recon first."
return 1
fi
nuclei -l "$OUTPUT\_DIR/all\_hosts.txt" \
-severity critical,high,medium \
-json \
-o "$OUTPUT\_DIR/scans/nuclei\_results.json" \
2>/dev/null || true
# Also run tech detection
echo "[\*] Running technology detection..."
httpx -l "$OUTPUT\_DIR/all\_hosts.txt" \
-tech-detect \
-json \
-o "$OUTPUT\_DIR/scans/tech\_detect.json" \
2>/dev/null || true
# Port scan top ports
echo "[\*] Running quick port scan..."
nmap -T4 -F --open -oG "$OUTPUT\_DIR/scans/ports.txt" "$TARGET" 2>/dev/null || true
echo "[+] Scan complete. Results in $OUTPUT\_DIR/scans/"
}
fuzz() {
URL="$1"
if [ -z "$URL" ]; then
echo "[-] Fuzz mode requires a URL"
echo "Usage: bugbounty-fuzz https://example.com"
return 1
fi
echo "[\*] Fuzzing $URL"
# Directory fuzzing
echo "[\*] Directory fuzzing..."
ffuf -u "${URL}FUZZ" \
-w /usr/share/wordlists/dirb/common.txt \
-mc 200,204,301,302,307,401,403 \
-o "$OUTPUT\_DIR/fuzz/directories.json" \
2>/dev/null || true
# Parameter fuzzing
echo "[\*] Parameter fuzzing..."
ffuf -u "${URL}?FUZZ=test" \
-w /usr/share/wordlists/dirb/common.txt \
-mc 200,204,301,302,307,401,403 \
-o "$OUTPUT\_DIR/fuzz/parameters.json" \
2>/dev/null || true
echo "[+] Fuzz complete"
}
case "$MODE" in
recon)
recon
;;
scan)
scan
;;
fuzz)
fuzz "$TARGET"
;;
full)
recon
scan
;;
\*)
echo "[-] Unknown mode: $MODE"
usage
;;
esac
echo ""
echo "=============================================="
echo " Complete! Results in: $OUTPUT\_DIR"
echo "=============================================="
echo ""
echo "Files created:"
ls -la "$OUTPUT\_DIR" 2>/dev/null || true