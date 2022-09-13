#!/bin/sh
# Create the EKS cluster
eksctl create cluster --config-file go-web.yaml
# The OIDC Connect
eksctl utils associate-iam-oidc-provider --cluster go-web --approve

# The ServiceAccount and Role given the policy
eksctl create iamserviceaccount \
    --name go-web-service-account \
    --namespace default \
    --cluster go-web \
    --role-name "go-web-dynamodb-access" \
    --attach-policy-arn arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess \
    --approve --override-existing-serviceaccounts

# Load the cluster with the Kubernetes deployment and service
eksctl get clusters --region ap-southeast-2
kubectl apply -f ../kubernetes/mainfest.yml
kubectl get svc
