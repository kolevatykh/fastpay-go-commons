package test_helper

import (
	"fmt"
	"testing"

	"github.com/SolarLabRU/fastpay-go-commons/helpers"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/stretchr/testify/assert"
)

const (
	InvokeType   = "INVOKE"
	InitType     = "INIT"
	QueryType    = "QUERY"
	StandardTxID = "1234567890"
)

func CallContractFunctionAndCheckError(t *testing.T, mockStub *MockStub, arguments []string, callType string, expectedMessage string) {
	t.Helper()

	CallContractFunctionAndCheckResponse(t, mockStub, arguments, callType, expectedMessage, "error")
}

func CallContractFunctionAndCheckSuccess(t *testing.T, mockStub *MockStub, arguments []string, callType string, expectedMessage string) {
	t.Helper()

	CallContractFunctionAndCheckResponse(t, mockStub, arguments, callType, expectedMessage, "success")
}

func CallContractFunctionAndCheckResponse(t *testing.T, mockStub *MockStub, arguments []string, callType string, expectedMessage string, expectedType string) {
	t.Helper()

	args := helpers.StringArrayToByteArray(arguments)

	var response peer.Response

	if callType == InitType {
		response = mockStub.MockInit(StandardTxID, args)
	} else if callType == InvokeType {
		response = mockStub.MockInvoke(StandardTxID, args)
	} else if callType == QueryType {
		response = mockStub.MockInvoke(StandardTxID, args)
	} else {
		panic(fmt.Sprintf("Call type passed should be %s or %s. Value passed was %s", InitType, InvokeType, callType))
	}

	expectedResponse := shim.Success([]byte(expectedMessage))

	if expectedType == "error" {
		expectedResponse = shim.Error(expectedMessage)
	}

	assert.Equal(t, expectedResponse, response)
}

func SetCreator(t *testing.T, stub *MockStub, mspID string, idbytes []byte) {
	sid := &msp.SerializedIdentity{Mspid: mspID, IdBytes: idbytes}
	b, err := proto.Marshal(sid)
	if err != nil {
		t.FailNow()
	}

	stub.Creator = b
}
