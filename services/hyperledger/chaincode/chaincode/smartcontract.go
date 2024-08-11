package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type HealthRecord struct {
	ID      string
	Content string
	Address string
}

func (s *SmartContract) CreateHealthRecord(ctx contractapi.TransactionContextInterface, id, content, address string) error {
	exists, err := s.HealthRecordExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the healthRecord %s already exists", id)
	}

	healthRecord := HealthRecord{
		ID:      id,
		Address: address,
		Content: content,
	}
	healthRecordJSON, err := json.Marshal(healthRecord)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, healthRecordJSON)
}

// ReadHealthRecord returns the healthRecord stored in the world state with given id.
func (s *SmartContract) ReadHealthRecord(ctx contractapi.TransactionContextInterface, id string) (*HealthRecord, error) {
	healthRecordJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if healthRecordJSON == nil {
		return nil, fmt.Errorf("the healthRecord %s does not exist", id)
	}

	var healthRecord HealthRecord
	err = json.Unmarshal(healthRecordJSON, &healthRecord)
	if err != nil {
		return nil, err
	}

	return &healthRecord, nil
}

// HealthRecordExists returns true when healthRecord with given ID exists in world state
func (s *SmartContract) HealthRecordExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	healthRecordJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return healthRecordJSON != nil, nil
}

// GetAllHealthRecords returns all healthRecords found in world state
func (s *SmartContract) GetAllHealthRecords(ctx contractapi.TransactionContextInterface) ([]*HealthRecord, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all healthRecords in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var healthRecords []*HealthRecord
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var healthRecord HealthRecord
		err = json.Unmarshal(queryResponse.Value, &healthRecord)
		if err != nil {
			return nil, err
		}
		healthRecords = append(healthRecords, &healthRecord)
	}

	return healthRecords, nil
}

// GetAllHealthRecords returns all healthRecords found in world state
func (s *SmartContract) GetAllHealthRecordsByAddress(ctx contractapi.TransactionContextInterface, address string) ([]*HealthRecord, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all healthRecords in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var healthRecords []*HealthRecord
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var healthRecord HealthRecord
		err = json.Unmarshal(queryResponse.Value, &healthRecord)
		if err != nil {
			return nil, err
		}

		if healthRecord.Address == address {
			healthRecords = append(healthRecords, &healthRecord)
		}
	}

	return healthRecords, nil
}
