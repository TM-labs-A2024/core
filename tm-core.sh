#!/bin/bash

DOCKER_COMPOSE='docker compose'
# Get docker sock path from environment variable
SOCK="${DOCKER_HOST:-/var/run/docker.sock}"
DOCKER_SOCK="${SOCK##unix://}"
CMD=$1

echo "$(pwd)"

compose-up() {
    DOCKER_SOCK="${DOCKER_SOCK}" ${DOCKER_COMPOSE}  -f ./docker-compose.yaml up
}

compose-build() {
    DOCKER_SOCK="${DOCKER_SOCK}" ${DOCKER_COMPOSE}  -f ./docker-compose.yaml build
}

case $CMD in
    "up")
        compose-up
        ;;
    "build")
        compose-build
        ;;
    "up-build")
        compose-build
        compose-up
        ;;
    *)
        echo "Unknown command."
        ;;
esac
