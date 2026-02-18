# 🛡️ Linux Security Monitoring System (Go Agent + AWS + Dashboard)

A lightweight Linux endpoint security monitoring system built using **Golang** and **AWS Serverless services**.
The system collects installed packages and performs **CIS Benchmark (Ubuntu 22.04 LTS Level 1)** security checks, then sends results securely to AWS and displays them in a web dashboard.

---

## 📌 Project Features

### ✅ Linux Agent (Golang)

* Collects host information:

  * Hostname
  * OS version
  * Kernel version
  * IP address
* Collects installed packages (max 200 packages)
* Performs **10 CIS security compliance checks**
* Sends results to AWS API Gateway in JSON format
* Packaged as a `.deb` installer
* Runs automatically every hour using **systemd timer**

---

### ✅ AWS Serverless Backend

* API Gateway endpoints:

  * `POST /ingest` → store scan report
  * `GET /hosts` → list monitored hosts
  * `GET /latest?hostname=<host>` → fetch latest scan report
* Lambda Functions:

  * `LinuxAgentIngest` → stores reports into DynamoDB
  * `LinuxAgentAPI` → retrieves data from DynamoDB
* DynamoDB Table:

  * `LinuxAgentReports`
  * Partition Key: `hostname`
  * Sort Key: `timestamp`

---

### ✅ Frontend Dashboard

* Hosted using Nginx on EC2 (port 80)
* Displays:

  * Host details
  * CIS compliance score (%)
  * CIS check results (Pass/Fail + Evidence)
  * Installed packages list

---

# 🏗️ Architecture

```
┌───────────────────────────┐
│     Linux Agent (Go)      │
│ - Collect Host Info       │
│ - Collect Packages        │
│ - Run CIS Checks          │
│ - Send JSON Report        │
└─────────────┬─────────────┘
              │ HTTPS POST /ingest
              ▼
┌───────────────────────────┐
│        API Gateway         │
│      POST /ingest          │
└─────────────┬─────────────┘
              ▼
┌───────────────────────────┐
│     Lambda Ingest API      │
│ - Parse JSON report        │
│ - Store into DynamoDB      │
└─────────────┬─────────────┘
              ▼
┌───────────────────────────┐
│         DynamoDB           │
│ Table: LinuxAgentReports   │
│ PK: hostname               │
│ SK: timestamp              │
└─────────────┬─────────────┘
              ▲
              │ Query latest / scan hosts
┌─────────────┴─────────────┐
│       Lambda Query API     │
│ - GET /hosts               │
│ - GET /latest?hostname=X   │
└─────────────┬─────────────┘
              ▲
              │ HTTPS GET requests
┌─────────────┴─────────────┐
│        API Gateway         │
│   GET /hosts, GET /latest  │
└─────────────┬─────────────┘
              ▲
              │ Fetch JSON
┌─────────────┴─────────────┐
│   Web Dashboard (HTML)  │
│ - Host dropdown            │
│ - CIS compliance score     │
│ - Packages + evidence      │
└───────────────────────────┘

```

---

# 🔍 CIS Benchmark Checks Implemented (Ubuntu 22.04 LTS - Level 1)

| Check ID           | CIS Check                                |
| ------------------ | ---------------------------------------- |
| CIS-SSH-ROOT       | Root login disabled over SSH             |
| CIS-FW-UFW         | Firewall enabled (UFW)                   |
| CIS-TIME-SYNC      | Time synchronization configured (chrony) |
| CIS-AUDITD         | Auditd service running                   |
| CIS-APPARMOR       | AppArmor enabled                         |
| CIS-PASS-EXP       | Password expiration policy enforced      |
| CIS-PASS-COMPLEX   | Password complexity policy enabled       |
| CIS-WORLD-WRITABLE | No world-writable files in /tmp          |
| CIS-CRAMFS         | cramfs filesystem disabled               |
| CIS-GDM-AUTOLOGIN  | GDM auto-login disabled                  |

---

# 📂 Project Structure

```
linux-agent-project/
├── cmd/
│   └── main.go
├── internal/
│   ├── models/
│   │   └── report.go
│   ├── collector/
│   │   ├── host.go
│   │   └── packages.go
│   ├── checks/
│   │   ├── cis.go
│   │   ├── ssh_root.go
│   │   ├── firewall.go
│   │   ├── time_sync.go
│   │   ├── auditd.go
│   │   ├── apparmor.go
│   │   ├── password_expiry.go
│   │   ├── password_complexity.go
│   │   ├── world_writable.go
│   │   ├── cramfs.go
│   │   ├── gdm_autologin.go
│   │   └── all_checks.go
│   └── sender/
│       └── aws_sender.go
├── frontend/
│   └── index.html
├── go.mod
└── go.sum
```

---

# 🚀 Setup & Deployment Guide

## 1️⃣ Clone Repository

```bash
git clone https://github.com/Akhil-jagadale/security-monitoring-system.git
cd security-monitoring-system
```

---

## 2️⃣ Install Dependencies (Ubuntu 22.04)

```bash
sudo apt update -y
sudo apt install -y golang-go curl ufw auditd chrony
```

Verify:

```bash
go version
```

---

# 🖥️ Running the Agent Manually

```bash
go run cmd/main.go
```

Or build binary:

```bash
go build -o linux-agent cmd/main.go
./linux-agent
```

---

# 📦 Building the .deb Package

Build binary:

```bash
go build -o linux-agent cmd/main.go
```

Create packaging structure:

```bash
mkdir -p package/linux-agent/usr/local/bin
mkdir -p package/linux-agent/DEBIAN
cp linux-agent package/linux-agent/usr/local/bin/
```

Create control file:

```bash
nano package/linux-agent/DEBIAN/control
```

Paste:

```
Package: linux-agent
Version: 1.0
Section: base
Priority: optional
Architecture: amd64
Maintainer: Akhilesh Jagadale
Description: Lightweight Linux CIS Security Agent
```

Build package:

```bash
dpkg-deb --build package/linux-agent
```

Install package:

```bash
sudo dpkg -i package/linux-agent.deb
```

Verify installation:

```bash
dpkg -l | grep linux-agent
```

Run agent:

```bash
linux-agent
```

---

# ⏰ Running Agent Automatically (systemd Timer)

### Create systemd service

```bash
sudo nano /etc/systemd/system/linux-agent.service
```

Paste:

```ini
[Unit]
Description=Linux Security Monitoring Agent
After=network.target

[Service]
Type=oneshot
ExecStart=/usr/local/bin/linux-agent
User=root
```

---

### Create systemd timer

```bash
sudo nano /etc/systemd/system/linux-agent.timer
```

Paste:

```ini
[Unit]
Description=Run Linux Agent every 1 hour

[Timer]
OnBootSec=2min
OnUnitActiveSec=1h
Unit=linux-agent.service

[Install]
WantedBy=timers.target
```

Enable timer:

```bash
sudo systemctl daemon-reload
sudo systemctl enable linux-agent.timer
sudo systemctl start linux-agent.timer
```

Verify timer:

```bash
sudo systemctl list-timers --all | grep linux-agent
```

Check logs:

```bash
sudo journalctl -u linux-agent.service -n 50 --no-pager
```

---

# ☁️ AWS Setup

## DynamoDB Table

Table Name:

```
LinuxAgentReports
```

Partition Key:

```
hostname (String)
```

Sort Key:

```
timestamp (String)
```

---

## API Gateway Endpoints

### POST report ingestion

```
POST /ingest
```

### List monitored hosts

```
GET /hosts
```

### Fetch latest scan

```
GET /latest?hostname=<hostname>
```

---

# 🌐 Hosting Frontend Dashboard on EC2 (Nginx)

Install nginx:

```bash
sudo apt update -y
sudo apt install nginx -y
```

Copy frontend:

```bash
sudo rm -rf /var/www/html/*
sudo cp -r frontend/* /var/www/html/
sudo systemctl restart nginx
```

Allow port 80 in EC2 Security Group.

Access dashboard:

```
http://<EC2_PUBLIC_IP>
```

---

# 🔐 Security Notes (Production Improvements)

For simplicity, APIs are open for demo purposes.
In a production environment, security improvements would include:

* API Gateway authentication (API Key / IAM / Cognito)
* TLS mutual authentication for agent communication
* Encrypt DynamoDB with KMS CMK
* Store logs in CloudWatch + alerts
* Use least privilege IAM policies

---

# 🛠️ Future Enhancements

* Multi-host support with filtering/search
* Add more CIS checks (20+)
* Add CloudWatch monitoring dashboard
* Add authentication for frontend
* Store package data separately for faster querying
* Add remediation suggestions for failed CIS checks

---

# 👨‍💻 Author

**Akhilesh Jagadale**
GitHub: [https://github.com/Akhil-jagadale](https://github.com/Akhil-jagadale)
