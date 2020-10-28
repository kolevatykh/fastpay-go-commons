package test_helper

import (
	"github.com/SolarLabRU/fastpay-go-commons/test-helper/proxy_cid"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

var txCtxHolder = struct {
	stub     shim.ChaincodeStubInterface
	cidProxy *proxy_cid.ProxyClientID
}{}

// Реализация мока contractapi.TransactionContextInterface для юнит тестов.
// Проксирует cid.ClientIdentity через proxy_cid.ProxyClientID.
// Корректность поведения не гарантриуется при параллельном выполнении тестов.
type MockTransactionContext struct{}

func (*MockTransactionContext) GetStub() shim.ChaincodeStubInterface {
	return txCtxHolder.stub
}

// Возвращает глобальный прокси *proxy_cid.ProxyClientID.
func (*MockTransactionContext) GetClientIdentity() cid.ClientIdentity {
	return txCtxHolder.cidProxy
}

func (*MockTransactionContext) SetStub(stub shim.ChaincodeStubInterface) {
	txCtxHolder.stub = stub
}

// Устанавливает ci в Target глобального прокси proxy_cid.ProxyClientID.
func (*MockTransactionContext) SetClientIdentity(ci cid.ClientIdentity) {
	if txCtxHolder.cidProxy == nil {
		txCtxHolder.cidProxy = &proxy_cid.ProxyClientID{}
	}
	txCtxHolder.cidProxy.Target = ci
}

// Устанавливает глобальный прокси.
// Корректность поведения не гарантриуется при параллельном выполнении тестов.
func SetProxyClientIdentity(proxyClientID *proxy_cid.ProxyClientID) {
	txCtxHolder.cidProxy = proxyClientID
}

// Возвращает глобальный прокси.
// Корректность поведения не гарантриуется при параллельном выполнении тестов.
func GetProxyClientIdentity() *proxy_cid.ProxyClientID {
	return txCtxHolder.cidProxy
}
