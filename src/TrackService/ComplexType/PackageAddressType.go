package ComplexType

import "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"

type PackageAddressType struct {
	Type *CommonCodeDescriptionType `xml:"Type,omitempty"`

	Address *ComplexType.AddressType `xml:"Address,omitempty"`
}
