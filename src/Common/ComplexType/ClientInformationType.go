package ComplexType

import "encoding/xml"

type ClientInformationType struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/Common/v1.0 ClientInformation"`

	Property []struct {
		Value string

		Key string `xml:"Key,attr,omitempty"`
	} `xml:"Property,omitempty"`
}
