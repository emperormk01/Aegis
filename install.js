#!/usr/bin/env node
const fs = require("fs");
const path = require("path");
const https = require("https");
const { execSync } = require("child_process");

const REPO = "emperormk01/Aegis";
const BINARY = "aegis";

function detect() {
  const os = process.platform === "win32" ? "windows" : process.platform === "darwin" ? "darwin" : "linux";
  const arch = process.arch === "arm64" ? "arm64" : "x64";
  return { os, arch };
}

function assetName(os, arch) {
  // matches Go release assets: aegis-linux-x64, etc. but Go uses x64 naming as x64 not amd64
  // our Go workflow uses memrie naming; for Aegis we use aegis-
  return `aegis-${os}-${arch}`;
}

async function download(url, dest) {
  return new Promise((resolve, reject) => {
    const file = fs.createWriteStream(dest);
    https.get(url, { headers: { "User-Agent": "aegis-installer" } }, (res) => {
      if (res.statusCode === 302 || res.statusCode === 301) {
        // follow redirect
        https.get(res.headers.location, (res2) => {
          res2.pipe(file);
          file.on("finish", () => file.close(resolve));
        }).on("error", reject);
        return;
      }
      if (res.statusCode !== 200) {
        reject(new Error(`Download failed ${res.statusCode} for ${url}`));
        return;
      }
      res.pipe(file);
      file.on("finish", () => file.close(resolve));
    }).on("error", reject);
  });
}

async function main() {
  const { os, arch } = detect();
  const asset = assetName(os, arch);
  const url = `https://github.com/${REPO}/releases/latest/download/${asset}`;
  const binDir = path.join(__dirname, "bin");
  const binPath = path.join(binDir, BINARY + (os === "windows" ? ".exe" : ""));

  // Skip if binary already exists and is executable
  if (fs.existsSync(binPath)) {
    try { fs.accessSync(binPath, fs.constants.X_OK); console.log(`[aegis] binary already exists at ${binPath}`); return; } catch {}
  }

  console.log(`[aegis] Downloading ${asset} from ${url}...`);
  fs.mkdirSync(binDir, { recursive: true });

  try {
    await download(url, binPath);
    fs.chmodSync(binPath, 0o755);
    console.log(`[aegis] Installed to ${binPath}`);
  } catch (e) {
    console.error(`[aegis] Download failed: ${e.message}`);
    console.error(`[aegis] Try: go install github.com/${REPO}@latest`);
    // Do not fail postinstall - allow user to use go install fallback
  }
}

main();
