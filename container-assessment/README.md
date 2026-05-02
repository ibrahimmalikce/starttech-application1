# MuchToDo - Container Assessment

## Overview
This assessment containerizes the MuchToDo backend API (Golang + MongoDB) for StartupTech. The project covers two phases:
- **Phase 1**: Docker setup for local development using docker-compose
- **Phase 2**: Kubernetes deployment using Docker Desktop's built-in Kubernetes cluster

## Tech Stack
- **Backend**: Golang API (port 8080)
- **Database**: MongoDB 4.4
- **Containerization**: Docker + Docker Compose
- **Orchestration**: Kubernetes (Docker Desktop)

---

## Prerequisites
Make sure you have the following installed:
- [Docker Desktop](https://www.docker.com/products/docker-desktop/) (with Kubernetes enabled)
- kubectl (comes with Docker Desktop)
- Git

---

## Project Structurecontainer-assessment/
├── Dockerfile
├── docker-compose.yml
├── .dockerignore
├── kubernetes/
│   ├── namespace.yaml
│   ├── mongodb/
│   │   ├── mongodb-secret.yaml
│   │   ├── mongodb-configmap.yaml
│   │   ├── mongodb-pvc.yaml
│   │   ├── mongodb-deployment.yaml
│   │   └── mongodb-service.yaml
│   ├── backend/
│   │   ├── backend-secret.yaml
│   │   ├── backend-configmap.yaml
│   │   ├── backend-deployment.yaml
│   │   └── backend-service.yaml
│   └── ingress.yaml
├── scripts/
│   ├── docker-build.sh
│   ├── docker-run.sh
│   ├── k8s-deploy.sh
│   └── k8s-cleanup.sh
└── README.md

## Phase 1: Docker Setup

### Build the image
```bash
bash container-assessment/scripts/docker-build.sh
```

### Run with docker-compose
```bash
bash container-assessment/scripts/docker-run.sh
```

### Test
```bash
curl http://localhost:8080/health
```

## Phase 2: Kubernetes Deployment

### Deploy to Kubernetes
```bash
bash container-assessment/scripts/k8s-deploy.sh
```

### Access the application
```bash
kubectl port-forward service/backend-service 8081:8080 -n muchtodo
curl http://localhost:8081/health
```

### Check status
```bash
kubectl get all -n muchtodo
```

### Cleanup
```bash
bash container-assessment/scripts/k8s-cleanup.sh
```

## Environment Variables
| Variable | Description |
|----------|-------------|
| PORT | Application port (default: 8080) |
| MONGO_URI | MongoDB connection string |
| DB_NAME | Database name |
| JWT_SECRET_KEY | JWT signing key |
| JWT_EXPIRATION_HOURS | JWT token expiry |
| LOG_LEVEL | Logging level |
| LOG_FORMAT | Log format (json/text) |
