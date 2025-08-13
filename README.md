# Go HTTP Server

A simple HTTP server built with Go that provides three endpoints with environment variable integration:

## Application Features

- **Environment Variable Integration**: Reads from `PORT`, `USER`, and `URL` environment variables
- **Configurable Port**: Server port can be set via `PORT` environment variable (defaults to 8080)
- **Health Check Ready**: Endpoints designed for Kubernetes health checks
- **Secret Management**: Demonstrates secure environment variable handling

## Endpoints

1. **`/` or `/hello`** - Returns a hello message
2. **`/env-echo`** - Echoes a message containing an environment variable (USER)
3. **`/secret-echo`** - Echoes a message containing a secret environment variable (URL)

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

### Test with Different Ports
```bash
# Test on default port 8080
curl http://localhost:8080/hello

# Test on custom port 3000
PORT=3000 go run main.go &
curl http://localhost:3000/hello
kill %1
```

### Test Environment Variable Integration
```bash
# Test USER environment variable
USER=John curl http://localhost:8080/env-echo

# Test URL environment variable  
URL=secret-data curl http://localhost:8080/secret-echo
```

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

### Secret Echo Endpoint
```bash
curl http://localhost:8080/secret-echo
```

## Environment Variables

- `PORT` - Server port (default: 8080)
- `USER` - User environment variable (used in the env-echo endpoint)
- `URL` - Secret environment variable (used in the secret-echo endpoint)

## Go Code Structure

The application is structured with:
- **Handler Functions**: Separate functions for each endpoint (`helloHandler`, `envEchoHandler`, `secretEchoHandler`)
- **Environment Variable Handling**: Graceful fallbacks when environment variables are not set
- **HTTP Server Setup**: Configurable port with environment variable support
- **Error Handling**: Proper HTTP response writing and error logging

## Building

To build an executable:
```bash
go build -o server main.go
./server
```

## Kubernetes Architecture

Your application is deployed using the following Kubernetes components:

### Components Overview
- **ConfigMap** (`config-map.yaml`): Stores application configuration (PORT, USER)
- **Secret** (`secret-db.yaml`): Stores sensitive database connection string
- **Deployment** (`deployment.yaml`): Manages application pods with 2 replicas
- **Service** (`service.yaml`): Exposes the application via NodePort

### Resource Configuration
- **Replicas**: 2 (high availability)
- **Resources**: 
  - Requests: 100m CPU, 128Mi Memory
  - Limits: 500m CPU, 512Mi Memory
- **Port**: 8080 (container and service)
- **Image**: `kanavghai/my-app:latest`

## Kubernetes Deployment

### Prerequisites
- Kubernetes cluster (local: minikube, Docker Desktop, or cloud: GKE, EKS, AKS)
- `kubectl` CLI tool installed and configured
- Docker image built and pushed to a registry

### 1. Build and Push Docker Image

First, build and push your Docker image to a registry:

```bash
# Build the image
docker build -t kanavghai/my-app:latest .

# Push to registry (replace with your registry)
docker push kanavghai/my-app:latest
```

### 2. Deploy to Kubernetes

#### Option A: Deploy All Components at Once
```bash
# Apply all Kubernetes manifests
kubectl apply -f config-map.yaml
kubectl apply -f secret-db.yaml
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

#### Option B: Deploy Step by Step
```bash
# 1. Create ConfigMap for application configuration
kubectl apply -f config-map.yaml

# 2. Create Secret for database connection
kubectl apply -f secret-db.yaml

# 3. Deploy the application
kubectl apply -f deployment.yaml

# 4. Create service for external access
kubectl apply -f service.yaml
```

### 3. Verify Deployment

```bash
# Check deployment status
kubectl get deployments

# Check pods
kubectl get pods

# Check services
kubectl get services

# Check configmaps and secrets
kubectl get configmaps
kubectl get secrets
```

### 4. Access the Application

#### Get Service Details
```bash
# Get service information
kubectl get service my-app-service

# Get detailed service info
kubectl describe service my-app-service
```

#### Access via Port Forward (for testing)
```bash
# Forward local port to service
kubectl port-forward service/my-app-service 8080:8080

# Then access: http://localhost:8080
```

#### Access via NodePort (if using minikube)
```bash
# Get minikube IP
minikube ip

# Get NodePort
kubectl get service my-app-service -o jsonpath='{.spec.ports[0].nodePort}'

# Access: http://<minikube-ip>:<nodePort>
```

### 5. Monitor and Debug

```bash
# View pod logs
kubectl logs -l app=my-app

# Follow logs in real-time
kubectl logs -f -l app=my-app

# Describe deployment
kubectl describe deployment my-app-deployment

# Describe pods
kubectl describe pods -l app=my-app
```

### 6. Scale the Application

```bash
# Scale to 5 replicas
kubectl scale deployment my-app-deployment --replicas=5

# Check scaling status
kubectl get deployment my-app-deployment
```

### 7. Update the Application

```bash
# Update image (trigger rolling update)
kubectl set image deployment/my-app-deployment my-app=kanavghai/my-app:v2

# Or update deployment YAML and reapply
kubectl apply -f deployment.yaml
```

### 8. Clean Up

```bash
# Delete all resources
kubectl delete -f deployment.yaml
kubectl delete -f service.yaml
kubectl delete -f config-map.yaml
kubectl delete -f secret-db.yaml

# Or delete all at once
kubectl delete -f .
```

## Troubleshooting

### Common Issues

#### 1. Image Pull Errors
```bash
# Check if image exists in registry
docker pull kanavghai/my-app:latest

# Verify image name in deployment.yaml matches your registry
```

#### 2. Pod Startup Issues
```bash
# Check pod status
kubectl get pods -l app=my-app

# View pod events
kubectl describe pod <pod-name>

# Check pod logs
kubectl logs <pod-name>
```

#### 3. Service Access Issues
```bash
# Verify service is running
kubectl get service my-app-service

# Check endpoints
kubectl get endpoints my-app-service

# Test service connectivity
kubectl run test-pod --image=busybox --rm -it --restart=Never -- wget -O- http://my-app-service:8080/hello
```

#### 4. Configuration Issues
```bash
# Verify ConfigMap
kubectl get configmap my-app-config -o yaml

# Verify Secret
kubectl get secret db-secret -o yaml

# Check if pods are using the config
kubectl describe pod <pod-name> | grep -A 10 "Environment:"
```

### Health Checks
Your deployment includes readiness and liveness probes. Check them with:
```bash
# View deployment details
kubectl describe deployment my-app-deployment

# Check probe status in pod description
kubectl describe pod <pod-name>
```

### Performance Monitoring
```bash
# Monitor resource usage
kubectl top pods -l app=my-app

# Monitor resource usage by node
kubectl top nodes
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
