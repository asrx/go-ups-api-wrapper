package ComplexType

type ReceiptType struct {
	Image *ImageType `xml:"Image,omitempty"`

	HTMLImage string `xml:"HTMLImage,omitempty"`

	URL string `xml:"URL,omitempty"`
}

