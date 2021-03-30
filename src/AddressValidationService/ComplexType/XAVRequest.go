package ComplexType

import (
	"encoding/xml"
	"github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
)

type XAVRequest struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/xav/v1.0 XAVRequest"`

	Request *ComplexType.RequestType `xml:"Request,omitempty"`

	RegionalRequestIndicator string `xml:"RegionalRequestIndicator,omitempty"`

	MaximumCandidateListSize string `xml:"MaximumCandidateListSize,omitempty"`

	AddressKeyFormat *AddressKeyFormatType `xml:"AddressKeyFormat,omitempty"`
}
