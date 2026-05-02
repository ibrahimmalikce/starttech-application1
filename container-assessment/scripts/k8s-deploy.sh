#!/bin/bash
echo "Deploying to Kubernetes..."
kubectl apply -f container-assessment/kubernetes/namespace.yaml
kubectl apply -f container-assessment/kubernetes/mongodb/
kubectl apply -f container-assessment/kubernetes/backend/
kubectl apply -f container-assessment/kubernetes/ingress.yaml
echo "Deployment complete!"
kubectl get all -n muchtodo
