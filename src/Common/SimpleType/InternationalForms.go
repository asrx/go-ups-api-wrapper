package SimpleType


const TYPE_INVOICE = "01"
const TYPE_CO = "03"
const TYPE_NAFTA_CO = "04"
const TYPE_PARTIAL_INVOICE = "05"
const TYPE_PACKINGLIST = "06"
const TYPE_CUSTOMER_GENERATED_FORMS = "07"
const TYPE_AIR_FREIGHT_PACKING_LIST = "08"
const TYPE_CN22_FORMS = "09"
const TYPE_UPS_PREMIUM_CARE = "10"
const TYPE_EEI_SHIPMENT_WITH_RETURN_SERVICE = "11"

/*
		"01' => 'Invoice'
        '03' => 'CO',
        '04' => 'NAFTA CO',
        '05' => 'Partial Invoice',
        '06' => 'Packinglist',
        '07' => 'Customer Generated Forms',
        '08' => 'Air Freight Packing List',
        '09' => 'CN22 Forms',
        '10' => 'UPS Premium Care',
        '11' => 'EEI. For shipment with return service',
*/

const TOS_COST_AND_FREIGHT = "CFR"
const TOS_COST_INSURANCE_AND_FREIGHT = "CIF"
const TOS_CARRIAGE_AND_INSURANCE_PAID = "CIP"
const TOS_CARRIAGE_PAID_TO = "CPT"
const TOS_DELIVERED_AT_FRONTIER = "DAF"
const TOS_DELIVERY_DUTY_PAID = "DDP"
const TOS_DELIVERY_DUTY_UNPAID = "DDU"
const TOS_DELIVERED_EX_QUAY = "DEQ"
const TOS_DELIVERED_EX_SHIP = "DES"
const TOS_EX_WORKS = "EXW"
const TOS_FREE_ALONGSIDE_SHIP = "FAS"
const TOS_FREE_CARRIER = "FCA"
const TOS_FREE_ON_BOARD = "FOB"

/*
		'CFR' => 'Cost and Freight',
        'CIF' => 'Cost, Insurance and Freight',
        'CIP' => 'Carriage and Insurance Paid',
        'CPT' => 'Carriage Paid To',
        'DAF' => 'Delivered at Frontier',
        'DDP' => 'Delivery Duty Paid',
        'DDU' => 'Delivery Duty Unpaid',
        'DEQ' => 'Delivered Ex Quay',
        'DES' => 'Delivered Ex Ship',
        'EXW' => 'Ex Works',
        'FAS' => 'Free Alongside Ship',
        'FCA' => 'Free Carrier',
        'FOB' => 'Free On Board',
*/

const RFE_SALE = "SALE"
const RFE_GIFT = "GIFT"
const RFE_SAMPLE = "SAMPLE"
const RFE_RETURN = "RETURN"
const RFE_REPAIR = "REPAIR"
const RFE_INTERCOMPANYDATA = "INTERCOMPANYDATA"