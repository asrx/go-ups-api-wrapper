package ComplexType

import (
	. "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	"encoding/xml"
)

type VoidShipmentResponse struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/Void/v1.1 VoidShipmentResponse"`

	Response *ResponseType `xml:"Response,omitempty"`

	SummaryResult *SummaryResult `xml:"SummaryResult,omitempty"`

	PackageLevelResult []*PackageLevelResult `xml:"PackageLevelResult,omitempty"`
}

type SummaryResult struct {
	Status *CodeDescriptionType `xml:"Status,omitempty"`
}
