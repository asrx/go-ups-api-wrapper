package ComplexType

import "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"

type ActivityLocationType struct {
	Address *ComplexType.AddressType `xml:"Address,omitempty"`

	TransportFacility *TransportFacilityType `xml:"TransportFacility,omitempty"`

	Code string `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`

	SignedForByName string `xml:"SignedForByName,omitempty"`
}
