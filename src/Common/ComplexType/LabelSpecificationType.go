package ComplexType


type LabelSpecificationType struct {

	HTTPUserAgent string `xml:"HTTPUserAgent,omitempty"`

	LabelImageFormat *LabelImageFormatType `xml:"LabelImageFormat,omitempty"`

	LabelStockSize *LabelStockSizeType `xml:"LabelStockSize,omitempty"`


	Instruction []*InstructionCodeDescriptionType `xml:"Instruction,omitempty"`

	CharacterSet string `xml:"CharacterSet,omitempty"`
}
