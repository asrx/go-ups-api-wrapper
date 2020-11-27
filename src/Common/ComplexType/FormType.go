package ComplexType

type FormType struct {

	Image *FormImageType `xml:"Image,omitempty"`

	Code string `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`

	FormGroupId string `xml:"FormGroupId,omitempty"`

	FormGroupIdName string `xml:"FormGroupIdName,omitempty"`
}
