package schema

import "github.com/google/uuid"

// ValidPayment an example of a valid payment
func ValidPayment() *Payment {
	ID := uuid.New().String()
	return &Payment{
		ID:             ID,
		Type:           "Payment",
		Version:        0,
		OrganisationID: uuid.New().String(),
		Attributes:     ValidPaymentAttributes(),
	}
}

// ValidPaymentAttributes an example of a valid paymentattributes
func ValidPaymentAttributes() PaymentAttributes {
	return PaymentAttributes{
		Amount:               "200.10",
		Currency:             "great",
		EndToEndReference:    "here it is",
		NumericReference:     "1245",
		PaymentID:            "343535",
		PaymentPurpose:       "stuff",
		PaymentScheme:        "best",
		PaymentType:          "Credit",
		ProcessingDate:       "now",
		Reference:            "that guy",
		SchemePaymentSubType: "InternetBanking",
		SchemePaymentType:    "ImmediatePayment",
		BeneficiaryParty:     ValidParty(),
		DebtorParty:          ValidParty(),
	}
}

// ValidParty an example of a valid party
func ValidParty() Party {
	return Party{
		AccountName:       "My account",
		AccountNumber:     "12345",
		AccountNumberCode: "93847",
		AccountType:       5,
		Address:           "123 lane",
		BankID:            "best bank",
		BankIDCode:        "12AFR",
		Name:              "Mr Man",
	}
}
