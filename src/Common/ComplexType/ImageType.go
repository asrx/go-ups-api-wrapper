package ComplexType

type ImageType struct {
	ImageFormat *ImageFormatType `xml:"ImageFormat,omitempty"`

	GraphicImage string `xml:"GraphicImage,omitempty"`
}