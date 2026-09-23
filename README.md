# learn-go-http-server

This repository is a beginner-friendly Go HTTP server that demonstrates routing, status codes, query parameters, JSON responses, and method handling.

## What this project covers

* Creating an HTTP server with the standard library
* Using `http.ServeMux` for route registration
* Handling GET and POST requests
* Returning different status codes
* Serving plain text and JSON responses
* Reading query parameters
* Using 404 and 405 responses

## Routes

| Route | Method | Response |
| --- | --- | --- |
| `/` | GET | Root page |
| `/home` | GET | Welcome message |
| `/about` | GET | About page |
| `/health` | GET | JSON: `{ "status": "ok" }` |
| `/greet?name=Alice` | GET | Hello message with query param |
| `/api/users` | GET | JSON list of users |
| `/api/users` | POST | Creates a new user |
| `/api/users/1` | GET | User detail page |
| Undefined route | GET | `404 Not Found` |
| Any route with unsupported method | Any | `405 Method Not Allowed` |

## Run the server

```bash
go run .
```

The server starts on port `:8080` by default. You can override it with:

```bash
PORT=9000 go run .
```

## Example requests

```bash
curl http://localhost:8080/
curl http://localhost:8080/home
curl http://localhost:8080/about
curl http://localhost:8080/health
curl "http://localhost:8080/greet?name=Alice"
curl http://localhost:8080/api/users
curl -X POST -d "name=Charlie" http://localhost:8080/api/users
curl http://localhost:8080/api/users/1
```

## Notes

This project is meant for learning. It keeps the code simple while showing practical patterns you will use when building APIs or web apps in Go.

```

Expected:

```text
HTTP/1.1 200 OK

Welcome
```

---

### Method Not Allowed

```bash
curl -i -X POST http://localhost:3000/home
```

Expected:

```text
HTTP/1.1 405 Method Not Allowed

Method Not Allowed
```

---

### Route Not Found

```bash
curl -i http://localhost:3000/about
```

Expected:

```text
HTTP/1.1 404 Not Found

404 page not found
```

---

## Concepts Demonstrated

* HTTP Server Creation
* Routing with `ServeMux`
* Request Handling
* HTTP Methods
* Status Codes
* Error Handling
* Response Writing

---

## Technologies Used

* Go
* Standard Library (`net/http`)

No third-party packages are required.

---

## Learning Goals

This project is intended for beginners learning backend development with Go.

After understanding this project, consider learning:

* JSON Responses
* Request Bodies
* Query Parameters
* URL Parameters
* REST APIs
* Middleware
* Authentication

## Author

Created by Chekus-dev as a learning project for understanding HTTP servers and status codes in Go.

