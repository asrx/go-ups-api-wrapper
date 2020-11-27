package ComplexType

import "encoding/xml"

type ErrorDetail struct {
	XMLName xml.Name `xml:"detail,omitempty"`
	Errors *Errors `xml:"http://www.ups.com/XMLSchema/XOLTWS/Error/v1.1 Errors",omitempty`
}

type Errors struct {
	//XMLName xml.Name `xml:"http://www.ups.com/XMLSchema/XOLTWS/Error/v1.1 Errors"`

	ErrorDetail []*ErrorDetailType `xml:"ErrorDetail,omitempty"`
}

type ErrorDetailType struct {
	Severity string `xml:"Severity,omitempty"`

	PrimaryErrorCode *CodeType `xml:"PrimaryErrorCode,omitempty"`

	//MinimumRetrySeconds string `xml:"MinimumRetrySeconds,omitempty"`
	//
	//Location *LocationType `xml:"Location,omitempty"`
	//
	//SubErrorCode []*CodeType `xml:"SubErrorCode,omitempty"`
	//
	//AdditionalInformation []*AdditionalInfoType `xml:"AdditionalInformation,omitempty"`
	//
	//ElementLevelInformation *ElementLevelInformationType `xml:"ElementLevelInformation,omitempty"`
}


type CodeType struct {
	Code string `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`

	Digest string `xml:"Digest,omitempty"`
}


type LocationType struct {
	LocationElementName string `xml:"LocationElementName,omitempty"`

	XPathOfElement string `xml:"XPathOfElement,omitempty"`

	OriginalValue string `xml:"OriginalValue,omitempty"`
}


type AdditionalInfoType struct {
	Type string `xml:"Type,omitempty"`

	Value []*AdditionalCodeDescType `xml:"Value,omitempty"`
}

type AdditionalCodeDescType struct {
	Code string `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`
}


type ElementLevelInformationType struct {
	Level string `xml:"Level,omitempty"`

	ElementIdentifier []*ElementIdentifierType `xml:"ElementIdentifier,omitempty"`
}


type ElementIdentifierType struct {
	Code string `xml:"Code,omitempty"`

	Value string `xml:"Value,omitempty"`
}

type DetailType struct {
	Code string `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`

	ElementLevelInformation *ElementLevelInformationType `xml:"ElementLevelInformation,omitempty"`
}

