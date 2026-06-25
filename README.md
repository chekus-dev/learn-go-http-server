# learn-go-http-server
i created this repository to teach go-http-server about status code and pattern macthing
# Go HTTP Server with Status Codes

A beginner-friendly Go project that demonstrates how to build a simple HTTP server using the standard `net/http` package.

This project covers:

* Creating an HTTP server
* Using `http.ServeMux` for routing
* Handling multiple routes
* Restricting HTTP methods
* Returning HTTP status codes
* Handling `404 Not Found`
* Handling `405 Method Not Allowed`

## Features

### Routes

| Route               | Method                    | Response                 |
| ------------------- | ------------------------- | ------------------------ |
| `/`                 | GET                       | `This is the root page`  |
| `/home`             | GET                       | `Welcome`                |
| `/`                 | Any method other than GET | `405 Method Not Allowed` |
| `/home`             | Any method other than GET | `405 Method Not Allowed` |
| Any undefined route | GET                       | `404 Not Found`          |

---

## Project Structure

```text
.
├── main.go
├── go.mod
└── README.md
```

---

## How It Works

### Root Route

The root route (`/`) is handled by `rootHandler`.

```go
mux.HandleFunc("/", rootHandler)
```

Example:

```bash
curl http://localhost:3000/
```

Response:

```text
This is the root page
```

---

### Home Route

The `/home` route is handled by `homeHandler`.

```go
mux.HandleFunc("/home", homeHandler)
```

Example:

```bash
curl http://localhost:3000/home
```

Response:

```text
Welcome
```

---

### Method Validation

Both handlers only allow `GET` requests.

Example:

```bash
curl -X POST http://localhost:3000/home
```

Response:

```text
Method Not Allowed
```

Status Code:

```text
405 Method Not Allowed
```

---

### 404 Handling

The root handler checks the requested path:

```go
if r.URL.Path != "/" {
	http.NotFound(w, r)
	return
}
```

This ensures that undefined routes return:

```text
404 Not Found
```

instead of incorrectly displaying the root page.

Example:

```bash
curl http://localhost:3000/about
```

Response:

```text
404 page not found
```

---

## Installation

### Clone the Repository

```bash
git clone https://github.com/your-username/go-http-server-status-codes.git
```

### Navigate into the Project

```bash
cd go-http-server-status-codes
```

### Run the Server

```bash
go run .
```

You should see:

```text
starting the server on port :3000
```

The server will be available at:

```text
http://localhost:3000
```

---

## Testing with curl

### Root Route

```bash
curl -i http://localhost:3000/
```

Expected:

```text
HTTP/1.1 200 OK

This is the root page
```

---

### Home Route

```bash
curl -i http://localhost:3000/home
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

