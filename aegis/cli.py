import json
import click
from rich.console import Console
from rich.table import Table
from rich import box

from aegis.checks import s3 as s3check
from aegis.checks import headers as hcheck
from aegis.checks import race as racecheck
from aegis.checks import recon as rcheck

console = Console()


@click.group()
@click.version_option("0.1.0")
def main():
    """Aegis - security audit CLI (Python). Combines bugbounty-findings, kambegoye-scan, and security-research into one tool."""
    pass


@main.command("scan")
@click.argument("target")
@click.option("--json", "as_json", is_flag=True, help="Output JSON")
def scan(target, as_json):
    """Full audit: headers, sensitive files, stack, subdomains, and ports for a URL or domain."""
    url = target if target.startswith("http") else f"https://{target}"
    domain = target.replace("https://", "").replace("http://", "").split("/")[0]

    data = {}
    data["target"] = target
    data["headers"] = hcheck.check_headers(url)
    data["sensitive"] = hcheck.check_sensitive(url)
    data["stack"] = rcheck.detect_stack(url)
    data["subdomains"] = rcheck.enumerate_subdomains(domain)
    data["ports"] = rcheck.scan_ports(domain)

    if as_json:
        click.echo(json.dumps(data, indent=2))
        return

    # Pretty
    console.rule(f"[bold]Aegis scan: {target}")

    t = Table(title="Security headers", box=box.SIMPLE)
    t.add_column("Check"); t.add_column("Result")
    hdr = data["headers"]
    for label in ["HSTS", "CSP", "X-Frame-Options", "X-Content-Type-Options", "Referrer-Policy", "Permissions-Policy"]:
        present = label not in hdr.get("missing", [])
        t.add_row(label, "[green]present" if present else "[red]missing")
    if hdr.get("cookie_issues"):
        for iss in hdr["cookie_issues"]:
            t.add_row("Cookie", f"[yellow]{iss}")
    console.print(t)

    t2 = Table(title="Sensitive files", box=box.SIMPLE)
    t2.add_column("Path"); t2.add_column("Status")
    for row in data["sensitive"]:
        status = str(row.get("status"))
        exposed = "[red]exposed" if row.get("exposed") else "[green]hidden"
        t2.add_row(row["path"], f"{status} {exposed}")
    console.print(t2)

    t3 = Table(title="Stack", box=box.SIMPLE)
    t3.add_column("Signal"); t3.add_column("Value")
    st = data["stack"]
    t3.add_row("Server", st.get("server", "-"))
    t3.add_row("Tech", ", ".join(st.get("tech", [])) or "-")
    console.print(t3)

    t4 = Table(title="Subdomains", box=box.SIMPLE)
    t4.add_column("Host"); t4.add_column("Resolved")
    for row in data["subdomains"]:
        t4.add_row(row["host"], "[green]yes" if row["resolved"] else "[dim]no")
    console.print(t4)

    t5 = Table(title="Ports", box=box.SIMPLE)
    t5.add_column("Port"); t5.add_column("Open")
    for row in data["ports"]:
        t5.add_row(str(row["port"]), "[red]open" if row["open"] else "[dim]closed")
    console.print(t5)


@main.command("s3")
@click.argument("base")
def s3(base):
    """Enumerate S3 buckets from a base name with permutation attack."""
    results = s3check.enumerate(base)
    t = Table(title=f"S3 enumeration for {base}", box=box.SIMPLE)
    t.add_column("Bucket"); t.add_column("HTTP"); t.add_column("Result")
    for r in results:
        status = r["status"]
        color = "green" if status == "exists" and r.get("listable") else "yellow" if status == "exists" else "dim"
        t.add_row(r["bucket"], str(r.get("http", "-")), f"[{color}]{status}[/]")
    console.print(t)


@main.command("headers")
@click.argument("url")
def headers(url):
    """Check security headers and cookie flags for a URL."""
    res = hcheck.check_headers(url)
    console.print_json(data=res)
    sens = hcheck.check_sensitive(url)
    t = Table(title="Sensitive files", box=box.SIMPLE)
    t.add_column("Path"); t.add_column("Status"); t.add_column("Exposed")
    for row in sens:
        t.add_row(row["path"], str(row.get("status")), "[red]yes" if row.get("exposed") else "no")
    console.print(t)


@main.command("race")
@click.argument("url")
@click.option("--method", default="POST", help="HTTP method")
@click.option("--count", default=25, help="Number of parallel requests")
@click.option("--data", default=None, help="Body data")
def race(url, method, count, data):
    """Fire parallel requests to test race conditions (TOCTOU)."""
    import json as _json
    payload = None
    if data:
        try:
            payload = _json.loads(data)
        except Exception:
            payload = data
    # If payload looks like JSON dict, send as json, else data
    kwargs = {"json": payload} if isinstance(payload, dict) else {"data": payload}
    results = racecheck.race_sync(method.upper(), url, count=count, **kwargs)
    summary = racecheck.analyze(results)
    console.print_json(data={"summary": summary, "results": results[:5]})
    if summary["diverged"]:
        console.print("[red]Diverged responses detected - possible race condition[/]")
    else:
        console.print("[green]No divergence in response lengths[/]")


@main.command("recon")
@click.argument("domain")
def recon(domain):
    """Subdomains, ports, and stack detection for a domain."""
    subs = rcheck.enumerate_subdomains(domain)
    ports = rcheck.scan_ports(domain)
    stack = rcheck.detect_stack(f"https://{domain}")
    console.print_json(data={"subdomains": subs, "ports": ports, "stack": stack})


if __name__ == "__main__":
    main()
