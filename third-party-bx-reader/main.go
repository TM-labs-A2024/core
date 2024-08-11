package main

import (
	"log"

	"github.com/TM-labs-A2024/core/services/backend-server/pkg/blockchain"
	"github.com/TM-labs-A2024/core/services/backend-server/pkg/config"
)

func main() {
	config, err := config.LoadConfig("/home/mariajdab/Jorge/core/config.yaml")
	if err != nil {
		panic(err)
	}

	client, err := blockchain.New(
		config.ChaincodeName,
		config.ChannelName,
		"/home/mariajdab/Jorge/core/services/hyperledger/tm-network/organizations/peerOrganizations/org1.tmlabs.com",
		"localhost:7051",
	)
	if err != nil {
		panic(err)
	}

	result, err := client.GetAllHealthRecords()
	if err != nil {
		panic(err)
	}

	log.Println(result)
}
