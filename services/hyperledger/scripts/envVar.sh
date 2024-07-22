#!/bin/bash

. ./scripts/utils.sh

export CORE_PEER_TLS_ENABLED=true
export ORDERER_CA=./orgs/ordererOrganizations/tmlabs.com/tlsca/tlsca.tmlabs.com-cert.pem
export PEER0_ORG1_CA=./orgs/peerOrganizations/main.tmlabs.com/tlsca/tlsca.main.tmlabs.com-cert.pem
export PEER0_ORG2_CA=./orgs/peerOrganizations/aux.tmlabs.com/tlsca/tlsca.aux.tmlabs.com-cert.pem

# Set environment variables for the peer org
setGlobals() {
  local USING_ORG=""
  if [ -z "$OVERRIDE_ORG" ]; then
    USING_ORG=$1
  else
    USING_ORG="${OVERRIDE_ORG}"
  fi
  infoln "Using organization ${USING_ORG}"
  if [ "$USING_ORG" = "main" ]; then
    export CORE_PEER_LOCALMSPID=mainMSP
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_main_CA
    export CORE_PEER_MSPCONFIGPATH=./orgs/peerOrganizations/main.tmlabs.com/users/Admin@main.tmlabs.com/msp
    export CORE_PEER_ADDRESS=localhost:7051
  elif [ "$USING_ORG" = "aux" ]; then
    export CORE_PEER_LOCALMSPID=auxMSP
    export CORE_PEER_TLS_ROOTCERT_FILE=$PEER0_aux_CA
    export CORE_PEER_MSPCONFIGPATH=./orgs/peerOrganizations/aux.tmlabs.com/users/Admin@aux.tmlabs.com/msp
    export CORE_PEER_ADDRESS=localhost:9051
  fi

  if [ "$VERBOSE" = "true" ]; then
    env | grep CORE
  fi
}

# parsePeerConnectionParameters $@
# Helper function that sets the peer connection parameters for a chaincode
# operation
parsePeerConnectionParameters() {
  PEER_CONN_PARMS=()
  PEERS=""
  while [ "$#" -gt 0 ]; do
    setGlobals $1
    PEER="peer0.$1"
    ## Set peer addresses
    if [ -z "$PEERS" ]
    then
	PEERS="$PEER"
    else
	PEERS="$PEERS $PEER"
    fi
    PEER_CONN_PARMS=("${PEER_CONN_PARMS[@]}" --peerAddresses $CORE_PEER_ADDRESS)
    ## Set path to TLS certificate
    CA=PEER0_ORG$1_CA
    TLSINFO=(--tlsRootCertFiles "${!CA}")
    PEER_CONN_PARMS=("${PEER_CONN_PARMS[@]}" "${TLSINFO[@]}")
    # shift by one to get to the next organization
    shift
  done
}

verifyResult() {
  if [ $1 -ne 0 ]; then
    fatalln "$2"
  fi
}
