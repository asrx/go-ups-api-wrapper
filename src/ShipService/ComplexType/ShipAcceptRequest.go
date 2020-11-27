package ComplexType

import (
	"github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	"encoding/xml"
)

type ShipAcceptRequest struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/Ship/v1.0 ShipAcceptRequest"`

	Request *ComplexType.RequestType `xml:"Request,omitempty"`

	ShipmentDigest string `xml:"ShipmentDigest,omitempty"`
}
