package ComplexType


type ReceiptSpecificationType struct {
	ImageFormat *ReceiptImageFormatType `xml:"ImageFormat,omitempty"`
}


type ReceiptImageFormatType struct {
	Code string `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`
}