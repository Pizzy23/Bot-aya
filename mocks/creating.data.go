package mocks

import (
	"whats/db"

	"golang.org/x/exp/rand"
)

func CreateInvestments(clientID uint) []db.Investment {
	count := rand.Intn(20) + 1
	investments := make([]db.Investment, count)

	for i := 0; i < count; i++ {
		investments[i] = CreateInvestment(clientID)
	}

	return investments
}

func CreateSlips(clientID uint) []db.Slip {
	count := rand.Intn(20) + 1
	slips := make([]db.Slip, count)

	for i := 0; i < count; i++ {
		slips[i] = CreateSlip(clientID)
	}

	return slips
}

func CreateInvestment(clientID uint) db.Investment {
	value, income, invType := generateInvestment()
	name := generateInvestmentPlaceName()

	return db.Investment{
		ClientID:     clientID,
		Name:         name,
		Value:        value,
		Income:       income,
		TypeOfInvest: invType.Name,
	}
}

func CreateSlip(clientID uint) db.Slip {
	value := generateBoletoValue()
	name := generatePlaceName()
	barcode := generateBarcode(48)
	return db.Slip{
		ClientID: clientID,
		Name:     name,
		Value:    value,
		BarCode:  barcode,
	}
}

func GenerateBalance(clientID uint) db.Balance {
	return db.Balance{
		ClientID:     clientID,
		Balance:      generateAccountBalance(),
		TotalPayment: roundToTwoDecimals(1000 + rand.Float64()*(30000-1000)), // TotalPayment fictício, ajuste conforme necessário
	}
}
