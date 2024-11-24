#!/bin/bash

docker buildx build --file Dockerfile . --tag lazzerbeam/spectra:ute1732432112 --platform linux/amd64,linux/arm64 --push
