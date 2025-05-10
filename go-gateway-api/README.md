# Payment Gateway- API Gateway (Go)

This is the API Gateway microservice developed in Go, part of the Payment Gateway project.

## About the Project

The Payment Gateway is a distributed system composed of:
- Frontend in Next.js
- API Gateway in Go (this repository)
- Anti-fraud system in Nest.js
- Apache Kafka for asynchronous communication

## Current status of the project
So far, we have implemented:
- Setup and base structure of the project
- Account management endpoints (creation and consultation)
- Complete invoice system with:
- Automatic creation and processing of payments
- Limit validation (invoices > R$ 10,000 are pending)
- Individual consultation and listing of invoices
- Automatic update of account balance

Pending features:
- Integration with Apache Kafka for:
- Sending transactions to the anti-fraud microservice
- Consumption of responses from the anti-fraud service
- Payment processing based on fraud analysis


## Application architecture
[View the full architecture here](https://link.excalidraw.com/readonly/Nrz6WjyTrn7IY8ZkrZHy)

## Prerequisites

- [Go](https://golang.org/doc/install) 1.24 or higher
- [Docker](https://www.docker.com/get-started)
  - For Windows: [WSL2](https://docs.docker.com/desktop/windows/wsl/) it's necessary
- [golang-migrate](https://github.com/golang-migrate/migrate)
  - Installation: `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
- [Extensão REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) (optional, for run tests)

## Project Setup

1. Clone the repository:
```bash
git clone https://github.com/romeolacerda/payment-gateway
cd imersao22/go-gateway
```

2. Configure the environment variables:
```bash
cp .env.example .env
```

3. Init the database:
```bash
docker compose up -d
```

4. Run the migrations:
```bash
migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/gateway?sslmode=disable" up
```

5. Access the container of the app:
```bash
docker compose exec app sh
```

6. Run the application:
```bash
go run cmd/app/main.go
```

## API Endpoints

### Create account
```http
POST /accounts
Content-Type: application/json

{
    "name": "John Doe",
    "email": "john@doe.com"
}
```
Returns the data of the created account, including the API Key for authentication.

### Consult Account

```http
GET /accounts
X-API-Key: {api_key}
```
Returns the account data associated with the API Key.

### Create Invoice
```http
POST /invoice
Content-Type: application/json
X-API-Key: {api_key}

{
    "amount": 100.50,
    "description": "Compra de produto",
    "payment_type": "credit_card",
    "card_number": "4111111111111111",
    "cvv": "123",
    "expiry_month": 12,
    "expiry_year": 2025,
    "cardholder_name": "John Doe"
}
```
Creates a new invoice and processes the payment. Invoices over R$10,000 are pending manual review.

### Consult Invoice
```http
GET /invoice/{id}
X-API-Key: {api_key}
```
Returns data for a specific invoice.

### List Invoices
```http
GET /invoice
X-API-Key: {api_key}
```
Lists all invoices for the account.

## Testing the API

The project includes a `test.http` file that can be used with the VS Code REST Client extension. This file contains:
- Pre-configured global variables
- Examples of all requests
- Automatic API Key capture after account creation

To use:
1. Install the REST Client extension in VS Code
2. Open the `test.http` file
3. Click "Send Request" above each request