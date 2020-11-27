package SimpleType



const C_PRINT_AND_MAIL = "01" // UPS prints the label and mails the label to the customer.
const C_ONE_ATTEMPT    = "02" // UPS driver makes 1 attempt to bring the package label to the pickup location and pickup the package.
const C_THREE_ATTEMPT  = "03" // UPS driver makes 3 attempt to bring the package label to the pickup location and pickup the package.
const C_ELECTRONIC     = "04" // UPS electronically notifiesx the custoemr via e-mail that a label and receipt are available.
const C_PRINT          = "05" // The shipper prints the label and the Import Control Customer Receipt and includes w/ outbound shipment.