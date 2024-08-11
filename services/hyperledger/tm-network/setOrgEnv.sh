#!/bin/bash
#
# SPDX-License-Identifier: Apache-2.0




# default to using Main
ORG=${1:-Main}

# Exit on first error, print all commands.
set -e
set -o pipefail

# Where am I?
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." && pwd )"

ORDERER_CA=${DIR}/tm-network/organizations/ordererOrganizations/tmlabs.com/tlsca/tlsca.tmlabs.com-cert.pem
PEER0_Main_CA=${DIR}/tm-network/organizations/peerOrganizations/main.tmlabs.com/tlsca/tlsca.main.tmlabs.com-cert.pem
PEER0_ORG2_CA=${DIR}/tm-network/organizations/peerOrganizations/aux.tmlabs.com/tlsca/tlsca.aux.tmlabs.com-cert.pem
PEER0_ORG3_CA=${DIR}/tm-network/organizations/peerOrganizations/org3.tmlabs.com/tlsca/tlsca.org3.tmlabs.com-cert.pem


if [[ ${ORG,,} == "main" || ${ORG,,} == "digibank" ]]; then

   CORE_PEER_LOCALMSPID=MainMSP
   CORE_PEER_MSPCONFIGPATH=${DIR}/tm-network/organizations/peerOrganizations/main.tmlabs.com/users/Admin@main.tmlabs.com/msp
   CORE_PEER_ADDRESS=localhost:7051
   CORE_PEER_TLS_ROOTCERT_FILE=${DIR}/tm-network/organizations/peerOrganizations/main.tmlabs.com/tlsca/tlsca.main.tmlabs.com-cert.pem

elif [[ ${ORG,,} == "aux" || ${ORG,,} == "magnetocorp" ]]; then

   CORE_PEER_LOCALMSPID=AuxMSP
   CORE_PEER_MSPCONFIGPATH=${DIR}/tm-network/organizations/peerOrganizations/aux.tmlabs.com/users/Admin@aux.tmlabs.com/msp
   CORE_PEER_ADDRESS=localhost:9051
   CORE_PEER_TLS_ROOTCERT_FILE=${DIR}/tm-network/organizations/peerOrganizations/aux.tmlabs.com/tlsca/tlsca.aux.tmlabs.com-cert.pem

else
   echo "Unknown \"$ORG\", please choose Main/Digibank or Aux/Magnetocorp"
   echo "For example to get the environment variables to set upa Aux shell environment run:  ./setOrgEnv.sh Aux"
   echo
   echo "This can be automated to set them as well with:"
   echo
   echo 'export $(./setOrgEnv.sh Aux | xargs)'
   exit 1
fi

# output the variables that need to be set
echo "CORE_PEER_TLS_ENABLED=true"
echo "ORDERER_CA=${ORDERER_CA}"
echo "PEER0_Main_CA=${PEER0_Main_CA}"
echo "PEER0_ORG2_CA=${PEER0_ORG2_CA}"
echo "PEER0_ORG3_CA=${PEER0_ORG3_CA}"

echo "CORE_PEER_MSPCONFIGPATH=${CORE_PEER_MSPCONFIGPATH}"
echo "CORE_PEER_ADDRESS=${CORE_PEER_ADDRESS}"
echo "CORE_PEER_TLS_ROOTCERT_FILE=${CORE_PEER_TLS_ROOTCERT_FILE}"

echo "CORE_PEER_LOCALMSPID=${CORE_PEER_LOCALMSPID}"
