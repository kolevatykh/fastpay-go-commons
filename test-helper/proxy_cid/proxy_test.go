package proxy_cid

import (
	"crypto/x509"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-chaincode-go/pkg/attrmgr"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	clientID, _ = cid.New(buildStub())
)

func TestProxyClientID_AssertAttributeValue(t *testing.T) {
	t.Parallel()
	notExistingKey := "not exists key"
	notExistingValue := "not exists val"
	affiliationKey := "hf.Affiliation"
	vtbValue := "vtb"
	notVtbValue := "not vtb"
	typeKey := "hf.Type"
	typeValue := "peer"
	proxy := &ProxyClientID{}

	err := proxy.AssertAttributeValue(typeKey, typeValue)
	assert.Equal(t, ErrNoAttrs, err)
	err = proxy.AssertAttributeValue(affiliationKey, vtbValue)
	assert.Equal(t, ErrNoAttrs, err)
	err = proxy.AssertAttributeValue(affiliationKey, notVtbValue)
	assert.Equal(t, ErrNoAttrs, err)
	err = proxy.AssertAttributeValue(notExistingKey, notExistingValue)
	assert.Equal(t, ErrNoAttrs, err)

	proxy.Target = clientID

	err = proxy.AssertAttributeValue(notExistingKey, notExistingValue)
	var notFound *ErrAssertAttrNotFound
	assert.True(t, errors.As(err, &notFound))
	assert.Equal(t, notExistingKey, notFound.AttrName)
	err = clientID.AssertAttributeValue(notExistingKey, notExistingValue)
	var notFoundOrigin *ErrAssertAttrNotFound
	assert.False(t, errors.As(err, &notFoundOrigin))
	assert.Equal(t, notFound.Error(), err.Error())

	wrongAff := "wrong"
	err = proxy.AssertAttributeValue(affiliationKey, wrongAff)
	var notEqual *ErrAssertAttrNotEqual
	assert.True(t, errors.As(err, &notEqual))
	assert.Equal(t, affiliationKey, notEqual.AttrName)
	err = clientID.AssertAttributeValue(affiliationKey, wrongAff)
	var notEqualOrigin *ErrAssertAttrNotEqual
	assert.False(t, errors.As(err, &notEqualOrigin))
	assert.Equal(t, notEqual.Error(), err.Error())

	proxy.Attrs = &attrmgr.Attributes{Attrs: make(map[string]string)}

	proxy.Attrs.Attrs[notExistingKey] = notExistingValue
	assert.Nil(t, proxy.AssertAttributeValue(notExistingKey, notExistingValue))
	err = clientID.AssertAttributeValue(notExistingKey, notExistingValue)
	assert.NotNil(t, err)

	assert.Nil(t, proxy.AssertAttributeValue(affiliationKey, vtbValue))
	assert.Nil(t, clientID.AssertAttributeValue(affiliationKey, vtbValue))
	proxy.Attrs.Attrs[affiliationKey] = notVtbValue
	assert.Nil(t, proxy.AssertAttributeValue(affiliationKey, notVtbValue))
	assert.Nil(t, clientID.AssertAttributeValue(affiliationKey, vtbValue))
	assert.Nil(t, proxy.AssertAttributeValue(typeKey, typeValue))

	proxy.Target = nil
	assert.Nil(t, proxy.AssertAttributeValue(affiliationKey, notVtbValue))
	err = proxy.AssertAttributeValue(typeKey, typeValue)
	var anotherNotFound *ErrAssertAttrNotFound
	assert.True(t, errors.As(err, &anotherNotFound))
}

func TestProxyClientID_GetAttributeValue(t *testing.T) {
	t.Parallel()
	proxy := &ProxyClientID{}
	val, found, err := proxy.GetAttributeValue("a1")
	assert.False(t, found)
	assert.Equal(t, "", val)
	assert.Equal(t, ErrNoAttrs, err)

	affiliationKey := "hf.Affiliation"
	vtbValue := "vtb"
	proxy.Target = clientID
	actual, found, err := proxy.GetAttributeValue(affiliationKey)
	assert.Nil(t, err)
	assert.True(t, found)
	fromTarget, found, err := clientID.GetAttributeValue(affiliationKey)
	assert.Nil(t, err)
	assert.True(t, found)
	assert.Greater(t, len(actual), 0)
	assert.Equal(t, fromTarget, actual)
	assert.Equal(t, fromTarget, vtbValue)

	proxy.Attrs = &attrmgr.Attributes{Attrs: make(map[string]string)}
	actual, found, err = proxy.GetAttributeValue(affiliationKey)
	assert.Nil(t, err)
	assert.True(t, found)
	fromTarget, found, err = clientID.GetAttributeValue(affiliationKey)
	assert.Nil(t, err)
	assert.True(t, found)
	assert.Greater(t, len(actual), 0)
	assert.Equal(t, fromTarget, actual)
	assert.Equal(t, fromTarget, vtbValue)

	notVtbValue := "not vtb"
	proxy.Attrs.Attrs[affiliationKey] = notVtbValue
	actual, found, err = proxy.GetAttributeValue(affiliationKey)
	assert.Nil(t, err)
	assert.True(t, found)
	fromTarget, found, err = clientID.GetAttributeValue(affiliationKey)
	assert.Nil(t, err)
	assert.True(t, found)
	assert.Equal(t, actual, notVtbValue)
	assert.Equal(t, fromTarget, vtbValue)

	typeKey := "hf.Type"
	typeValue := "peer"
	actual, found, err = proxy.GetAttributeValue(typeKey)
	assert.Nil(t, err)
	assert.True(t, found)
	fromTarget, found, err = clientID.GetAttributeValue(typeKey)
	assert.Nil(t, err)
	assert.True(t, found)
	assert.Equal(t, actual, typeValue)
	assert.Equal(t, fromTarget, typeValue)

	proxy.Target = nil
	actual, found, err = proxy.GetAttributeValue(affiliationKey)
	assert.Nil(t, err)
	assert.True(t, found)
	assert.Equal(t, actual, notVtbValue)
}

func TestProxyClientID_GetID(t *testing.T) {
	t.Parallel()
	proxy := &ProxyClientID{}
	id, err := proxy.GetID()
	assert.Equal(t, "", id)
	assert.Equal(t, ErrNoID, err)

	proxy.Target = clientID
	actual, err := proxy.GetID()
	assert.Nil(t, err)
	fromTarget, err := clientID.GetID()
	assert.Nil(t, err)
	assert.Greater(t, len(actual), 0)
	assert.Equal(t, fromTarget, actual)

	own := "id"
	proxy.ID = &own
	actual, err = proxy.GetID()
	assert.Nil(t, err)
	assert.Greater(t, len(actual), 0)
	assert.Equal(t, own, actual)

	proxy.Target = nil
	actual, err = proxy.GetID()
	assert.Nil(t, err)
	assert.Greater(t, len(actual), 0)
	assert.Equal(t, own, actual)
}

func TestProxyClientID_GetMSPID(t *testing.T) {
	t.Parallel()
	proxy := &ProxyClientID{}
	mspID, err := proxy.GetMSPID()
	assert.Equal(t, "", mspID)
	assert.Equal(t, ErrNoMSPID, err)

	proxy.Target = clientID
	actual, err := proxy.GetMSPID()
	assert.Nil(t, err)
	fromTarget, err := clientID.GetMSPID()
	assert.Nil(t, err)
	assert.Greater(t, len(actual), 0)
	assert.Equal(t, fromTarget, actual)

	own := "msp id"
	proxy.MspID = &own
	actual, err = proxy.GetMSPID()
	assert.Nil(t, err)
	assert.Greater(t, len(actual), 0)
	assert.Equal(t, own, actual)

	proxy.Target = nil
	actual, err = proxy.GetMSPID()
	assert.Nil(t, err)
	assert.Greater(t, len(actual), 0)
	assert.Equal(t, own, actual)
}

func TestProxyClientID_GetX509Certificate(t *testing.T) {
	t.Parallel()
	proxy := &ProxyClientID{}
	cert, err := proxy.GetX509Certificate()
	assert.Nil(t, cert)
	assert.Equal(t, ErrNoCert, err)

	proxy.Target = clientID
	actual, err := proxy.GetX509Certificate()
	assert.Nil(t, err)
	fromTarget, err := clientID.GetX509Certificate()
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, fromTarget, actual)

	own := &x509.Certificate{}
	proxy.Cert = own
	actual, err = proxy.GetX509Certificate()
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, own, actual)

	proxy.Target = nil
	actual, err = proxy.GetX509Certificate()
	assert.Nil(t, err)
	assert.NotNil(t, actual)
	assert.Equal(t, own, actual)
}

func buildStub() (stub *shimtest.MockStub) {
	stub = shimtest.NewMockStub("myname", &contractapi.ContractChaincode{})
	sid := &msp.SerializedIdentity{
		Mspid: "mspID",
		IdBytes: []byte(`-----BEGIN CERTIFICATE-----
MIICujCCAmCgAwIBAgIUdx26WDSNm3qwsURZKIOOAzWs9uEwCgYIKoZIzj0EAwIw
bjELMAkGA1UEBhMCVVMxFzAVBgNVBAgTDk5vcnRoIENhcm9saW5hMRQwEgYDVQQK
EwtIeXBlcmxlZGdlcjEPMA0GA1UECxMGRmFicmljMR8wHQYDVQQDExZjYS5mYWJy
aWMubGludXhjdGwuY29tMB4XDTIwMDMxNzEyMzgwMFoXDTIxMDMxNzEyNDMwMFow
ZzELMAkGA1UEBhMCVVMxFzAVBgNVBAgTDk5vcnRoIENhcm9saW5hMRQwEgYDVQQK
EwtIeXBlcmxlZGdlcjEZMAsGA1UECxMEcGVlcjAKBgNVBAsTA3Z0YjEOMAwGA1UE
AxMFcGVlcjEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASGakVFMf1DjtAtgQ/N
9OwIeKlt/5ovRM3+05Mqpj0zped19YiUGtLHEG3w2QRnkJfzqYIG+xCST7rFXfdj
DM+Co4HiMIHfMA4GA1UdDwEB/wQEAwIHgDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQW
BBR0DRQxTJ519P/q40gsUN5rLbc8tTAfBgNVHSMEGDAWgBRak/KzSst325PeMnxP
5o3aukjHVjAkBgNVHREEHTAbghNmYXN0cGF5LWFkbWluLXBlZXIxhwSsGwhCMFkG
CCoDBAUGBwgBBE17ImF0dHJzIjp7ImhmLkFmZmlsaWF0aW9uIjoidnRiIiwiaGYu
RW5yb2xsbWVudElEIjoicGVlcjEiLCJoZi5UeXBlIjoicGVlciJ9fTAKBggqhkjO
PQQDAgNIADBFAiEAkOSk/HCnG4nvbBjV+xPPA328UOVbULSoem8R6Y4VE5YCIAVs
sud4CPPL080UzfaQPVIXKOygzXmfbGIvaGbMFjS7
-----END CERTIFICATE-----
`),
	}
	b, err := proto.Marshal(sid)
	if err != nil {
		panic(err)
	}

	stub.Creator = b
	return
}
