#!/bin/bash
echo "Building Docker image..."
docker build -f container-assessment/Dockerfile -t muchtodo-backend .
echo "Build complete!"
