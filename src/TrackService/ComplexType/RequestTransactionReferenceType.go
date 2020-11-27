package ComplexType

type RequestTransactionReferenceType struct {
	CustomerContext string `xml:"CustomerContext,omitempty"`

	TransactionIdentifierPlayback string `xml:"TransactionIdentifierPlayback,omitempty"`
}
