#!/usr/bin/env node
const { spawn } = require("child_process");
const path = require("path");
const fs = require("fs");

const bin = path.join(__dirname, process.platform === "win32" ? "aegis.exe" : "aegis");

function run() {
  if (!fs.existsSync(bin)) {
    console.error(`[aegis] Binary not found at ${bin}. Try reinstalling or use: go install github.com/emperormk01/Aegis@latest`);
    process.exit(1);
  }
  const child = spawn(bin, process.argv.slice(2), { stdio: "inherit" });
  child.on("exit", (code) => process.exit(code ?? 0));
}

run();
