#!/bin/bash

docker buildx build --file Dockerfile . --tag lazzerbeam/api-v1:ute1726809903 --platform linux/amd64,linux/arm64 --push
