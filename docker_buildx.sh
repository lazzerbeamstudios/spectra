#!/bin/bash

docker buildx build --file Dockerfile . --tag lazzerbeam/spectra:ute1732552258 --platform linux/amd64,linux/arm64 --push
