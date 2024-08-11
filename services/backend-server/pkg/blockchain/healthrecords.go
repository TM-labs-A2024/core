package blockchain

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	mspID       = "Org1MSP"
	certPath    = "/users/User1@org1.tmlabs.com/msp/signcerts"
	keyPath     = "/users/User1@org1.tmlabs.com/msp/keystore"
	tlsCertPath = "/peers/peer0.org1.tmlabs.com/tls/ca.crt"
	gatewayPeer = "peer0.org1.tmlabs.com"
)

type Client struct {
	contract     *client.Contract
	mspID        string
	cryptoPath   string
	certPath     string
	keyPath      string
	tlsCertPath  string
	peerEndpoint string
	gatewayPeer  string
}

func New(chaincodeName, channelName, cryptoPath, peerEndpoint string) (*Client, error) {
	c := &Client{
		mspID:        mspID,
		cryptoPath:   cryptoPath,
		certPath:     filepath.Join(cryptoPath, certPath),
		keyPath:      filepath.Join(cryptoPath, keyPath),
		tlsCertPath:  filepath.Join(cryptoPath, tlsCertPath),
		peerEndpoint: peerEndpoint,
		gatewayPeer:  gatewayPeer,
	}

	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection, err := c.newGrpcConnection()
	if err != nil {
		return nil, err
	}

	id, err := c.newIdentity()
	if err != nil {
		return nil, err
	}

	sign, err := c.newSign()
	if err != nil {
		return nil, err
	}

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		return nil, err
	}

	network := gw.GetNetwork(channelName)
	c.contract = network.GetContract(chaincodeName)

	return c, nil
}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func (c *Client) newGrpcConnection() (*grpc.ClientConn, error) {
	certificatePEM, err := os.ReadFile(c.tlsCertPath)
	if err != nil {
		return nil, err
	}

	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, c.gatewayPeer)

	connection, err := grpc.NewClient(c.peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		return nil, err
	}

	return connection, nil
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func (c *Client) newIdentity() (*identity.X509Identity, error) {
	certificatePEM, err := readFirstFile(c.certPath)
	if err != nil {
		return nil, err
	}

	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		return nil, err
	}

	id, err := identity.NewX509Identity(c.mspID, certificate)
	if err != nil {
		return nil, err
	}

	return id, nil
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func (c *Client) newSign() (identity.Sign, error) {
	privateKeyPEM, err := readFirstFile(c.keyPath)
	if err != nil {
		return nil, err
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		return nil, err
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		return nil, err
	}

	return sign, nil
}

func readFirstFile(dirPath string) ([]byte, error) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}

	fileNames, err := dir.Readdirnames(1)
	if err != nil {
		return nil, err
	}

	return os.ReadFile(path.Join(dirPath, fileNames[0]))
}

func (c *Client) GetAllHealthRecords() (string, error) {
	evaluateResult, err := c.contract.EvaluateTransaction("GetAllHealthRecords")
	if err != nil {
		return "", err
	}
	result, err := formatJSON(evaluateResult)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) GetAllHealthRecordsByAddress(address string) (string, error) {
	evaluateResult, err := c.contract.EvaluateTransaction("GetAllHealthRecordsByAddress", address)
	if err != nil {
		return "", err
	}

	result, err := formatJSON(evaluateResult)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) CreateHealthRecord(address, content string) error {
	_, err := c.contract.SubmitTransaction("CreateHealthRecord", uuid.NewString(), content, address)
	if err != nil {
		return err
	}

	return nil
}

// Format JSON data
func formatJSON(data []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
