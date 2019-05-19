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
		SponsorParty:         ValidParty(),
		ForeignExchange:      ValidCurrencyExchange(),
		ChargesInformation:   ValidCharges(),
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

// ValidCurrencyExchange an example of a valid currency exchange
func ValidCurrencyExchange() CurrencyExchange {
	return CurrencyExchange{
		ContractReference: "FX123",
		ExchangeRate:      "2.00000",
		OriginalAmount:    "200.42",
		OriginalCurrency:  "USD",
	}
}

// ValidCharges an example of a valid currency exchange
func ValidCharges() Charges {
	moneyA, moneyB := ValidMoney(), ValidMoney()
	return Charges{
		BearerCode:              "SHAR",
		ReceiverChargesAmount:   "1.00",
		ReceiverChargesCurrency: "USD",
		SenderCharges: []*Money{
			&moneyA,
			&moneyB,
		},
	}
}

// ValidMoney an example of a valid money
func ValidMoney() Money {
	return Money{
		Amount:   "5.00",
		Currency: "USD",
	}
}
