package ComplexType

import (
	"github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	"encoding/xml"
)

type TrackResponse struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/Track/v2.0 TrackResponse"`

	Response *ComplexType.ResponseType `xml:"Response,omitempty"`

	Shipment *ShipmentType `xml:"Shipment,omitempty"`

	Disclaimer []string `xml:"Disclaimer,omitempty"`
}

