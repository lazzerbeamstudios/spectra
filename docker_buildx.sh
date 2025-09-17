#!/bin/bash

docker buildx build --file Dockerfile . --tag ghcr.io/lazzerbeam-studios/spectra:ute1758120276 --platform linux/amd64,linux/arm64 --push
