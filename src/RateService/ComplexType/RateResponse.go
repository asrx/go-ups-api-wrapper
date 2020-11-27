package ComplexType

import (
	"github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	"encoding/xml"
)

type RateResponse struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/Rate/v1.1 RateResponse"`

	Response *ComplexType.ResponseType `xml:"Response,omitempty"`

	RatedShipment []*RatedShipmentType `xml:"RatedShipment,omitempty"`
}
