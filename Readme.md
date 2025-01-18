# Discount Engine Application

The Discount Engine is a Go-based web application designed to calculate and apply discounts based on predefined rules. It provides a simple and extensible way to handle dynamic discounting logic for orders.

## Features
- Apply discounts based on rules like minimum order value, customer type, and more.
- Supports both percentage-based and fixed discounts.
- Handles multiple discount rules with prioritization.
- Exposes a REST API to process discount calculations.

## Project Structure
```
/discount-engine
├── controllers   # Contains the core logic for processing discounts
├── models        # Defines the data structures for orders and rules
├── main.go       # Application entry point
├── go.mod        # Go module dependencies
├── go_task_rules.json  # JSON file containing discount rules
```

## Prerequisites
- Go 1.22 or later
- Docker (optional for containerized deployment)

## Installation

### Local Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/Ris-Codes/discount-engine.git
   cd discount-engine
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

4. By default, the application runs on port `8000`. You can access the API at `http://localhost:8000`.

### Docker Setup
1. Build the Docker image:
   ```bash
   docker build -t discount-engine .
   ```

2. Run the Docker container:
   ```bash
   docker run -p 8000:8000 discount-engine
   ```

3. The application will be accessible at `http://localhost:8000`.

## API Usage
### Endpoint
- `POST /apply-discount`

### Request Body
```json
{
  "order_total": 120,
  "customer_type": "premium"
}
```

### Response
```json
{
    "discount_amount": 20,
    "fianl_total": 100,
    "applied_rules": [
        "rule_2"
    ]
}
```

## Configuration
### Discount Rules
Discount rules are defined in `go_task_rules.json`. Modify this file to add or update rules. A sample rule looks like:
```json
[
  {
    "id": "rule_1",
    "description": "10% off for orders over $100",
    "condition": {
      "min_order_value": 100
    },
    "discount_percentage": 10
  },
  {
    "id": "rule_2",
    "description": "$20 off for premium customers",
    "condition": {
      "customer_type": "premium"
    },
    "discount_fixed": 20
  },
  {
    "id": "rule_3",
    "description": "5% off for orders over $50",
    "condition": {
      "min_order_value": 50
    },
    "discount_percentage": 5
  },
  {
    "id": "rule_4",
    "description": "$10 off for regular customers on orders over $75",
    "condition": {
      "customer_type": "regular",
      "min_order_value": 75
    },
    "discount_fixed": 10
  }
]
```

### Running Unit Tests
Run the following command to execute unit tests:
```bash
go test ./... -v
```