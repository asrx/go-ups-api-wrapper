package ComplexType

import (
	"encoding/xml"
	"github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
)

type XAVResponse struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/xav/v1.0 XAVResponse"`

	Response *ComplexType.ResponseType `xml:"Response,omitempty"`

	AddressClassification *AddressClassificationType `xml:"AddressClassification,omitempty"`

	Candidate []*CandidateType `xml:"Candidate,omitempty"`

	ValidAddressIndicator string `xml:"ValidAddressIndicator,omitempty"`

	AmbiguousAddressIndicator string `xml:"AmbiguousAddressIndicator,omitempty"`

	NoCandidatesIndicator string `xml:"NoCandidatesIndicator,omitempty"`
}
