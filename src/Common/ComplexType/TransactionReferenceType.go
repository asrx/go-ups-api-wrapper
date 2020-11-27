package ComplexType

type TransactionReferenceType struct {
	CustomerContext string `xml:"CustomerContext,omitempty"`

	TransactionIdentifier string `xml:"TransactionIdentifier,omitempty"`
}
