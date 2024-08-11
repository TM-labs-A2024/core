#!/bin/bash

function createOrg1() {
  infoln "Enrolling the CA admin"
  mkdir -p organizations/peerOrganizations/org1.tmlabs.com/

  export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/org1.tmlabs.com/

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:7054 --caname ca-org1 --tls.certfiles "${PWD}/organizations/fabric-ca/org1/ca-cert.pem"
  { set +x; } 2>/dev/null

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-7054-ca-org1.pem
    OrganizationalUnitIdentifier: orderer' > "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/msp/config.yaml"

  # Since the CA serves as both the organization CA and TLS CA, copy the org's root cert that was generated by CA startup into the org level ca and tlsca directories

  # Copy org1's CA cert to org1's /msp/tlscacerts directory (for use in the channel MSP definition)
  mkdir -p "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/msp/tlscacerts"
  cp "${PWD}/organizations/fabric-ca/org1/ca-cert.pem" "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/msp/tlscacerts/ca.crt"

  # Copy org1's CA cert to org1's /tlsca directory (for use by clients)
  mkdir -p "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/tlsca"
  cp "${PWD}/organizations/fabric-ca/org1/ca-cert.pem" "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/tlsca/tlsca.org1.tmlabs.com-cert.pem"

  # Copy org1's CA cert to org1's /ca directory (for use by clients)
  mkdir -p "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/ca"
  cp "${PWD}/organizations/fabric-ca/org1/ca-cert.pem" "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/ca/ca.org1.tmlabs.com-cert.pem"

  infoln "Registering peer0"
  set -x
  fabric-ca-client register --caname ca-org1 --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles "${PWD}/organizations/fabric-ca/org1/ca-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering user"
  set -x
  fabric-ca-client register --caname ca-org1 --id.name user1 --id.secret user1pw --id.type client --tls.certfiles "${PWD}/organizations/fabric-ca/org1/ca-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering the org admin"
  set -x
  fabric-ca-client register --caname ca-org1 --id.name org1admin --id.secret org1adminpw --id.type admin --tls.certfiles "${PWD}/organizations/fabric-ca/org1/ca-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Generating the peer0 msp"
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/peers/peer0.org1.tmlabs.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org1/ca-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/peers/peer0.org1.tmlabs.com/msp/config.yaml"

  infoln "Generating the peer0-tls certificates, use --csr.hosts to specify Subject Alternative Names"
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/peers/peer0.org1.tmlabs.com/tls" --enrollment.profile tls --csr.hosts peer0.org1.tmlabs.com --csr.hosts localhost --tls.certfiles "${PWD}/organizations/fabric-ca/org1/ca-cert.pem"
  { set +x; } 2>/dev/null

  # Copy the tls CA cert, server cert, server keystore to well known file names in the peer's tls directory that are referenced by peer startup config
  cp "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/peers/peer0.org1.tmlabs.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/peers/peer0.org1.tmlabs.com/tls/ca.crt"
  cp "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/peers/peer0.org1.tmlabs.com/tls/signcerts/"* "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/peers/peer0.org1.tmlabs.com/tls/server.crt"
  cp "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/peers/peer0.org1.tmlabs.com/tls/keystore/"* "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/peers/peer0.org1.tmlabs.com/tls/server.key"

  infoln "Generating the user msp"
  set -x
  fabric-ca-client enroll -u https://user1:user1pw@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/users/User1@org1.tmlabs.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org1/ca-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/users/User1@org1.tmlabs.com/msp/config.yaml"

  infoln "Generating the org admin msp"
  set -x
  fabric-ca-client enroll -u https://org1admin:org1adminpw@localhost:7054 --caname ca-org1 -M "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/users/Admin@org1.tmlabs.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org1/ca-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org1.tmlabs.com/users/Admin@org1.tmlabs.com/msp/config.yaml"
}

function createOrg2() {
  infoln "Enrolling the CA admin"
  mkdir -p organizations/peerOrganizations/org2.tmlabs.com/

  export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/peerOrganizations/org2.tmlabs.com/

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:8054 --caname ca-org2 --tls.certfiles "${PWD}/organizations/fabric-ca/org2/ca-cert.pem"
  { set +x; } 2>/dev/null

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-org2.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-org2.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-org2.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-8054-ca-org2.pem
    OrganizationalUnitIdentifier: orderer' > "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/msp/config.yaml"

  # Since the CA serves as both the organization CA and TLS CA, copy the org's root cert that was generated by CA startup into the org level ca and tlsca directories

  # Copy org2's CA cert to org2's /msp/tlscacerts directory (for use in the channel MSP definition)
  mkdir -p "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/msp/tlscacerts"
  cp "${PWD}/organizations/fabric-ca/org2/ca-cert.pem" "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/msp/tlscacerts/ca.crt"

  # Copy org2's CA cert to org2's /tlsca directory (for use by clients)
  mkdir -p "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/tlsca"
  cp "${PWD}/organizations/fabric-ca/org2/ca-cert.pem" "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/tlsca/tlsca.org2.tmlabs.com-cert.pem"

  # Copy org2's CA cert to org2's /ca directory (for use by clients)
  mkdir -p "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/ca"
  cp "${PWD}/organizations/fabric-ca/org2/ca-cert.pem" "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/ca/ca.org2.tmlabs.com-cert.pem"

  infoln "Registering peer0"
  set -x
  fabric-ca-client register --caname ca-org2 --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles "${PWD}/organizations/fabric-ca/org2/ca-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering user"
  set -x
  fabric-ca-client register --caname ca-org2 --id.name user1 --id.secret user1pw --id.type client --tls.certfiles "${PWD}/organizations/fabric-ca/org2/ca-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering the org admin"
  set -x
  fabric-ca-client register --caname ca-org2 --id.name org2admin --id.secret org2adminpw --id.type admin --tls.certfiles "${PWD}/organizations/fabric-ca/org2/ca-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Generating the peer0 msp"
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:8054 --caname ca-org2 -M "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/peers/peer0.org2.tmlabs.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org2/ca-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/peers/peer0.org2.tmlabs.com/msp/config.yaml"

  infoln "Generating the peer0-tls certificates, use --csr.hosts to specify Subject Alternative Names"
  set -x
  fabric-ca-client enroll -u https://peer0:peer0pw@localhost:8054 --caname ca-org2 -M "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/peers/peer0.org2.tmlabs.com/tls" --enrollment.profile tls --csr.hosts peer0.org2.tmlabs.com --csr.hosts localhost --tls.certfiles "${PWD}/organizations/fabric-ca/org2/ca-cert.pem"
  { set +x; } 2>/dev/null

  # Copy the tls CA cert, server cert, server keystore to well known file names in the peer's tls directory that are referenced by peer startup config
  cp "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/peers/peer0.org2.tmlabs.com/tls/tlscacerts/"* "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/peers/peer0.org2.tmlabs.com/tls/ca.crt"
  cp "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/peers/peer0.org2.tmlabs.com/tls/signcerts/"* "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/peers/peer0.org2.tmlabs.com/tls/server.crt"
  cp "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/peers/peer0.org2.tmlabs.com/tls/keystore/"* "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/peers/peer0.org2.tmlabs.com/tls/server.key"

  infoln "Generating the user msp"
  set -x
  fabric-ca-client enroll -u https://user1:user1pw@localhost:8054 --caname ca-org2 -M "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/users/User1@org2.tmlabs.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org2/ca-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/users/User1@org2.tmlabs.com/msp/config.yaml"

  infoln "Generating the org admin msp"
  set -x
  fabric-ca-client enroll -u https://org2admin:org2adminpw@localhost:8054 --caname ca-org2 -M "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/users/Admin@org2.tmlabs.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/org2/ca-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/msp/config.yaml" "${PWD}/organizations/peerOrganizations/org2.tmlabs.com/users/Admin@org2.tmlabs.com/msp/config.yaml"
}

function createOrderer() {
  infoln "Enrolling the CA admin"
  mkdir -p organizations/ordererOrganizations/tmlabs.com

  export FABRIC_CA_CLIENT_HOME=${PWD}/organizations/ordererOrganizations/tmlabs.com

  set -x
  fabric-ca-client enroll -u https://admin:adminpw@localhost:9054 --caname ca-orderer --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg/ca-cert.pem"
  { set +x; } 2>/dev/null

  echo 'NodeOUs:
  Enable: true
  ClientOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: client
  PeerOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: peer
  AdminOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: admin
  OrdererOUIdentifier:
    Certificate: cacerts/localhost-9054-ca-orderer.pem
    OrganizationalUnitIdentifier: orderer' > "${PWD}/organizations/ordererOrganizations/tmlabs.com/msp/config.yaml"

  # Since the CA serves as both the organization CA and TLS CA, copy the org's root cert that was generated by CA startup into the org level ca and tlsca directories

  # Copy orderer org's CA cert to orderer org's /msp/tlscacerts directory (for use in the channel MSP definition)
  mkdir -p "${PWD}/organizations/ordererOrganizations/tmlabs.com/msp/tlscacerts"
  cp "${PWD}/organizations/fabric-ca/ordererOrg/ca-cert.pem" "${PWD}/organizations/ordererOrganizations/tmlabs.com/msp/tlscacerts/tlsca.tmlabs.com-cert.pem"

  # Copy orderer org's CA cert to orderer org's /tlsca directory (for use by clients)
  mkdir -p "${PWD}/organizations/ordererOrganizations/tmlabs.com/tlsca"
  cp "${PWD}/organizations/fabric-ca/ordererOrg/ca-cert.pem" "${PWD}/organizations/ordererOrganizations/tmlabs.com/tlsca/tlsca.tmlabs.com-cert.pem"

  infoln "Registering orderer"
  set -x
  fabric-ca-client register --caname ca-orderer --id.name orderer --id.secret ordererpw --id.type orderer --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg/ca-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Registering the orderer admin"
  set -x
  fabric-ca-client register --caname ca-orderer --id.name ordererAdmin --id.secret ordererAdminpw --id.type admin --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg/ca-cert.pem"
  { set +x; } 2>/dev/null

  infoln "Generating the orderer msp"
  set -x
  fabric-ca-client enroll -u https://orderer:ordererpw@localhost:9054 --caname ca-orderer -M "${PWD}/organizations/ordererOrganizations/tmlabs.com/orderers/orderer.tmlabs.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg/ca-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/ordererOrganizations/tmlabs.com/msp/config.yaml" "${PWD}/organizations/ordererOrganizations/tmlabs.com/orderers/orderer.tmlabs.com/msp/config.yaml"

  infoln "Generating the orderer-tls certificates, use --csr.hosts to specify Subject Alternative Names"
  set -x
  fabric-ca-client enroll -u https://orderer:ordererpw@localhost:9054 --caname ca-orderer -M "${PWD}/organizations/ordererOrganizations/tmlabs.com/orderers/orderer.tmlabs.com/tls" --enrollment.profile tls --csr.hosts orderer.tmlabs.com --csr.hosts localhost --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg/ca-cert.pem"
  { set +x; } 2>/dev/null

  # Copy the tls CA cert, server cert, server keystore to well known file names in the orderer's tls directory that are referenced by orderer startup config
  cp "${PWD}/organizations/ordererOrganizations/tmlabs.com/orderers/orderer.tmlabs.com/tls/tlscacerts/"* "${PWD}/organizations/ordererOrganizations/tmlabs.com/orderers/orderer.tmlabs.com/tls/ca.crt"
  cp "${PWD}/organizations/ordererOrganizations/tmlabs.com/orderers/orderer.tmlabs.com/tls/signcerts/"* "${PWD}/organizations/ordererOrganizations/tmlabs.com/orderers/orderer.tmlabs.com/tls/server.crt"
  cp "${PWD}/organizations/ordererOrganizations/tmlabs.com/orderers/orderer.tmlabs.com/tls/keystore/"* "${PWD}/organizations/ordererOrganizations/tmlabs.com/orderers/orderer.tmlabs.com/tls/server.key"

  # Copy orderer org's CA cert to orderer's /msp/tlscacerts directory (for use in the orderer MSP definition)
  mkdir -p "${PWD}/organizations/ordererOrganizations/tmlabs.com/orderers/orderer.tmlabs.com/msp/tlscacerts"
  cp "${PWD}/organizations/ordererOrganizations/tmlabs.com/orderers/orderer.tmlabs.com/tls/tlscacerts/"* "${PWD}/organizations/ordererOrganizations/tmlabs.com/orderers/orderer.tmlabs.com/msp/tlscacerts/tlsca.tmlabs.com-cert.pem"

  infoln "Generating the admin msp"
  set -x
  fabric-ca-client enroll -u https://ordererAdmin:ordererAdminpw@localhost:9054 --caname ca-orderer -M "${PWD}/organizations/ordererOrganizations/tmlabs.com/users/Admin@tmlabs.com/msp" --tls.certfiles "${PWD}/organizations/fabric-ca/ordererOrg/ca-cert.pem"
  { set +x; } 2>/dev/null

  cp "${PWD}/organizations/ordererOrganizations/tmlabs.com/msp/config.yaml" "${PWD}/organizations/ordererOrganizations/tmlabs.com/users/Admin@tmlabs.com/msp/config.yaml"
}
