# Over-Engineered Calculator API

A RESTful calculator service with history, separation of logic, and authentication. Built in Go with best-effort microservice architecture.

---

## 🧮 Supported Operations

The `/calculate` endpoint evaluates mathematical expressions using standard infix notation:

| Operation      | Symbol | Example   | Result |
| -------------- | ------ | --------- | ------ |
| Addition       | `+`    | `2+3`     | 5      |
| Subtraction    | `-`    | `5-2`     | 3      |
| Multiplication | `*`    | `4*2`     | 8      |
| Division       | `/`    | `8/2`     | 4      |
| Parentheses    | `( )`  | `(2+3)*5` | 25     |

**Notes:**

* Spaces are ignored (`2 + 3` = `2+3`)
* Division uses floating-point precision (`5/2` → 2.5)
* Invalid characters return an error JSON response:

```json
{ "error": "invalid character in expression" }
```

---

## 🔐 Authentication

All endpoints require **Basic Auth (email/password)**.

**Demo credentials:**

* Email: `user@example.com`
* Password: `123456`

### Using Postman

1. Open the request (e.g., `/calculate`)
2. Go to **Authorization → Basic Auth**
3. Enter the credentials above
4. Send the request

---

## 🚀 Endpoints

### POST `/calculate`

Evaluate a mathematical expression.

**Request Body:**

```json
{
  "expression": "(2+3)*4"
}
```

**Response Body:**

```json
{
  "result": 20
}
```

---

### GET `/history`

Retrieve the list of previous calculations.

**Response Body:**

```json
[
  { "expression": "(2+3)*4", "result": 20 },
  { "expression": "5/2", "result": 2.5 }
]
```

---

## 💻 Testing with curl

**Calculate expression:**

```bash
curl -u user@example.com:123456 \
     -X POST \
     -H "Content-Type: application/json" \
     -d '{"expression":"(2+3)*4"}' \
     http://localhost:8080/calculate
```

**Get history:**

```bash
curl -u user@example.com:123456 \
     http://localhost:8080/history
```

---

## 📦 Postman Collection

A Postman collection is provided for easy testing:

📄 [`postman_collection.json`](./postman_collection.json)

**How to use:**

1. Open [Postman Web](https://web.postman.co) or the desktop app
2. Click **Import → File**
3. Select `postman_collection.json`
4. Set environment variable `base_url` if needed (`http://localhost:8080` for local testing)
5. Run the requests (`/calculate` and `/history`)

---

## ⚙️ Architecture Overview

* **`cmd/server`** → Entry point for the API
* **`internal/api`** → Handlers and routers + authentication
* **`internal/calculator`** → Core calculator logic (lexer, parser (AST construction), evaluator)
* **`internal/storage`** → In-memory history store

This modular structure follows **best-effort microservice principles**, with clear separation of concerns and testable components.

---

## 🔄 Example Workflow

1. Send a POST request to `/calculate` with an expression: `(2+3)*4`
2. Receive result: `{ "result": 20 }`
3. Send a GET request to `/history`
4. Receive list of previous calculations, including the one just executed

## This demonstrates how the calculator and history endpoints work together.
