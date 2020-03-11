package base

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SolarLabRU/fastpay-go-commons/enums/roles"
	. "github.com/SolarLabRU/fastpay-go-commons/errors"
	. "github.com/SolarLabRU/fastpay-go-commons/models"
	"github.com/SolarLabRU/fastpay-go-commons/requests"
	"github.com/SolarLabRU/fastpay-go-commons/responses"
	. "github.com/SolarLabRU/fastpay-go-commons/validation"
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

func InvokeChaincode(stub shim.ChaincodeStubInterface, chaincodeName string, nameFunc string, params interface{}) ([]byte, error) {
	var args [][]byte

	paramsAsBytes, err := json.Marshal(params)

	if err != nil {
		return nil, fmt.Errorf("Ошибка парсинга входных параметров. %s", err.Error())
	}

	args = append(args, []byte(nameFunc))
	args = append(args, paramsAsBytes)

	response := stub.InvokeChaincode(chaincodeName, args, "")

	if response.GetStatus() == 500 {
		return nil, fmt.Errorf("Ошибка при вызове чейнкода: %s", response.GetMessage())
	}

	return response.GetPayload(), nil

}

func GetBankByRemoteContract(stub shim.ChaincodeStubInterface, mspId string, address string) (*Bank, error) {
	request := requests.GetBank{
		Address: address,
		MSPId:   mspId,
	}

	// TODO Вынести название чейнкода в конастанту
	response, err := InvokeChaincode(stub, "banks", "getBankByMspIdAddress", request)
	if err != nil {
		return nil, err
	}

	var bankResponse responses.BankResponse
	err = json.Unmarshal(response, &bankResponse)

	if err != nil {
		return nil, fmt.Errorf("Ошибка вызова чейнкода banks. %s", err.Error())
	}

	fmt.Println("bank", bankResponse)

	return &bankResponse.Data, nil
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
		return CreateError(ErrorForbidden, "Права доступа к методу не определены")
	case roles.Regulator:
		if !bank.IsRegulator {
			return CreateError(ErrorForbidden, "Для досупа банк должен быть регулятором")
		}
	case roles.Owner:
		if !bank.IsRegulator {
			return CreateError(ErrorForbidden, "Для досупа банк должен быть владельцем")
		}
	}

	return nil
}

func CreateError(code uint, message string) error {
	baseError := BaseError{
		Code:    code,
		Message: message,
	}

	return createError(&baseError)
}

func CreateErrorWithData(code uint, message, data string) error {
	baseError := BaseError{
		Code:    code,
		Message: message,
		Data:    data,
	}

	return createError(&baseError)
}

func createError(baseError *BaseError) error {
	byteError, err := json.Marshal(baseError)
	if err != nil {
		return err // TODO Доделать вернуть JSON
	}

	return errors.New(string(byteError))
}

func CheckArgs(args string, request interface{}) error {
	err := json.Unmarshal([]byte(args), &request)

	if err != nil {
		return err
	}

	err = Validate.Struct(request)
	if err != nil {
		return err
	}

	return nil
}
