# Todo Go API

## Dependencies
* **Chi** for HTTP server
* **Godotenv** for .env
* **PQ** because PostgreSQL database

## API Setup
* create database with Docker with `docker run --name postgres-db-whatever -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres`
* run stuff with `go run .`

## Endpoints
### /todos
For todos - duuuuh.
(This was just in case we'd like to add some other endpoint)

| METHOD | PARAMS | BODY | DESCRIPTION     |
|--------|--------|------|-----------------|
| GET    | -      | -    | Get all todos   |
| GET    | id     | -    | Get single todo |
| POST   | -      | TODO | Create todo     |
| PUT    | -      | TODO | Edit todo       |
| DELETE | id     | -    | Delete todo     |

## Types

### TODO
| KEY     | TYPE   | JSON    |
|---------|--------|---------|
| ID      | in64   | id      |
| Label   | string | label   |
| Checked | bool   | checked |

### TODO CREATE
_Specific for creation JSON to handle postgres ID auto-increment_

| KEY     | TYPE   | JSON    |
|---------|--------|---------|
| Label   | string | label   |
| Checked | bool   | checked |

