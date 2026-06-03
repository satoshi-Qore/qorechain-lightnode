# 🐙 QoreChain Light Node

> Community-maintained fork by **Satoshi-Qore** for exploring, running, and supporting QoreChain light node infrastructure.

![QoreChain](https://img.shields.io/badge/QoreChain-Light%20Node-blueviolet)
![Post Quantum](https://img.shields.io/badge/Post--Quantum-Ready-success)
![Docker](https://img.shields.io/badge/Docker-Supported-blue)
![License](https://img.shields.io/badge/License-Apache%202.0-green)

---

## 🚀 Overview

QoreChain Light Node is a lightweight client for the QoreChain network.

It is designed for users who want to contribute to network infrastructure without running a full validator. A light node can relay traffic, serve light client queries, contribute uptime, and participate in the QoreChain ecosystem.

This fork is focused on:

- Helping new users understand light node setup
- Testing the SX and UX editions
- Documenting community installation steps
- Supporting QoreChain node operators
- Tracking useful commands and troubleshooting notes

---

## 🧩 Editions

QoreChain Light Node provides two editions:

| Edition | Name | Best For |
|---|---|---|
| **SX** | Server eXperience | VPS / server deployments |
| **UX** | User eXperience | Dashboard-based desktop or web use |

---

## ✨ Key Features

- Header verification via skipping verification light client
- Delegated staking with multi-validator split
- Auto-compound rewards with configurable intervals
- Reputation-aware validator rebalancing
- Real-time network telemetry
- On-chain registration with heartbeat liveness proofs
- 3% block reward eligibility for active light nodes
- Post-quantum cryptography support
- Interactive onboarding wizard
- Local-only mode for pre-mainnet testing
- Live PQC self-test command

---

## ⚡ Quick Start

### 1. Build SX edition

```bash
CGO_ENABLED=1 go build -o build/lightnode-sx ./cmd/lightnode-sx/
```

### 2. Run onboarding

```bash
build/lightnode-sx onboard
```

During onboarding, the wizard asks for:

- Chain RPC endpoint
- Private key or new key generation
- Node configuration preferences

### 3. Start the node

```bash
build/lightnode-sx start
```

---

## 🐳 Docker Usage

```bash
docker compose up lightnode-sx
```

For the UX dashboard edition:

```bash
docker compose up lightnode-ux
```

---

## 🛡️ Verify PQC Stack

Run the self-test command anytime:

```bash
lightnode-sx selftest
```

This checks:

- Key generation
- Sign operation
- Valid signature verification
- Rejection of tampered signatures
- Rejection of tampered messages

---

## 🛠️ Useful VPS Commands

Check running containers:

```bash
docker ps
```

Start services:

```bash
docker compose up -d
```

View logs:

```bash
docker compose logs -f
```

Restart services:

```bash
docker compose restart
```

---

## 📌 Community Notes

This repository is part of my QoreChain learning and community support work.

I use this fork to:

- Test light node setup flows
- Collect useful commands
- Help community members with node questions
- Improve my understanding of post-quantum blockchain infrastructure

---

## 🔗 Official Repository

Original project:

```text
https://github.com/qorechain/qorechain-lightnode
```

---

## 📄 License

Apache 2.0
