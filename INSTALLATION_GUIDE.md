# QoreChain Light Node Installation Guide

This guide explains the basic setup flow for running a QoreChain Light Node from this fork.

> Maintained by **Satoshi-Qore** as part of QoreChain community support and infrastructure learning.

---

## 1. Requirements

Before starting, make sure your system has:

- Linux VPS or local Linux environment
- Go installed
- Docker and Docker Compose installed
- Git installed
- Basic terminal access

For VPS usage, Ubuntu is recommended for beginners.

---

## 2. Clone the Repository

```bash
git clone https://github.com/satoshi-Qore/qorechain-lightnode.git
cd qorechain-lightnode
```

---

## 3. Build the Light Node

### SX Edition

The SX edition is designed for server / VPS usage.

```bash
CGO_ENABLED=1 go build -o build/lightnode-sx ./cmd/lightnode-sx/
```

### UX Edition

The UX edition includes a dashboard experience.

```bash
CGO_ENABLED=1 go build -o build/lightnode-ux ./cmd/lightnode-ux/
```

---

## 4. Run Onboarding

```bash
build/lightnode-sx onboard
```

The onboarding wizard may ask for:

- Chain RPC endpoint
- Private key or new key generation
- Node configuration details

If the network endpoint is not available yet, local-only mode can be used for testing.

---

## 5. Start the Node

```bash
build/lightnode-sx start
```

---

## 6. Docker Setup

Start the SX edition:

```bash
docker compose up lightnode-sx
```

Start the UX edition:

```bash
docker compose up lightnode-ux
```

Run in background:

```bash
docker compose up -d
```

---

## 7. Check Logs

```bash
docker compose logs -f
```

---

## 8. Restart Services

```bash
docker compose restart
```

---

## 9. Verify PQC Self-Test

```bash
lightnode-sx selftest
```

This verifies that the post-quantum cryptography stack is working correctly.

---

## 10. Notes

This guide is intended for community learning and support. Always check the official QoreChain documentation and repository for the latest production instructions.

Official repository:

```text
https://github.com/qorechain/qorechain-lightnode
```
