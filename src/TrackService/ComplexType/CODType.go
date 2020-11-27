package ComplexType

type CODType struct {
	Amount *AmountType `xml:"Amount,omitempty"`

	Status *CODStatusType `xml:"Status,omitempty"`

	ControlNumber string `xml:"ControlNumber,omitempty"`
}
