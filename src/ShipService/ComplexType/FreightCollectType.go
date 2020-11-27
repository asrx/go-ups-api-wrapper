package ComplexType

type FreightCollectType struct {
	BillReceiver *BillReceiverType `xml:"BillReceiver,omitempty"`
}