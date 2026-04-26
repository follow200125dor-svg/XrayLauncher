# 🚀 Xray Launcher — VPN Manager on Go

A fast and simple VPN launcher written in Go. Downloads and runs Xray-core automatically with your config.

---

## 📋 Features

| Feature | Description |
|---------|-------------|
| ⚡ **Fast** | Written in Go, compiles to a single EXE file |
| 📥 **Auto-download** | Downloads Xray-core automatically on first run |
| 🔧 **Config support** | Works with VLESS, VMess, Trojan, Shadowsocks configs |
| 🪶 **Lightweight** | No dependencies, one executable file |
| 🎨 **Cross-platform** | Works on Windows, Linux, macOS |

---

## 🛠 How to Install

### Method 1: Build from source (recommended)

1. Install Go from [go.dev/dl](https://go.dev/dl)
2. Clone the repository:
git clone https://github.com/YOUR_USERNAME/XrayLauncher.git
cd XrayLauncher

text
3. Build:
go build -o XrayLauncher.exe main.go

text
4. Run:
XrayLauncher.exe

text

### Method 2: Use pre-built EXE

1. Download `XrayLauncher.exe` from [Releases](https://github.com/YOUR_USERNAME/XrayLauncher/releases)
2. Double-click to run

---

## 🚀 How to Use

### First run
1. Launch `XrayLauncher.exe`
2. It will automatically download `xray.exe` from GitHub
3. Unzip and prepare everything

### With config
1. Create `config.json` with your VLESS/VMess config
2. Place it next to `XrayLauncher.exe`
3. Launch — VPN starts working

### Without config
1. Find a free config online (search for "free VLESS config")
2. Save as `config.json` in the same folder
3. Launch `XrayLauncher.exe`

---

## 📁 Project Structure
XrayLauncher/
├── main.go ← main source code
├── go.mod ← Go module file
├── go.sum ← Go dependencies
├── config.json ← your VPN config (you create this)
├── README.md ← this file
└── install.cmd ← auto-installer for Windows

text

---

## ⚙️ Config Example

Create `config.json` with this structure:

```json
{
  "inbounds": [{
    "port": 10808,
    "listen": "127.0.0.1",
    "protocol": "socks",
    "settings": {
      "udp": true
    }
  }],
  "outbounds": [{
    "protocol": "vless",
    "settings": {
      "vnext": [{
        "address": "server.com",
        "port": 443,
        "users": [{
          "id": "your-uuid-here",
          "encryption": "none"
        }]
      }]
    },
    "streamSettings": {
      "network": "ws",
      "security": "tls",
      "wsSettings": {
        "path": "/"
      }
    }
  }]
}
📊 How It Works
text
[Launch] → [Check xray.exe] → [Download if missing] → [Unzip] → [Run with config] → [VPN ON]
❓ FAQ
Q: Is it free?
A: Yes! The launcher is completely free. You can use free configs or your own server.

Q: Does it work in Russia?
A: Yes, with VLESS/XTLS protocol and a working config — YouTube and other sites open.

Q: Do I need a server?
A: No, you can use free public configs. Or set up your own server for better speed.

Q: Is it safe?
A: The launcher is open source. Xray-core is trusted by millions. Always use configs from reliable sources.

Q: Can I use it on macOS/Linux?
A: Yes! Go compiles for all platforms. Just build with go build.

🛠 Technologies
Go — programming language

Xray-core — VPN engine

VLESS / VMess / XTLS — modern protocols for bypassing censorship

👤 Author
Mark, 10 years old

Learning Go and building tools for the world. One project at a time. 🔥

📄 License
Free to use, modify, and share.
