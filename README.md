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

## Available Test Wallets

The following test wallets are preconfigured for use in the project. Each wallet is associated with a UUID and a specific balance (in units):

| Wallet UUID                                 | Balance  |
|--------------------------------------------|----------|
| `103960a0-9a79-43ed-bff0-052a19eaa98e`     | 5,000    |
| `f2bfc720-c67a-4f79-9bcb-47c146deb8e3`     | 10,000   |
| `c6a51a4e-7c5e-4354-8576-d22c7649ad65`     | 15,000   |
| `a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11`     | 20,000   |
| `b81d4fae-7dec-4c8e-8e0e-34d5dd4d88b0`     | 0        |
| `5a82d8af-1b3c-4f07-80f4-1c9be144c87e`     | 7,500    |
| `29f360c5-8466-4dfb-8f94-ec4e004d632f`     | 500      |
| `8e4c1a17-4dbd-4910-a0d1-7e6eced20a5e`     | 100,000  |
| `cad2e9e6-56cc-4aa9-b5c2-8722e0a1c7d7`     | 700      |
| `f4a6c8b1-34d3-4dc6-87b2-9b8a439dcc59`     | 1        |
| `77d7d5b5-54b7-41c3-9887-9a50b80cf6df`     | 92,254,775,807 |

### How to Use

These wallets can be used for testing various API endpoints or application features involving wallet operations. Ensure that the wallet UUIDs are correctly referenced in your requests or test scripts. 

--- 

### Notes:
- Replace `localhost` with the actual IP addres Docker provides.
