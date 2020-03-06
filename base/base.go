package base

import (
	"encoding/json"
	"fmt"
	. "github.com/SolarLabRU/fastpay-go-commons/models"
	. "github.com/SolarLabRU/fastpay-go-commons/requests"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func GetSenderBank(ctx contractapi.TransactionContextInterface) (*Bank, error) {
	clientIdentity := ctx.GetClientIdentity()
	stub := ctx.GetStub()
	mspId, err := clientIdentity.GetMSPID()
	if err != nil {
		return nil, fmt.Errorf("Невозможно получить MSP ID. %s", err.Error())
	}

	address, isFound, err := clientIdentity.GetAttributeValue("address")
	if err != nil {
		return nil, fmt.Errorf("Невозможно получить атрибут из сертификата. %s", err.Error())
	}
	if isFound == false {
		return nil, fmt.Errorf("Отсутвует атрибут address в сертификате")
	}

	return GetBankByRemoteContract(stub, mspId, address)

}

func GetBankByRemoteContract(stub shim.ChaincodeStubInterface, mspId string, address string) (*Bank, error) {
	request := GetBankRequest{
		Address: address,
		MSPId:   mspId,
	}

	requestAsBytes, err := json.Marshal(request)

	if err != nil {
		return nil, fmt.Errorf("Ошибка парсинга входных параметров. %s", err.Error())
	}

	var args [][]byte

	args = append(args, []byte("getBank"))
	args = append(args, requestAsBytes)

	response := stub.InvokeChaincode("banks", args, "")

	var bank Bank
	err = json.Unmarshal(response.GetPayload(), &bank)

	if err != nil {
		return nil, fmt.Errorf("Ошибка вызова чейнкода banks. %s", err.Error())
	}

	return &bank, nil
}
