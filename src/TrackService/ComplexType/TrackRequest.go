package ComplexType

import (
	"github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
	"encoding/xml"
)

type TrackRequest struct {
	XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/Track/v2.0 TrackRequest"`

	Request *ComplexType.RequestType `xml:"Request,omitempty"`

	InquiryNumber string `xml:"InquiryNumber,omitempty"`

	IncludeMailInnovationIndicator string `xml:"IncludeMailInnovationIndicator,omitempty"`

	TrackingOption string `xml:"TrackingOption,omitempty"`

	CandidateBookmark string `xml:"CandidateBookmark,omitempty"`

	ShipperAccountInfo *ShipperAccountInfoType `xml:"ShipperAccountInfo,omitempty"`

	PreauthorizedReturnIndicator string `xml:"PreauthorizedReturnIndicator,omitempty"`

	Locale string `xml:"Locale,omitempty"`
}

