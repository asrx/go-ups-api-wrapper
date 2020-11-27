package ComplexType

type FormImageType struct {
	ImageFormat *ImageFormatType `xml:"ImageFormat,omitempty"`

	GraphicImage string `xml:"GraphicImage,omitempty"`
}
