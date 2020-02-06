#!/bin/bash

docker build -t me-go -f docker/Dockerfile .

if ! [ -z "$1" ] || [ "$1" == "--save" ]; then
    echo "Saving the image as '/me-go.tar'";
    mkdir -p files/image
    docker save -o files/image/me-go.tar me-go
fi
