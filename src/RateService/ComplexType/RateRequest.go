package ComplexType

import (
	. "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	"encoding/xml"
)

type RateRequest struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/Rate/v1.1 RateRequest"`

	Request *RequestType `xml:"Request,omitempty"`

	PickupType *CodeDescriptionType `xml:"PickupType,omitempty"`

	CustomerClassification *CodeDescriptionType `xml:"CustomerClassification,omitempty"`

	Shipment *ShipmentType `xml:"Shipment,omitempty"`
}

