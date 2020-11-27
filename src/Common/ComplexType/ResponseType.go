package ComplexType

import "encoding/xml"

type ResponseType struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/Common/v1.0 Response"`

	ResponseStatus *CodeDescriptionType `xml:"ResponseStatus,omitempty"`

	Alert []*CodeDescriptionType `xml:"Alert,omitempty"`

	AlertDetail []*DetailType `xml:"AlertDetail,omitempty"`

	TransactionReference *TransactionReferenceType `xml:"TransactionReference,omitempty"`
}
