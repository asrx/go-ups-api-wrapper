package ComplexType

import "encoding/xml"

type RequestType struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/Common/v1.0 Request"`

	RequestOption []string `xml:"RequestOption,omitempty"`

	SubVersion string `xml:"SubVersion,omitempty"`

	TransactionReference *TransactionReferenceType `xml:"TransactionReference,omitempty"`
}
