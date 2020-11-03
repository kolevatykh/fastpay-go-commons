package base

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/SolarLabRU/fastpay-go-commons/cc-errors"
	"github.com/SolarLabRU/fastpay-go-commons/crypto"
	"github.com/SolarLabRU/fastpay-go-commons/enums/access-role-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/state_enum"
	"github.com/SolarLabRU/fastpay-go-commons/logger"
	"github.com/SolarLabRU/fastpay-go-commons/models"
	"github.com/SolarLabRU/fastpay-go-commons/requests"
	"github.com/SolarLabRU/fastpay-go-commons/responses"
	"github.com/SolarLabRU/fastpay-go-commons/validation"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

var Logger *logger.Logger

func init() {
	Logger = new(logger.Logger)
	Logger.Init()

	// TODO убрать
	Logger.Info("GO_ENV: ", os.Getenv("GO_ENV"))
	Logger.Info("CORE_CHAINCODE_LOGGING_LEVEL: ", os.Getenv("CORE_CHAINCODE_LOGGING_LEVEL"))
	Logger.Info("-----------------All_ENV---------------------")
	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}
	Logger.Error("Error")
	Logger.Info("Info")
	Logger.Debug("Debug")
}

const (
	ChaincodeBankName       = "banks"
	ChaincodeAccountsName   = "accounts"
	ChaincodeClientBankName = "client-banks"
	compositeExpSignKey     = "typeExpSign~sign"
	EventBatchName          = "EventBatch"
)

// Метод получения банка отправителя
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

func GetClientBank(ctx contractapi.TransactionContextInterface, address string) (*responses.ClientBankItemResponse, error) {
	stub := ctx.GetStub()

	request := requests.GetClientBankByAddressRequest{
		Address: address,
	}

	response, err := InvokeChaincode(stub, ChaincodeClientBankName, "getClientBankById", request)
	if err != nil {
		return nil, err
	}

	var bankResponse responses.ClientBankResponse
	err = json.Unmarshal(response, &bankResponse)

	if err != nil {
		return nil, CreateError(cc_errors.ErrorJsonUnmarshal, fmt.Sprintf("Ошибка десерилизации ответа вовремя вызова чейнкода client-banks. %s", err.Error()))
	}

	return &bankResponse.Data, nil
}

// Метод получения адреса банка из сертификата
func GetSenderAddressFromCertificate(identity cid.ClientIdentity) (string, error) {
	address, isFound, _ := identity.GetAttributeValue("address")

	//address, isFound, _ = func() (string, bool, error) { return "263093b1c21f98c5f9b6433bf9bbb97bb87b6e79", true, nil }() // TODO Убрать

	if !isFound {
		return "", CreateError(cc_errors.ErrorCertificateNotValid, "Отсутвует атрибут address в сертификате")
	}

	return address, nil
}

// Метод вызова внешнего чейнкода
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
		fmt.Println("Ошибка при вызове чейнкода. Payload: ", string(response.GetPayload()))
		return nil, parseErrorFromAnotherChaincode(response.GetMessage())
	}

	return response.GetPayload(), nil
}

// Метод вызова внешнего чейнкода без входных параметров
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

// Метод получения банка по адресу банка и mspId
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

// Метод проверки доступности банка отправителя
func SenderBankIsAvailable(ctx contractapi.TransactionContextInterface) error {
	bank, _ := GetSenderBank(ctx)
	return SenderBankIsAvailableWithBank(ctx, bank)
}

// Метод проверки доступности банка отправителя без получения банка
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

func SenderClientBankIsAvailable(ctx contractapi.TransactionContextInterface, senderClientBank *responses.ClientBankItemResponse, address string) error {
	if senderClientBank == nil {
		var err error = nil
		senderClientBank, err = GetClientBank(ctx, address)
		if err != nil {
			return err
		}
	}

	addressSender, err := GetSenderAddressFromCertificate(ctx.GetClientIdentity())

	if err != nil {
		return nil
	}

	if senderClientBank == nil || senderClientBank.State != state_enum.Available {
		return CreateError(cc_errors.ErrorClientBankNotAvailable, "Клиентский банк отправителя не доступен")
	}
	if senderClientBank.Owner != addressSender {
		return CreateError(cc_errors.ErrorClientBankOwnerNotEqualSender,
			fmt.Sprintf("Опорный банк(отправитель)(%s) не является владельцем клиентского банка(%s)", addressSender, senderClientBank.Address))
	}

	return nil
}

func CheckTechnicalAccountSign(ctx contractapi.TransactionContextInterface, technicalSignRequest requests.TechnicalSignRequest, bankSender *models.Bank) error {

	if bankSender == nil {
		var err error = nil
		bankSender, err = GetSenderBank(ctx)
		if err != nil {
			return err
		}
	}

	return CheckTechnicalAccountSignByAddress(ctx, technicalSignRequest, bankSender.Address)
}

func CheckTechnicalAccountSignByAddress(ctx contractapi.TransactionContextInterface, technicalSignRequest requests.TechnicalSignRequest, address string) error {

	err := CheckSign(technicalSignRequest.TechnicalAddress, technicalSignRequest.TechnicalMsgHash, technicalSignRequest.TechnicalSig)
	if err != nil {
		return err
	}

	// TODO Убрать проверку если в сертификате не будет указыватся адрес банка отправителя
	if address != technicalSignRequest.TechnicalAddress {
		return CreateError(cc_errors.ErrorAccountTechnicalNotEqlSender,
			"Адрес банка отправителя не совпадает с адресом технического аккаунта")
	}

	return nil
}

func CheckClientBankTechnicalSignAndAvailable(ctx contractapi.TransactionContextInterface, request requests.TechnicalSignRequest) error {

	err := CheckSign(request.TechnicalAddress, request.TechnicalMsgHash, request.TechnicalSig)
	if err != nil {
		return err
	}

	err = SenderClientBankIsAvailable(ctx, nil, request.TechnicalAddress)
	if err != nil {
		return err
	}

	return nil
}

func CheckAccessAndAvailableWithBank(ctx contractapi.TransactionContextInterface, bank *models.Bank, role access_role_enum.AccessRole) error {
	return CheckAccessWithBank(ctx, bank, role, true)
}

func CheckAccessAndAvailable(ctx contractapi.TransactionContextInterface, role access_role_enum.AccessRole) error {
	return CheckAccessWithBank(ctx, nil, role, true)
}

// Метод проверки доступа к методу чейнкода с переданым банком отправителя
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

	currentRoles := getRoles(bank)
	result := currentRoles & role

	if result == 0 {
		return CreateError(cc_errors.ErrorForbidden, "Недостаточно прав для вызова метода")
	}

	return nil
}

// Метод проверки, что вызваемый метод вызывался другим чейнкодом
func CheckCalledChaincode(stub shim.ChaincodeStubInterface, name, function string) (bool, error) {
	nameChaincode, nameFunc, err := GetChaincodeNameCalled(stub)
	if err != nil {
		return false, err
	}

	// TODO Убрать
	Logger.Info("nameChaincode: ", nameChaincode)
	Logger.Info("name: ", name)
	Logger.Info("nameFunc: ", nameFunc)
	Logger.Info("function: ", function)

	if strings.Contains(nameChaincode, name) {
		if len(function) == 0 {
			return true, nil
		}

		return function == nameFunc, nil
	}

	return false, nil
}

func GetChaincodeNameCalled(stub shim.ChaincodeStubInterface) (string, string, error) {
	nameChaincode := ""
	nameFunc := ""

	signedProposal, err := stub.GetSignedProposal()
	if err != nil {
		return nameChaincode, nameFunc, err
	}
	stringSignedProposal := string(signedProposal.GetProposalBytes())
	filterString := regexp.MustCompile(`[^a-zA-Z 0-9\n_-]`).ReplaceAllString(stringSignedProposal, "")
	filterString2 := regexp.MustCompile(`[\t\r\n]+`).ReplaceAllString(strings.TrimSpace(filterString), "\n")
	splitResult := strings.Split(filterString2, "\n")

	if len(stringSignedProposal) > 0 && len(splitResult) > 2 && stringSignedProposal[len(stringSignedProposal)-1] == '}' { // Если есть входные парметры
		nameFunc = splitResult[len(splitResult)-2]
		nameChaincode = splitResult[len(splitResult)-3]
	} else if len(splitResult) > 1 { // Если нет входных парметров
		nameFunc = splitResult[len(splitResult)-1]
		nameChaincode = splitResult[len(splitResult)-2]
	}

	return nameChaincode, nameFunc, nil
}

// Метод создания структуры ошибки
func CreateError(code int, message string) error {
	baseError := cc_errors.BaseError{
		Code:    code,
		Message: message,
	}

	return createError(&baseError)
}

// Метод создания структуры ошибки с доп. информацией
func CreateErrorWithData(code int, message, data string) error {
	baseError := cc_errors.BaseError{
		Code:    code,
		Message: message,
		Data:    data,
	}

	return createError(&baseError)
}

// Метод проверки входных параметров
func CheckArgs(args string, request interface{}) error {
	err := json.Unmarshal([]byte(args), &request)

	if err != nil {
		return CreateError(cc_errors.ErrorValidateDefault, fmt.Sprintf("Ошибка валидации: %s", err.Error())) // TODO
	}

	_, err = validation.ValidateStruct(request)
	if err != nil {
		return processValidationError(err)
	}

	requestInterface, ok := request.(interface{ SetDefaults() })

	if ok {
		requestInterface.SetDefaults()
	}

	return nil
}

// Проверка аккаунта на соответствие валюты по его адресу
func CheckAccountCurrencyCode(stub shim.ChaincodeStubInterface, address string, currencyCode int) error {
	account, err := GetAccountByAddress(stub, address)
	if err != nil {
		return err
	}

	if account.CurrencyCode != currencyCode {
		return CreateError(cc_errors.ErrorIncorrectCurrencyCodeAccount, fmt.Sprintf("Аккаунт %s не соответствует указанной валюте", account.Address))
	}

	return nil
}

// Метод публикации события в чейнкоде
func PublicEvent(stub shim.ChaincodeStubInterface, event interface{}, eventName string) error {
	return PublicEvents(stub, models.EventBatch{
		Events: []models.EventBatchItem{{
			EventName: eventName,
			Data:      event,
		}},
	})
}

// Метод публикации массива событий в чейнкоде
func PublicEvents(stub shim.ChaincodeStubInterface, events interface{}) error {
	eventsAsBytes, err := json.Marshal(events)
	if err != nil {
		return CreateError(cc_errors.ErrorJsonMarshal, fmt.Sprintf("Ошибка при сериализации событий. %s", err.Error()))
	}

	err = stub.SetEvent(EventBatchName, eventsAsBytes)
	if err != nil {
		return err
	}

	return nil
}

// Метод получения времени созджания транзакции(Из заголовка канала)
func GetTimestamp(stub shim.ChaincodeStubInterface) (int64, error) {
	timestamp, err := stub.GetTxTimestamp()
	if err != nil {
		return 0, CreateError(cc_errors.ErrorGetTxTime, fmt.Sprintf("Ошибка при получении времени создания транзакции. %s", err.Error()))
	}

	time := timestamp.GetSeconds()*1000 + int64(timestamp.GetNanos()/1e6)

	return time, nil
}

// Метод получения комиссии
func GetContractCommission(contract models.CurrencyExchangeContract, amountOutput int64) float64 {
	calcCommission := float64(amountOutput) * contract.Price * contract.FractionalCommission / (1 - contract.FractionalCommission)

	if contract.MaxCommission == 0 {
		return calcCommission
	}

	return math.Min(calcCommission, float64(contract.MaxCommission))
}

func ParseError(err error) *cc_errors.BaseError {
	var baseError *cc_errors.BaseError

	errUnmarshal := json.Unmarshal([]byte(err.Error()), &baseError)

	if errUnmarshal != nil {

		return &cc_errors.BaseError{
			Code:    cc_errors.ErrorDefault,
			Message: "Ошибка парсинга информации об ошибки: " + err.Error(),
			Data:    "",
		}
	}

	return baseError
}

func CheckTransactionError(err error) (error, *cc_errors.BaseError) {
	if err == nil {
		return nil, nil
	}

	baseError := ParseError(err)

	errorsRejectTransactions := []int{cc_errors.ErrorDefault, cc_errors.ErrorPutState, cc_errors.ErrorGetState, cc_errors.ErrorDeleteState,
		cc_errors.ErrorCreateCompositeKey, cc_errors.ErrorJsonMarshal, cc_errors.ErrorJsonUnmarshal}

	if Contains(errorsRejectTransactions, baseError.Code) {
		return err, baseError
	}

	return nil, baseError
}

func Contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// Метод создания структуры ошибки
func createError(baseError *cc_errors.BaseError) error {
	byteError, err := json.Marshal(baseError)
	if err != nil {
		return errors.New(
			fmt.Sprintf("{\"code\": %d, \"message\": \"Ошибка при формирование структуры ошибки. %s\", \"data\": \"\"}",
				cc_errors.ErrorDefault, err.Error()))
	}

	return errors.New(string(byteError))
}

// Метод получения ролей
func getRoles(bank *models.Bank) access_role_enum.AccessRole {
	roles := access_role_enum.Bank

	for _, v := range bank.Roles {
		roles |= v
	}

	return roles
}

// Метод серилизации ошибки при вызове чейнкода
func parseErrorFromAnotherChaincode(message string) error {
	var baseError cc_errors.BaseError

	err := json.Unmarshal([]byte(message), &baseError)

	if err != nil {
		return CreateError(cc_errors.ErrorDefault, fmt.Sprintf("Ошибка при вызове чейнкода: %s", message))
	}

	return CreateError(baseError.Code, fmt.Sprintf("Ошибка при вызове чейнкода: %s", baseError.Message))

}

// Метод проверки сигнатуры
func CheckSign(address, msgHash string, sign requests.SignDto) error {
	if msgHash == "" || sign.R == "" || sign.S == "" || sign.V == 0 {
		return CreateError(cc_errors.ErrorValidateDefault, "Сигнатура не передана")
	}

	isSigned, err := crypto.IsSigned(address, msgHash, sign.R, sign.S, sign.V)

	if err != nil {
		return CreateError(cc_errors.ErrorSignVerify, fmt.Sprintf("Ошибка проверки сигнатуры. %s", err.Error()))
	}

	if !isSigned {
		return CreateError(cc_errors.ErrorSignVerify, "Ошибка проверки сигнатуры.")
	}

	return nil
}

// Метод проверки сигнатуры и ее времени действия
func CheckSignAndExpiration(stub shim.ChaincodeStubInterface, address, msgHash string, sign requests.SignDto, expiration int64, now int64) error {

	if expiration == 0 {
		return CreateError(cc_errors.ErrorValidateDefault, "Поле Exp не передано")
	}

	err := CheckSign(address, msgHash, sign)
	if err != nil {
		return err
	}

	err = CheckSignExpiration(stub, sign, expiration, now)
	if err != nil {
		return err
	}

	return nil
}

func CheckSignExpiration(stub shim.ChaincodeStubInterface, sign requests.SignDto, expiration int64, now int64) error {

	if now > expiration {
		return CreateError(cc_errors.ErrorExpirationSign, "Время подписи истекло")
	}

	isExist, err := GetExpirationSign(stub, sign)
	if err != nil {
		return err
	}

	if isExist {
		return CreateError(cc_errors.ErrorSignUsed, "Подпись уже была использована")
	}

	err = SaveExpirationSign(stub, sign, expiration)
	if err != nil {
		return err
	}

	return nil
}

func GetExpirationSign(stub shim.ChaincodeStubInterface, sig requests.SignDto) (bool, error) {
	signKey := fmt.Sprintf("%s_%s_%d", sig.R, sig.S, sig.V)

	expirationAsBytes, err := stub.GetState(signKey)
	if err != nil {
		return false, CreateError(cc_errors.ErrorGetState, fmt.Sprintf("Ошибка при получении времени действия сигнатуры. %s", err.Error()))
	}

	if len(expirationAsBytes) == 0 {
		return false, nil
	}
	return true, nil
}

func SaveExpirationSign(stub shim.ChaincodeStubInterface, sig requests.SignDto, expiration int64) error {

	signKey := fmt.Sprintf("%s_%s_%d", sig.R, sig.S, sig.V)

	err := stub.PutState(signKey, []byte(strconv.FormatInt(expiration, 10)))
	if err != nil {
		return CreateError(cc_errors.ErrorPutState, fmt.Sprintf("Ошибка даты окончания действия подписи. %s.", err.Error()))
	}

	return nil
}

// Метод очистки используемых подписей
// TODO По возможности изменить алгоритм очистки подписей, потому что он является не самым оптимальным(возможна ошибка PHANTOM_READ_CONFLICT)
func DeleteExpirationSign(stub shim.ChaincodeStubInterface) error {
	resultsIterator, err := stub.GetStateByPartialCompositeKey(compositeExpSignKey, []string{"exp_sign"})
	if err != nil {
		return CreateError(cc_errors.ErrorGetState, fmt.Sprintf("Ошибка создания композитного ключа:  %s.", err.Error()))
	}

	defer resultsIterator.Close()

	timestampNow, err := GetTimestamp(stub)
	if err != nil {
		return err
	}

	for resultsIterator.HasNext() {
		responseRange, err := resultsIterator.Next()
		if err != nil {
			return CreateError(cc_errors.ErrorGetState, fmt.Sprintf("Ошибка получения данных итератора из ответа по запросу к БД. %s", err.Error()))
		}

		expirationSign, err := strconv.ParseInt(string(responseRange.GetValue()), 10, 64)
		if err != nil {
			return CreateError(cc_errors.ErrorDefault, fmt.Sprintf("Ошибка при десериализации времени действия подписи. %s", err.Error()))
		}

		if timestampNow > expirationSign {
			err = stub.DelState(responseRange.Key)
			if err != nil {
				return CreateError(cc_errors.ErrorDeleteState, fmt.Sprintf("Ошибка при удалении используемой сигнатуры. %s", err.Error()))
			}
		}
	}

	return nil
}

// Обработка ошибки при валидации запроса
func processValidationError(err error) error {
	var errorData cc_errors.Error

	if strings.Contains(err.Error(), ";") {
		errorData = cc_errors.ErrorMessages[strings.Split(err.Error(), ";")[0]]
	} else {
		errorData = cc_errors.ErrorMessages[err.Error()]
	}

	if errorData.Code != 0 {
		return CreateError(errorData.Code, errorData.Message)
	}

	return CreateError(cc_errors.ErrorValidateDefault, fmt.Sprintf("Ошибка валидации: %s", err.Error()))
}

// Метод получение аккаунта с использованием кеша
func GetAccount(stub shim.ChaincodeStubInterface, addressOrIdentifier string, cache map[string]*models.Account) (*models.Account, error) {

	var account *models.Account
	var err error
	if cache == nil {
		if len(addressOrIdentifier) == 40 {
			account, err = GetAccountByAddress(stub, addressOrIdentifier)
			if err != nil {
				return nil, err
			}
		} else {
			account, err = GetAccountByIdentifier(stub, addressOrIdentifier)
			if err != nil {
				return nil, err
			}
		}

		return account, nil
	}

	var isAccount bool
	account, isAccount = cache[addressOrIdentifier]
	if !isAccount {
		var account *models.Account
		var err error

		if len(addressOrIdentifier) == 40 {
			account, err = GetAccountByAddress(stub, addressOrIdentifier)
			if err != nil {
				return nil, err
			}
		} else {
			account, err = GetAccountByIdentifier(stub, addressOrIdentifier)
			if err != nil {
				return nil, err
			}

			cache[account.Address] = account
		}
		cache[addressOrIdentifier] = account

		return account, nil
	}

	return account, nil
}

// Метод получения аккаунта по его идентификатору
func GetAccountByIdentifier(stub shim.ChaincodeStubInterface, identifier string) (*models.Account, error) {

	request := requests.GetByIdentifierRequest{
		Identifier: identifier,
	}

	response, err := InvokeChaincode(stub, ChaincodeAccountsName, "getByIdentifier", request)
	if err != nil {
		return nil, err
	}

	var accountResponse responses.AccountResponse
	err = json.Unmarshal(response, &accountResponse)
	if err != nil {
		return nil, CreateError(cc_errors.ErrorJsonUnmarshal, fmt.Sprintf("Ошибка десериализации ответа после вызова чейнкода accounts. %s", err.Error()))
	}

	return &accountResponse.Data, nil
}

// Метод получения аккаунта по его адресу
func GetAccountByAddress(stub shim.ChaincodeStubInterface, address string) (*models.Account, error) {

	request := requests.GetByAddressRequest{
		Address: address,
	}

	response, err := InvokeChaincode(stub, ChaincodeAccountsName, "getByAddress", request)
	if err != nil {
		return nil, err
	}

	var accountResponse responses.AccountResponse
	err = json.Unmarshal(response, &accountResponse)
	if err != nil {
		return nil, CreateError(cc_errors.ErrorJsonUnmarshal, fmt.Sprintf("Ошибка десериализации ответа после вызова чейнкода accounts. %s", err.Error()))
	}

	return &accountResponse.Data, nil
}

func GetChaincodeNameCurrencyCode() (string, string) {

	str := os.Getenv("CORE_CHAINCODE_ID_NAME")

	Logger.Info("CORE_CHAINCODE_ID_NAME: ", str)

	chaincodeName := ""
	currencyCode := ""

	if len(str) == 0 {
		return chaincodeName, currencyCode
	}

	res := strings.Split(str, "_")
	if len(res) == 1 {
		tempSplit := strings.Split(res[0], ":")
		chaincodeName = tempSplit[0]
	}
	if len(res) == 2 {
		chaincodeName = res[0]
		tempSplit := strings.Split(res[1], ":")
		if len(tempSplit[1]) != 64 {
			currencyCode = tempSplit[0]
		}
	}
	if len(res) == 3 {
		chaincodeName = res[0]
		currencyCode = res[1]
	}

	return chaincodeName, currencyCode
}

func GetChaincodeCurrencyCode() string {
	_, currencyCode := GetChaincodeNameCurrencyCode()

	return currencyCode
}

func GetChaincodeName() string {
	name, currencyCode := GetChaincodeNameCurrencyCode()

	if currencyCode == "" {
		return name
	}

	return name + "_" + currencyCode
}

// print the contents of the obj
func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}
