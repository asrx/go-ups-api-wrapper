package ComplexType

import (
	. "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	"encoding/xml"
)

type ShipmentRequest struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/Ship/v1.0 ShipmentRequest"`

	Request *RequestType `xml:"Request,omitempty"`

	Shipment *ShipmentType `xml:"Shipment,omitempty"`

	LabelSpecification *LabelSpecificationType `xml:"LabelSpecification,omitempty"`

	ReceiptSpecification *ReceiptSpecificationType `xml:"ReceiptSpecification,omitempty"`
}
