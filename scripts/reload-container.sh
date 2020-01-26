#!/bin/bash

CONTAINER_NAME=me-go
PORT=8082
ARGUMENT_RESTART_CONTAINER=

if [ -z "$1" ] || [ "$1" != "--production" ]; then
    echo "Container will start in dev mode. Pass '--production' to run in production mode.";
else
    echo "Container will start in production mode.";
    ARGUMENT_RESTART_CONTAINER="--restart unless-stopped"
fi

echo "Removing existing containers.."
if [ "$(docker ps -a | grep -w $CONTAINER_NAME)" ]; then
    docker stop $CONTAINER_NAME && docker rm $CONTAINER_NAME && echo "Container removed"
else
    echo "Nothing to remove"
fi

echo "Starting container: '$CONTAINER_NAME'"
docker run --name $CONTAINER_NAME \
    $ARGUMENT_RESTART_CONTAINER \
    -d \
    -p $PORT:8080 \
    -v ~/.cdocker/vlm/me-go/pages:/home/go/pages \
    me-go \
    ./ssr
