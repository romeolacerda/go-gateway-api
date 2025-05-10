# Payment Gateway - Anti-Fraud Microservice (NestJS)

This is the Anti-Fraud microservice developed in NestJS

## About the Project

The Payment Gateway is a distributed system composed of:
- Frontend in Next.js
- API Gateway in Go
- Anti-Fraud System in Nest.js (this repository)
- Apache Kafka for asynchronous communication

## Application architecture
[View the complete architecture here](https://link.excalidraw.com/readonly/Nrz6WjyTrn7IY8ZkrZHy)

## Prerequisites

- [Docker](https://www.docker.com/get-started)
  - For Windows: [WSL2](https://docs.docker.com/desktop/windows/wsl/) it's necessary
- [Extension REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) (optionl, for run tests with api.http)

## Impportant!

⚠️ **You need to run the go-gateway-api service first** before this project, as this microservice uses the Docker network created by go-gateway.

## Porject setup

1. Clone the repository:
```bash
git clone https://github.com/romeolacerda/payment-gateway
cd payment-gateway/nestjs-anti-fraud
```

2. Check if the go-gateway service is already running

3. Start the services with Docker Compose:
```bash
docker compose up -d
```

4. Run Prisma migrations inside the container:
```bash
docker compose exec nestjs bash
npx prisma migrate dev
```

## Running the application

You can run the application in two different modes inside the container:

### 1. API REST + Consumer Kafka (standard)
```bash
docker compose exec nestjs bash
npm run start:dev
```

### 2. Only the Kafka
```bash
docker compose exec nestjs bash
npm run start:dev -- --entryFile cmd/kafka.cmd
```

## Project Structure
The project uses:
- NestJS as a framework
- Prisma ORM for accessing the PostgreSQL database
- Integration with Apache Kafka for asynchronous processing
- TypeScript for static typing

## Communication with Kafka

The Anti-Fraud service communicates with the API Gateway via Apache Kafka:

### Event consumption
- **Topic**: `pending_transactions`
- **Format**: JSON with full transaction data

### Event production
- **Topic**: `transactions_result`
- **Format**: JSON with the analysis result and risk score

## Analysis Rules

The system applies rules to detect possible fraud, such as:

1. **Transaction amount**:
- Transactions above certain thresholds receive a higher risk score

2. **Transaction frequency**:
- Many transactions in a short period of time increase risk

3. **Card behavior**:
- Use of multiple cards with suspicious patterns

## API Endpoints

The project includes an `api.http` file that can be used with the VS Code REST Client extension to test API endpoints:

1. Install the REST Client extension in VS Code
2. Open the `api.http` file
3. Click "Send Request" above each request

The available endpoints are documented in this file for easy testing and reference.

## Database access

The PostgreSQL of the Anti-Fraud service is configured to avoid conflicts with the go-gateway database.

To access Prisma Studio (visual database interface):

```bash
docker compose exec nestjs bash
npx prisma studio
```

## Logs and Monitoring

To view service logs:

```bash
docker logs -f nestjs-anti-fraud-nestjs-1
```

## Development

For development, you can run commands inside the container:

```bash
docker compose exec nestjs bash

# Example of commands available inside the container
npm run start:dev # Start in development mode
npm run start:dev -- --entryFile cmd/kafka.cmd # Start only the Kafka consumer
npx prisma studio # Visual database interface
npx prisma migrate dev # Run migrations
```

To modify the code, edit the files locally - they are mounted as a volume in the Docker container, which automatically restarts when it detects changes.
