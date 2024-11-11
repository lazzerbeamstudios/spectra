#!/bin/bash

docker buildx build --file Dockerfile . --tag lazzerbeam/spectra:ute1731364188 --platform linux/amd64,linux/arm64 --push
