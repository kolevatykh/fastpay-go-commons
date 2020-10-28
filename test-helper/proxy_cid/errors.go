package proxy_cid

import (
	"errors"
	"fmt"
)

var (

	// Означает, что ProxyClientID не содержит ни ProxyClientID.ID ни ProxyClientID.Target.
	ErrNoID = errors.New("no id")

	// Означает, что ProxyClientID не содержит ни ProxyClientID.Cert ни ProxyClientID.Target.
	ErrNoCert = errors.New("no cert")

	// Означает, что ProxyClientID не содержит ни ProxyClientID.MspID ни ProxyClientID.Target.
	ErrNoMSPID = errors.New("no msp id")

	// Означает, что ProxyClientID не содержит ни ProxyClientID.Attrs ни ProxyClientID.Target.
	ErrNoAttrs = errors.New("no attrs")
)

// Означает, что атрибут не был найден при попытке проверить значение искомого атрибута.
type ErrAssertAttrNotFound struct {

	// Имя искомого атрибута
	AttrName string
}

func (err *ErrAssertAttrNotFound) Error() string {
	return fmt.Sprintf("attribute '%s' was not found", err.AttrName)
}

// Означает, что искомый атрибут не соответствует данному эталону.
type ErrAssertAttrNotEqual struct {

	// Найденное значение атрибута.
	Actual string

	// Имя атрибута.
	AttrName string

	// Эталонное значение.
	Expected string
}

func (err *ErrAssertAttrNotEqual) Error() string {
	return fmt.Sprintf(
		"attribute '%s' equals '%s', not '%s'",
		err.AttrName,
		err.Actual,
		err.Expected,
	)
}
