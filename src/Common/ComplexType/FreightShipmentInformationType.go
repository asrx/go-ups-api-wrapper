package ComplexType

type FreightShipmentInformationType struct {
	FreightDensityInfo *FreightDensityInfoType `xml:"FreightDensityInfo,omitempty"`

	DensityEligibleIndicator string `xml:"DensityEligibleIndicator,omitempty"`
}

type FreightDensityInfoType struct {
	AdjustedHeightIndicator string `xml:"AdjustedHeightIndicator,omitempty"`

	AdjustedHeight *AdjustedHeightType `xml:"AdjustedHeight,omitempty"`

	HandlingUnits []*HandlingUnitsType `xml:"HandlingUnits,omitempty"`
}

type AdjustedHeightType struct {
	Value string `xml:"Value,omitempty"`

	UnitOfMeasurement *CodeDescriptionType `xml:"UnitOfMeasurement,omitempty"`
}

type HandlingUnitsType struct {
	Quantity string `xml:"Quantity,omitempty"`

	Type *CodeDescriptionType `xml:"Type,omitempty"`

	Dimensions *HandlingUnitsDimensionsType `xml:"Dimensions,omitempty"`
}

type HandlingUnitsDimensionsType struct {
	UnitOfMeasurement *CodeDescriptionType `xml:"UnitOfMeasurement,omitempty"`

	Length string `xml:"Length,omitempty"`

	Width string `xml:"Width,omitempty"`

	Height string `xml:"Height,omitempty"`
}

type HandlingUnitsResponseType struct {
	Quantity string `xml:"Quantity,omitempty"`

	Type *CodeDescriptionType `xml:"Type,omitempty"`

	Dimensions *HandlingUnitsDimensionsType `xml:"Dimensions,omitempty"`

	AdjustedHeight *AdjustedHeightType `xml:"AdjustedHeight,omitempty"`
}