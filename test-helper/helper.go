package test_helper

import (
	"fmt"
	"github.com/SolarLabRU/fastpay-go-commons/helpers"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric-protos-go/peer"
	"testing"
)

func CallContractFunctionAndCheckError(t *testing.T, mockStub *shimtest.MockStub, arguments []string, callType string, expectedMessage string) {
	t.Helper()

	callContractFunctionAndCheckResponse(t, mockStub, arguments, callType, expectedMessage, "error")
}

func CallContractFunctionAndCheckSuccess(t *testing.T, mockStub *shimtest.MockStub, arguments []string, callType string, expectedMessage string) {
	t.Helper()

	callContractFunctionAndCheckResponse(t, mockStub, arguments, callType, expectedMessage, "success")
}

func CallContractFunctionAndCheckResponse(t *testing.T, mockStub *shimtest.MockStub, arguments []string, callType string, expectedMessage string, expectedType string) {
	t.Helper()

	args := helpers.StringArrayToByteArray(arguments)

	var response peer.Response

	if callType == initType {
		response = mockStub.MockInit(standardTxID, args)
	} else if callType == invokeType {
		response = mockStub.MockInvoke(standardTxID, args)
	} else if callType == queryType {
		response = mockStub.MockInvoke(standardTxID, args)
	} else {
		panic(fmt.Sprintf("Call type passed should be %s or %s. Value passed was %s", initType, invokeType, callType))
	}

	expectedResponse := shim.Success([]byte(expectedMessage))

	if expectedType == "error" {
		expectedResponse = shim.Error(expectedMessage)
	}

	assert.Equal(t, expectedResponse, response)
}

func SetCreator(t *testing.T, stub *shimtest.MockStub, mspID string, idbytes []byte) {
	sid := &msp.SerializedIdentity{Mspid: mspID, IdBytes: idbytes}
	b, err := proto.Marshal(sid)
	if err != nil {
		t.FailNow()
	}

	stub.Creator = b
}
