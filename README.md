# Technical Assignment: Wallet Management System

## 🚀 Features

- API for wallet operations (`deposit` and `withdraw`).
- Integrated database for secure data persistence.
- Docker Compose support for seamless setup.
- Includes Prometheus metrics.
- Load testing with `k6`.

---

## 📚 Table of Contents

1. [Getting Started](#getting-started)
2. [Installation](#installation)
3. [Usage](#usage)

---

## 🛠️ Getting Started 

This repository is part of a technical assignment and demonstrates the implementation of a wallet management system. Follow the steps below to set up the project locally or in a containerized environment.

---

## 🖥️ Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/wallet-management-system.git
   cd wallet-management-system
   ```
2. **Build and run the application using Docker Compose:**

   ```bash
   make compose
   ```

   This will:
   - Build the application and its dependencies.
   - Start the API server, PostgreSQL database, Prometheus.

4. **Access the services:**
   - API: [http://localhost:8080](http://localhost:8080)
   - Prometheus: [http://localhost:8080/metrics](http://localhost:9090)
---

## 📄 Usage

Once the application is running, you can interact with the wallet API. Below are the main endpoints:

- **GET /api/v1/wallets/{wallet_id}**: Retrieve wallet details.
- **POST /api/v1/wallet**: Perform a `deposit` or `withdraw`.

### Example: Make a deposit

```bash
curl -X POST http://localhost:8080/api/v1/wallet \
  -H "Content-Type: application/json" \
  -d '{
    "id": "103960a0-9a79-43ed-bff0-052a19eaa98e",
    "amount": 1000,
    "operation": "deposit"
  }'
```

---

🚀 If you have any issues, feel free to submit a ticket or open a pull request.

--- 

### Notes:
- Replace `localhost` with the actual IP addres Docker provides.
