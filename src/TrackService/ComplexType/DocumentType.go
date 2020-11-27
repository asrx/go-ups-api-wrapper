package ComplexType

type DocumentType struct {
	Type *CommonCodeDescriptionType `xml:"Type,omitempty"`

	Content string `xml:"Content,omitempty"`

	Format *CommonCodeDescriptionType `xml:"Format,omitempty"`
}
