package ComplexType


type LabelDeliveryType struct {
	EMail *EmailDetailsType `xml:"EMail,omitempty"`

	LabelLinksIndicator string `xml:"LabelLinksIndicator,omitempty"`
}

type EmailDetailsType struct {
	EMailAddress []string `xml:"EMailAddress,omitempty"`

	UndeliverableEMailAddress string `xml:"UndeliverableEMailAddress,omitempty"`

	FromEMailAddress string `xml:"FromEMailAddress,omitempty"`

	FromName string `xml:"FromName,omitempty"`

	Memo string `xml:"Memo,omitempty"`

	Subject string `xml:"Subject,omitempty"`

	SubjectCode string `xml:"SubjectCode,omitempty"`
}
