package ComplexType

import "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"

type SignatureImageType struct {
	GraphicImage string `xml:"GraphicImage,omitempty"`

	ImageFormat *ComplexType.ImageFormatType `xml:"ImageFormat,omitempty"`
}
