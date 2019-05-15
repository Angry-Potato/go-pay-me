package schema

// {
// 	"amount": "100.21",
// 	"currency": "GBP",
// 	"end_to_end_reference": "Wil piano Jan",
// 	"numeric_reference": "1002001",
// 	"payment_id": "123456789012345678",
// 	"payment_purpose": "Paying for goods/services",
// 	"payment_scheme": "FPS",
// 	"payment_type": "Credit",
// 	"processing_date": "2017-01-18",
// 	"reference": "Payment for Em's piano lessons",
// 	"scheme_payment_sub_type": "InternetBanking",
// 	"scheme_payment_type": "ImmediatePayment",

// 	"beneficiary_party": {
// 	  "$ref": "#/definitions/Party/example"
// 	},
// 	"charges_information": {
// 	  "$ref": "#/definitions/Charges/example"
// 	},
// 	"debtor_party": {
// 	  "$ref": "#/definitions/Party/example"
// 	},
// 	"fx": {
// 	  "$ref": "#/definitions/CurrencyExchange/example"
// 	},
// 	"sponsor_party": {
// 	  "$ref": "#/definitions/Party/example"
// 	}
//   }

var paymentTypes = []string{"Credit"}
var schemePaymentSubTypes = []string{"InternetBanking"}
var schemePaymentTypes = []string{"ImmediatePayment"}

// PaymentAttributes resource
type PaymentAttributes struct {
	ID                   uint   `gorm:"primary_key" json:"id"`
	Amount               string `json:"amount"`
	Currency             string `json:"currency"`
	EndToEndReference    string `json:"end_to_end_reference"`
	NumericReference     string `json:"numeric_reference"`
	PaymentID            string `json:"payment_id"`
	PaymentPurpose       string `json:"payment_purpose"`
	PaymentScheme        string `json:"payment_scheme"`
	PaymentType          string `json:"payment_type"`
	ProcessingDate       string `json:"processing_date"`
	Reference            string `json:"reference"`
	SchemePaymentSubType string `json:"scheme_payment_sub_type"`
	SchemePaymentType    string `json:"scheme_payment_type"`
	InternalPaymentID    string `gorm:"unique;not null" json:"internal_payment_id"`
}
