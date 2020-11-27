package ComplexType

import "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"

type ShipmentActivityType struct {
	ActivityLocation *ComplexType.AddressType `xml:"ActivityLocation,omitempty"`

	Description string `xml:"Description,omitempty"`

	Date string `xml:"Date,omitempty"`

	Time string `xml:"Time,omitempty"`

	Trailer string `xml:"Trailer,omitempty"`
}
