package base

import (
	"encoding/json"
	"fmt"
	"github.com/SolarLabRU/fastpay-go-commons/enums/roles"
	. "github.com/SolarLabRU/fastpay-go-commons/models"
	"github.com/SolarLabRU/fastpay-go-commons/requests"
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
	request := requests.GetBank{
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

func CheckAccess(ctx contractapi.TransactionContextInterface, role roles.Roles) error {
	return CheckAccessWithBank(ctx, nil, role)
}

func CheckAccessWithBank(ctx contractapi.TransactionContextInterface, bank *Bank, role roles.Roles) error {
	if bank == nil {
		var err error = nil
		bank, err = GetSenderBank(ctx)
		if err != nil {
			return err
		}
	}

	switch role {
	case roles.Undefined:
		return fmt.Errorf("Права доступа к методу не определены", "60601")
	case roles.Regulator:
		if !bank.IsRegulator {
			return fmt.Errorf("Для досупа банк должен быть регулятором", "60601")
		}
	case roles.Owner:
		if !bank.IsRegulator {
			return fmt.Errorf("Для досупа банк должен быть владельцем", "60601")
		}
	}
}
