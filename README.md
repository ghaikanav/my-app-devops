# Go HTTP Server

A simple HTTP server built with Go that provides two endpoints:

## Endpoints

1. **`/` or `/hello`** - Returns a hello message
2. **`/env-echo`** - Echoes a message containing an environment variable (USER)

## Running the Server

### Prerequisites
- Go 1.21 or later installed on your system

### Steps
1. Navigate to the project directory:
   ```bash
   cd my-app-devops
   ```

2. Run the server:
   ```bash
   go run main.go
   ```

   The server will start on port 8080 by default.

3. To use a different port, set the `PORT` environment variable:
   ```bash
   PORT=3000 go run main.go
   ```

## Testing the Endpoints

Once the server is running, you can test the endpoints:

### Hello Endpoint
```bash
curl http://localhost:8080/hello
# or
curl http://localhost:8080/
```

### Environment Variable Echo Endpoint
```bash
curl http://localhost:8080/env-echo
```

## Environment Variables

- `PORT` - Server port (default: 8080)
- `USER` - User environment variable (used in the env-echo endpoint)

## Building

To build an executable:
```bash
go build -o server main.go
./server
```

## Docker

### Building and Running with Docker

1. **Build the Docker image:**
   ```bash
   docker build -t go-http-server .
   ```

2. **Run the container:**
   ```bash
   docker run -p 8080:8080 go-http-server
   ```

3. **Run with custom environment variables:**
   ```bash
   docker run -p 8080:8080 -e USER=custom-user go-http-server
   ```

### Using Docker Compose

1. **Build and run with Docker Compose:**
   ```bash
   docker-compose up --build
   ```

2. **Run in background:**
   ```bash
   docker-compose up -d
   ```

3. **Stop the service:**
   ```bash
   docker-compose down
   ```

4. **View logs:**
   ```bash
   docker-compose logs -f
   ```
