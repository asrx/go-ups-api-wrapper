package ComplexType

import (
	"github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	"encoding/xml"
)

type VoidShipmentRequest struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/Void/v1.1 VoidShipmentRequest"`

	Request *ComplexType.RequestType `xml:"Request,omitempty"`

	VoidShipment *VoidShipment `xml:"VoidShipment,omitempty"`
}

type VoidShipment struct {
	ShipmentIdentificationNumber string `xml:"ShipmentIdentificationNumber,omitempty"`

	TrackingNumber []string `xml:"TrackingNumber,omitempty"`
}