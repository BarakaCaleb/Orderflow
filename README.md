# Modular OrderFlow Platform
Modern, Overengineered Full-Stack System for Learning and Scalability

Tech Stack: Go • Django • React • TailwindCSS • ShadCN UI • PostgreSQL • Redis • Docker • OpenAI API

## 📖 Overview

OrderFlow is a modular, event-driven order management and inventory system designed to simulate production-level architecture for learning and experimentation.
It integrates high-performance Go microservices with Django for authentication, a React-based UI, and OpenAI-assisted documentation.


| Layer            | Technology                   | Purpose                      |
| ---------------- | ---------------------------- | ---------------------------- |
| Frontend         | React + Tailwind + ShadCN UI | Web interface                |
| API              | Django Rest Framework        | Auth, user management, admin |
| Backend Services | Go                           | Order, Inventory, Workers    |
| Database         | PostgreSQL                   | Persistent storage           |
| Cache            | Redis                        | Queueing, caching            |
| Docs             | OpenAI API                   | API doc generation           |
| Containerization | Docker + Compose             | Deployment consistency       |




/backend
│
├── Api/                       # Core Go backend (main API)
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── auth/             # JWT auth, middleware, password hashing
│   │   ├── orders/           # CRUD + order workflow logic
│   │   ├── payments/         # Stripe, M-Pesa, PayPal integration
│   │   ├── vendors/          # Vendor management
│   │   ├── products/         # Inventory management
│   │   ├── notifications/    # Email/SMS/WebSocket
│   │   ├── db/               # DB connection + migrations
│   │   ├── middleware/       # Auth, rate limiting, logging
│   │   └── utils/            # Shared helpers (env loader, etc.)
│   ├── pkg/
│   │   └── response/         # Standardized JSON responses
│   ├── go.mod
│   ├── go.sum
│   └── Dockerfile
│
├── admin/                   # Django admin + analytics + AI
│   ├── orderflow_ai/
│   │   ├── ai_engine/        # OpenAI integration, analytics logic
│   │   ├── reports/          # Reporting + insights
│   │   ├── dashboards/       # Custom dashboards
│   │   ├── notifications/    # Email alerts, scheduled tasks
│   │   └── __init__.py
│   ├── manage.py
│   └── Dockerfile
│
├── db/
│   ├── migrations/           # SQL migration scripts
│   └── init.sql              # Initial DB setup
│
├── redis/
│   └── redis.conf
│
├── docker/
│   ├── go.Dockerfile
│   ├── django.Dockerfile
│   ├── postgres.Dockerfile
│   ├── redis.Dockerfile
│   └── nginx.Dockerfile
│
├── docker-compose.yml
├── Makefile
└── .env.example


## ⚙️ PART 1 — Setting Up the Development Environment (Windows)

You’ll be running a multi-language stack (Go + Python + Node), so you need clean isolation, reproducible builds, and Docker for parity with your collaborator.

# 🧰 Step 1: Install Core Tools

## 1️⃣ Windows Subsystem for Linux (WSL2)
You must develop inside Linux for this stack — Go, Django, and Docker behave inconsistently on pure Windows.

Setup:
```
wsl --install
```

Choose Ubuntu 24.04 LTS when prompted.
Restart your system and open Ubuntu terminal.

## Then update packages:
```
sudo apt update && sudo apt upgrade -y
```
## 2️⃣ Install Docker Desktop

Download: https://www.docker.com/products/docker-desktop

## Enable WSL2 integration for Ubuntu.

Confirm Docker works inside WSL:
```
docker run hello-world
```

## 3️⃣ Install Core Languages
🐍 Python (for Django)
```
sudo apt install python3 python3-pip python3-venv -y
```
🦫 Go
```
sudo snap install go --classic
```

## Verify:
```
go version
```
🧱 Node.js + npm (for React)
```
sudo apt install nodejs npm -y
```
Then:
```
npm install -g pnpm
```
(pnpm is cleaner and faster than npm/yarn)

## 4️⃣ Install PostgreSQL + Redis (for local dev)

(Optional — if you won’t use Docker Compose yet)
```
sudo apt install postgresql redis-server -y
sudo service postgresql start
sudo service redis-server start
```

## 5️⃣ Install Git and Configure It
```
sudo apt install git -y
git config --global user.name "YourGithubName"
git config --global user.email "your@email.com"
```

## Then connect to GitHub using SSH:
```
ssh-keygen -t ed25519 -C "your@email.com"
```

```
cat ~/.ssh/id_ed25519.pub
```

Copy the output and add it under GitHub → Settings → SSH Keys.

## 🐳 Step 2: Run the Project with Docker Compose

Once the repo is cloned:
```
git clone git@github.com:BarakaCaleb/orderflow.git
```
```
cd orderflow
```
```
docker compose up --build
```

### This launches:

`Django API`

`Go services`

`PostgreSQL`

`Redis`

`React frontend`

## Access on:

[Frontend](http://localhost:3000)

[Django API](http://localhost:8000)

## Follow the following branching conventions on this code base
`main` → production-ready

`feature/*` → new features

`bugfix/*`  → bug fixes

`hotfix/*`  → urgent prod fixes

# 📜 License

MIT © 2025 Caleb Baraka