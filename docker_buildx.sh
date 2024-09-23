#!/bin/bash

docker buildx build --file Dockerfile . --tag lazzerbeam/spectra:ute1726809903 --platform linux/amd64,linux/arm64 --push
