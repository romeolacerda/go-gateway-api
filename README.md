# Payment-Gateway

## About the Project

The goal is to demonstrate the construction of a modern distributed system with separation of concerns, asynchronous communication, and real-time fraud analysis.

## Architecture

[View the complete architecture here](https://link.excalidraw.com/readonly/Nrz6WjyTrn7IY8ZkrZHy)

### System Components

- **Frontend (Next.js)**
  - User interface for account management and payment processing
  - Built with Next.js to ensure performance and a good user experience

- **Gateway (Go)**
  - Main system for payment processing
  - Manages accounts, transactions, and coordinates the payment flow
  - Publishes transaction events to Kafka for fraud analysis

- **Apache Kafka**
  - Responsible for asynchronous communication between the API Gateway and the Anti-fraud service
  - Ensures reliable message exchange between services
  - Specific topics for transactions and analysis results

- **Anti-fraud Service (Nest.js)**
  - Consumes transaction events from Kafka
  - Performs real-time analysis to detect potential fraud
  - Publishes analysis results back to Kafka

## Communication Flow

1. The frontend sends REST requests to the API Gateway
2. The Gateway processes the requests and publishes transaction events to Kafka
3. The Anti-fraud service consumes the events and performs real-time analysis
4. The analysis results are published back to Kafka
5. The Gateway consumes the results and finalizes the transaction processing

## Service Startup Order

To run the complete project, services must be started in the following order:

1. **API Gateway (Go)** – Should be started first as it sets up the Docker network
2. **Anti-fraud Service (Nest.js)** – Depends on Kafka, which is configured by the Gateway
3. **Frontend (Next.js)** – The user interface that communicates with the API Gateway

## Detailed Instructions

Each system component contains specific setup and configuration instructions in its respective folder:

- **API Gateway**: See the README in the `/go-gateway-api` folder
- **Anti-fraud Service**: See the README in the `/nestjs-anti-fraud` folder
- **Frontend**: See the README in the `/gateway-fe` folder

> **Important**: It's crucial to follow the startup order above, as each service depends on the previous ones to function correctly.

## General Prerequisites

To run all components of the project, you’ll need:

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- Git

## Key Business Rules

- Transactions over R$ 10,000 are automatically sent for analysis and marked as "pending"
- Smaller transactions are processed immediately
- The interface displays different statuses using colors: green (approved), yellow (pending), red (rejected)
