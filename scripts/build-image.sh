#!/bin/bash

docker build -t me-go -f docker/Dockerfile .

if ! [ -z "$1" ] || [ "$1" == "--save" ]; then
    echo "Saving the image as '/me-go.tar'";
    docker save -o me-go.tar me-go
fi
