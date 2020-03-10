package base

import (
	"encoding/json"
	"fmt"
	"github.com/SolarLabRU/fastpay-go-commons/enums/roles"
	. "github.com/SolarLabRU/fastpay-go-commons/models"
	"github.com/SolarLabRU/fastpay-go-commons/requests"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func GetSenderBank(ctx contractapi.TransactionContextInterface) (*Bank, error) {
	clientIdentity := ctx.GetClientIdentity()
	stub := ctx.GetStub()
	mspId, err := clientIdentity.GetMSPID()
	fmt.Println("mspId", mspId) // TODO Убрать
	if err != nil {
		return nil, fmt.Errorf("Невозможно получить MSP ID. %s", err.Error())
	}

	address, _ := GetSenderAddressFromCertificate(clientIdentity)

	if address == "" {
		return nil, fmt.Errorf("Отсутвует атрибут address в сертификате")
	}

	return GetBankByRemoteContract(stub, mspId, address)
}

func GetSenderAddressFromCertificate(identity cid.ClientIdentity) (string, error) {
	address, isFound, err := identity.GetAttributeValue("address")
	fmt.Println("address", address, isFound, err) // TODO Убрать

	address, isFound, err = func() (string, bool, error) { return "263093b1c21f98c5f9b6433bf9bbb97bb87b6e79", true, nil }() // TODO Убрать

	if err != nil {
		return "", fmt.Errorf("Невозможно получить атрибут из сертификата. %s", err.Error())
	}
	if isFound == false {
		return "", fmt.Errorf("Отсутвует атрибут address в сертификате")
	}

	return address, nil
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

	args = append(args, []byte("getBankByMspIdAddress"))
	args = append(args, requestAsBytes)

	response := stub.InvokeChaincode("banks", args, "")

	fmt.Println("response GetBankByRemoteContract", response)

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

	return nil
}
