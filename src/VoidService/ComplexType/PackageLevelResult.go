package ComplexType

import (
	. "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
)

type PackageLevelResult struct {
	TrackingNumber string `xml:"TrackingNumber,omitempty"`

	Status *CodeDescriptionType `xml:"Status,omitempty"`
}
