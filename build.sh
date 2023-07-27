#!/bin/bash

docker build -t volthai7us/url-shortener:latest .

docker push volthai7us/url-shortener:latest

kubectl apply -f deployment.yaml
kubectl apply -f service.yaml