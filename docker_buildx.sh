#!/bin/bash

docker buildx build --file Dockerfile . --tag ghcr.io/lazzerbeamstudios/spectra:ute1756763469 --platform linux/amd64,linux/arm64 --push
