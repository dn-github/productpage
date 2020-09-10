#!/bin/sh

set -e
export GO111MODULE=on
GOOS=linux go build -o productpage
docker build -t devnryn/productpage:v1 .
docker push devnryn/productpage:v1
rm productpage
