# Getting Started
This is a Go API server that also serves the frontend (embedded via `go:embed`) — no separate npm/Vite step needed.

Create a `.env` file in the project root with a port:
```
PORT=8080
```

Then run the server from the project root:
```
go run ./cmd/api
```

Visit `http://localhost:8080/app/` to see the frontend.