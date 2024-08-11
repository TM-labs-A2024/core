#!/bin/bash

set -e

DOCKER_COMPOSE='docker compose'
# Get docker sock path from environment variable
CMD=$1
ROOT=${PWD}
BLOCKCHAIN="$ROOT/services/hyperledger/tm-network"

echo "$(pwd)"

compose-up() {
    ${DOCKER_COMPOSE}  -f ./docker-compose.yaml up
}

compose-build() {
    ${DOCKER_COMPOSE}  -f ./docker-compose.yaml build
}

start-blockchain() {
    cd $BLOCKCHAIN
    ./network.sh up createChannel -c tm-healthcore -ca
    ./network.sh deployCC -ccn health-record -ccp ../chaincode/ -ccl go
}

clean-blockchain() {
    cd $BLOCKCHAIN
    ./network.sh down
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
    "clean")
        clean-blockchain
        ;;
    "chain-up")
        start-blockchain
        ;;
    *)
        echo "Unknown command."
        ;;
esac
