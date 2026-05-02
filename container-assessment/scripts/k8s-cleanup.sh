#!/bin/bash
echo "Cleaning up Kubernetes resources..."
kubectl delete namespace muchtodo
echo "Cleanup complete!"
