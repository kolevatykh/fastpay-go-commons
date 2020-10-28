package proxy_cid

import (
	"crypto/x509"
	"github.com/hyperledger/fabric-chaincode-go/pkg/attrmgr"
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
)

// ProxyClientID реализует cid.ClientIdentity, позволяя мокать выбранные данные.
// Возвращает собственные данные при их наличии.
// Иначе - проксирует запрос в целевой cid.ClientIdentity, при наличии оной.
// Исключением является AssertAttributeValue, см документацию метода для уточнения.
// НЕ является потокобезопасным.
type ProxyClientID struct {

	// Целевая cid.ClientID для проксирования.
	// Не является обязательной, для хранения собственных данных.
	// При отсутствии и собственных данных и Target будут возвращены соответствующие ошибки.
	Target cid.ClientIdentity
	ID     *string
	MspID  *string
	Cert   *x509.Certificate
	Attrs  *attrmgr.Attributes
}

// Возвращает ID. При отсутствии проксирует запрос в Target. Возвращает ErrNoID при отсутствии обоих.
func (pcid *ProxyClientID) GetID() (string, error) {
	if pcid.ID != nil {
		return *pcid.ID, nil
	}
	if pcid.Target != nil {
		return pcid.Target.GetID()
	}
	return "", ErrNoID
}

// Возвращает MspID. При отсутствии проксирует запрос в Target. Возвращает ErrNoMSPID при отсутствии обоих.
func (pcid *ProxyClientID) GetMSPID() (string, error) {
	if pcid.MspID != nil {
		return *pcid.MspID, nil
	}
	if pcid.Target != nil {
		return pcid.Target.GetMSPID()
	}
	return "", ErrNoMSPID
}

// Возвращает Cert. При отсутствии проксирует запрос в Target. Возвращает ErrNoCert при отсутствии обоих.
func (pcid *ProxyClientID) GetX509Certificate() (*x509.Certificate, error) {
	if pcid.Cert != nil {
		return pcid.Cert, nil
	}
	if pcid.Target != nil {
		return pcid.Target.GetX509Certificate()
	}
	return nil, ErrNoCert
}

// Выполняет поиск атрибута в Attrs.
// В случае отсутствия искомого атрибута или хранилища Attrs проксирует запрос в Target.
// В случае если Attrs и Target == nil возвращает ErrNoAttrs.
func (pcid *ProxyClientID) GetAttributeValue(attrName string) (value string, found bool, err error) {
	if pcid.Attrs == nil && pcid.Target == nil {
		err = ErrNoAttrs
		return
	}
	if pcid.Attrs != nil {
		value, found, err = pcid.Attrs.Value(attrName)
	}
	if (!found || err != nil) && pcid.Target != nil {
		value, found, err = pcid.Target.GetAttributeValue(attrName)
	}
	return
}

// Выполняет поиск атрибута в соответствии с GetAttributeValue.
// При успешном нахождении атрибута проверка будет проведена вручную без проксирования в Target.
// При отсутствии атрибута возвращает ErrAssertAttrNotFound.
// Если ожидаемый и найденный атрибут не совпадают будет возвращена ErrAssertAttrNotEqual.
// Гарантируется, что текстовое представление возвращаемых ошибок соответстует
// представлениям проксируемых ошибок.
func (pcid *ProxyClientID) AssertAttributeValue(attrName, attrValue string) error {
	val, ok, err := pcid.GetAttributeValue(attrName)
	if err != nil {
		return err
	}
	if !ok {
		return &ErrAssertAttrNotFound{AttrName: attrName}
	}
	if val != attrValue {
		return &ErrAssertAttrNotEqual{
			Actual:   val,
			AttrName: attrName,
			Expected: attrValue,
		}
	}
	return nil
}
