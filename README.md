# Todo API

A simple RESTful API for managing a list of todos. Built using [Gin Web Framework](https://github.com/gin-gonic/gin).

## Features

- View all todos
- Create a new todo
- Get a specific todo by ID
- Update a todo's completion status
- Delete a todo

---

## Installation

### Prerequisites

- Go 1.19 or later installed on your machine.

### Steps

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd <repository-name>
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

4. Open your browser or API client and access the API at:
   ```
   http://localhost:9090
   ```

---

## Endpoints

### Base URL

`http://localhost:9090`

### Endpoints Overview

| Method | Endpoint     | Description                            |
| ------ | ------------ | -------------------------------------- |
| GET    | `/todos`     | Retrieve all todos                     |
| GET    | `/todos/:id` | Retrieve a specific todo by ID         |
| POST   | `/todos`     | Create a new todo                      |
| PUT    | `/todos/:id` | Toggle the completion status of a todo |
| DELETE | `/todos/:id` | Delete a specific todo by ID           |
