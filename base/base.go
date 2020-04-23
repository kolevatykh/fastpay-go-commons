package base

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math"
	"regexp"
	"strings"

	"github.com/SolarLabRU/fastpay-go-commons/cc-errors"
	"github.com/SolarLabRU/fastpay-go-commons/enums/access-role-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"
	"github.com/SolarLabRU/fastpay-go-commons/models"
	"github.com/SolarLabRU/fastpay-go-commons/requests"
	"github.com/SolarLabRU/fastpay-go-commons/responses"
	"github.com/SolarLabRU/fastpay-go-commons/validation"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const (
	ChaincodeBankName = "banks"
)

func GetSenderBank(ctx contractapi.TransactionContextInterface) (*models.Bank, error) {
	clientIdentity := ctx.GetClientIdentity()
	stub := ctx.GetStub()
	mspId, err := clientIdentity.GetMSPID()

	if err != nil {
		return nil, CreateError(cc_errors.ErrorGetMspId, fmt.Sprintf("Невозможно получить MSP ID. %s", err.Error()))
	}

	address, err := GetSenderAddressFromCertificate(clientIdentity)

	if err != nil {
		return nil, err
	}

	return GetBankByRemoteContract(stub, mspId, address)
}

func GetSenderAddressFromCertificate(identity cid.ClientIdentity) (string, error) {
	address, isFound, _ := identity.GetAttributeValue("address")

	//address, isFound, _ = func() (string, bool, error) { return "263093b1c21f98c5f9b6433bf9bbb97bb87b6e79", true, nil }() // TODO Убрать

	if !isFound {
		return "", CreateError(cc_errors.ErrorCertificateNotValid, "Отсутвует атрибут address в сертификате")
	}

	return address, nil
}

func GetBankIdFromCertificate(identity cid.ClientIdentity) (string, error) {
	address, isFound, _ := identity.GetAttributeValue("bankId")

	//address, isFound, _ = func() (string, bool, error) { return "clientBank1", true, nil }() // TODO Убрать

	if !isFound {
		return "", CreateError(cc_errors.ErrorCertificateNotValid, "Отсутвует атрибут bankId в сертификате")
	}

	return address, nil
}

func InvokeChaincode(stub shim.ChaincodeStubInterface, chaincodeName string, nameFunc string, params interface{}) ([]byte, error) {
	var args [][]byte

	paramsAsBytes, err := json.Marshal(params)

	if err != nil {
		return nil, CreateError(cc_errors.ErrorDefault, fmt.Sprintf("Ошибка парсинга входных параметров. %s", err.Error()))
	}

	args = append(args, []byte(nameFunc))
	args = append(args, paramsAsBytes)

	response := stub.InvokeChaincode(chaincodeName, args, "")

	if response.GetStatus() == 500 {
		fmt.Println("Ошибка при вызове чейнкода: ", response.GetMessage())

		return nil, parseErrorFromAnotherChaincode(response.GetMessage())
	}

	return response.GetPayload(), nil
}

func GetEventName(chaincodeName, functionName string) string {
	return fmt.Sprintf("%s_%s", chaincodeName, functionName)
}

func InvokeChaincodeWithEmptyParams(stub shim.ChaincodeStubInterface, chaincodeName string, nameFunc string) ([]byte, error) {
	var args [][]byte

	args = append(args, []byte(nameFunc))
	args = append(args, []byte(""))

	response := stub.InvokeChaincode(chaincodeName, args, "")

	if response.GetStatus() == 500 {
		fmt.Println("Ошибка при вызове чейнкода: ", response.GetMessage())

		return nil, parseErrorFromAnotherChaincode(response.GetMessage())
	}

	return response.GetPayload(), nil
}

func parseErrorFromAnotherChaincode(message string) error {
	var baseError cc_errors.BaseError

	err := json.Unmarshal([]byte(message), &baseError)

	if err != nil {
		return CreateError(cc_errors.ErrorDefault, fmt.Sprintf("Ошибка при вызове чейнкода: %s", message))
	}

	return CreateError(baseError.Code, fmt.Sprintf("Ошибка при вызове чейнкода: %s", baseError.Message))

}

func GetBankByRemoteContract(stub shim.ChaincodeStubInterface, mspId string, address string) (*models.Bank, error) {
	request := requests.GetBankRequest{
		Address: address,
		MSPId:   mspId,
	}

	response, err := InvokeChaincode(stub, ChaincodeBankName, "getBankByMspIdAddress", request)
	if err != nil {
		return nil, err
	}

	var bankResponse responses.BankResponse
	err = json.Unmarshal(response, &bankResponse)

	if err != nil {
		return nil, CreateError(cc_errors.ErrorJsonUnmarshal, fmt.Sprintf("Ошибка десерилизации ответа после вызова чейнкода banks. %s", err.Error()))
	}

	return &bankResponse.Data, nil
}

func SenderBankIsAvailable(ctx contractapi.TransactionContextInterface) error {
	bank, _ := GetSenderBank(ctx)
	return SenderBankIsAvailableWithBank(ctx, bank)
}

func SenderBankIsAvailableWithBank(ctx contractapi.TransactionContextInterface, bank *models.Bank) error {
	if bank == nil {
		var err error = nil
		bank, err = GetSenderBank(ctx)
		if err != nil {
			return err
		}
	}

	if bank == nil || bank.State != state_enum.Available {
		return CreateError(cc_errors.ErrorBankNotAvailable, "Банк отправителя не доступен")
	}

	return nil
}

func CheckAccess(ctx contractapi.TransactionContextInterface, role access_role_enum.AccessRole, checkAvailable bool) error {
	return CheckAccessWithBank(ctx, nil, role, checkAvailable)
}

func CheckAccessWithBank(ctx contractapi.TransactionContextInterface, bank *models.Bank, role access_role_enum.AccessRole, checkAvailable bool) error {
	if role == access_role_enum.Any {
		return nil
	}

	if bank == nil {
		var err error = nil
		bank, err = GetSenderBank(ctx)
		if err != nil {
			return err
		}
	}

	if checkAvailable {
		err := SenderBankIsAvailableWithBank(ctx, bank)
		if err != nil {
			return err
		}
	}

	switch role {
	case access_role_enum.Undefined:
		return CreateError(cc_errors.ErrorForbidden, "Права доступа к методу не определены")
	case access_role_enum.Regulator:
		if !bank.IsRegulator {
			return CreateError(cc_errors.ErrorForbidden, "Для досупа банк должен быть регулятором")
		}
	case access_role_enum.Owner:
		if !bank.IsOwner {
			return CreateError(cc_errors.ErrorForbidden, "Для досупа банк должен быть владельцем")
		}
	}

	return nil
}

func CheckCalledChaincode(stub shim.ChaincodeStubInterface, name, function string) (bool, error) {
	signedProposal, err := stub.GetSignedProposal()
	if err != nil {
		return false, err
	}
	stringSignedProposal := string(signedProposal.GetProposalBytes())
	filterString := regexp.MustCompile(`[^a-zA-Z 0-9\n_-]`).ReplaceAllString(stringSignedProposal, "")
	filterString2 := regexp.MustCompile(`[\t\r\n]+`).ReplaceAllString(strings.TrimSpace(filterString), "\n")
	splitResult := strings.Split(filterString2, "\n")

	nameChaincode := "not set"
	nameFunc := "not set"

	if len(stringSignedProposal) > 0 && len(splitResult) > 2 && stringSignedProposal[len(stringSignedProposal)-1] == '}' { // Если есть входные парметры
		nameFunc = splitResult[len(splitResult)-2]
		nameChaincode = splitResult[len(splitResult)-3]
	} else if len(splitResult) > 1 { // Если нет входных парметров
		nameFunc = splitResult[len(splitResult)-1]
		nameChaincode = splitResult[len(splitResult)-2]
	}

	// TODO Убрать
	log.Info("nameChaincode: ", nameChaincode)
	log.Info("name: ", name)
	log.Info("nameFunc: ", nameFunc)
	log.Info("function: ", function)

	if name == nameChaincode {
		if len(function) == 0 {
			return true, nil
		}

		return function == nameFunc, nil
	}

	return false, nil
}

func CreateError(code uint, message string) error {
	baseError := cc_errors.BaseError{
		Code:    code,
		Message: message,
	}

	return createError(&baseError)
}

func CreateErrorWithData(code uint, message, data string) error {
	baseError := cc_errors.BaseError{
		Code:    code,
		Message: message,
		Data:    data,
	}

	return createError(&baseError)
}

func CheckArgs(args string, request interface{}) error {
	err := json.Unmarshal([]byte(args), &request)

	if err != nil {
		return CreateError(cc_errors.ErrorValidateDefault, fmt.Sprintf("Ошибка валидации: %s", err.Error())) // TODO
	}

	err = validation.Validate.Struct(request)
	if err != nil {
		return CreateError(cc_errors.ErrorValidateDefault, fmt.Sprintf("Ошибка валидации: %s", err.Error())) // TODO
	}

	requestInterface, ok := request.(interface{ SetDefaults() })

	if ok {
		requestInterface.SetDefaults()
	}

	return nil
}

func GetTimestamp(stub shim.ChaincodeStubInterface) (int64, error) {
	timestamp, err := stub.GetTxTimestamp()
	if err != nil {
		return 0, CreateError(cc_errors.ErrorGetTxTime, fmt.Sprintf("Ошибка при получении времени создания транзакции. %s", err.Error()))
	}

	time := timestamp.GetSeconds()*1000 + int64(timestamp.GetNanos()/1e6)

	return time, nil
}

func createError(baseError *cc_errors.BaseError) error {
	byteError, err := json.Marshal(baseError)
	if err != nil {
		return errors.New(
			fmt.Sprintf("{\"code\": %d, \"message\": \"Ошибка при формирование структуры ошибки. %s\", \"data\": \"\"}",
				cc_errors.ErrorDefault, err.Error()))
	}

	return errors.New(string(byteError))
}

func GetContractCommission(contract models.CurrencyExchangeContract, amountOutput int64) float64 {
	calcCommission := float64(amountOutput) * contract.Price * contract.FractionalCommission / (1 - contract.FractionalCommission)

	if contract.MaxCommission == 0 {
		return calcCommission
	}

	return math.Min(calcCommission, float64(contract.MaxCommission))
}
