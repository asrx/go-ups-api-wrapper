package ComplexType

import (
	"github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	"encoding/xml"
)

type ShipmentResponse struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/Ship/v1.0 ShipmentResponse"`

	Response *ComplexType.ResponseType `xml:"Response,omitempty"`

	ShipmentResults *ShipmentResultsType `xml:"ShipmentResults,omitempty"`
}