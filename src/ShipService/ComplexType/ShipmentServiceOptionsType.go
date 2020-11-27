package ComplexType

import (
	. "github.com/asrx/go-ups-api-wrapper/src/Common/ComplexType"
)

type ShipmentServiceOptionsType struct {
	SaturdayDeliveryIndicator string `xml:"SaturdayDeliveryIndicator,omitempty"`

	SaturdayPickupIndicator string `xml:"SaturdayPickupIndicator,omitempty"`

	COD *CODType `xml:"COD,omitempty"`

	AccessPointCOD *PackageServiceOptionsAccessPointCODType `xml:"AccessPointCOD,omitempty"`

	DeliverToAddresseeOnlyIndicator string `xml:"DeliverToAddresseeOnlyIndicator,omitempty"`

	DirectDeliveryOnlyIndicator string `xml:"DirectDeliveryOnlyIndicator,omitempty"`

	Notification []*NotificationType `xml:"Notification,omitempty"`

	LabelDelivery *LabelDeliveryType `xml:"LabelDelivery,omitempty"`

	InternationalForms *InternationalFormType `xml:"InternationalForms,omitempty"`

	DeliveryConfirmation *DeliveryConfirmationType `xml:"DeliveryConfirmation,omitempty"`

	ReturnOfDocumentIndicator string `xml:"ReturnOfDocumentIndicator,omitempty"`

	ImportControlIndicator string `xml:"ImportControlIndicator,omitempty"`

	LabelMethod *LabelMethodType `xml:"LabelMethod,omitempty"`

	CommercialInvoiceRemovalIndicator string `xml:"CommercialInvoiceRemovalIndicator,omitempty"`

	UPScarbonneutralIndicator string `xml:"UPScarbonneutralIndicator,omitempty"`

	PreAlertNotification []*PreAlertNotificationType `xml:"PreAlertNotification,omitempty"`

	ExchangeForwardIndicator string `xml:"ExchangeForwardIndicator,omitempty"`

	HoldForPickupIndicator string `xml:"HoldForPickupIndicator,omitempty"`

	DropoffAtUPSFacilityIndicator string `xml:"DropoffAtUPSFacilityIndicator,omitempty"`

	LiftGateForPickUpIndicator string `xml:"LiftGateForPickUpIndicator,omitempty"`

	LiftGateForDeliveryIndicator string `xml:"LiftGateForDeliveryIndicator,omitempty"`

	SDLShipmentIndicator string `xml:"SDLShipmentIndicator,omitempty"`

	EPRAReleaseCode string `xml:"EPRAReleaseCode,omitempty"`

	RestrictedArticles *RestrictedArticlesType `xml:"RestrictedArticles,omitempty"`

	InsideDelivery string `xml:"InsideDelivery,omitempty"`

	ItemDisposal string `xml:"ItemDisposal,omitempty"`
}