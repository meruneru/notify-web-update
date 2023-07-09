#!/usr/bin/bash

repository_url=$REPOSITORY_URL # Refer to AWS ECR dashbord

docker build -t notify-web-update .
docker tag notify-web-update:latest ${repository_url}:latest
docker push ${repository_url}:latest

