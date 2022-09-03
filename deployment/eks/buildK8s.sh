#!/bin/sh
eksctl create cluster --config-file go-web.yaml
eksctl get clusters --region ap-southeast-2
kubectl apply -f ../deployment/kubernetes/mainfest.yml
kubectl get svc
