package mocks

import (
	"fmt"
	"math/rand"
	"time"
)

type InvestmentType struct {
	Name       string
	Rendimento float64
}

var investmentTypes = []InvestmentType{
	{"Renda Fixa", 0.003},             // 0.3%
	{"Renda Variável", 0.005},         // 0.5%
	{"Tesouro Direto", 0.002},         // 0.2%
	{"CDB PagBank", 0.004},            // 0.4%
	{"Fundos de Investimento", 0.004}, // 0.4%
}

func generateBoletoValue() float64 {
	return roundToTwoDecimals(50 + rand.Float64()*(1000-50))
}

func generateInvestment() (investment float64, rendimento float64, tipo InvestmentType) {
	tipo = investmentTypes[rand.Intn(len(investmentTypes))]
	investment = roundToTwoDecimals(50 + rand.Float64()*(1000-50))
	rendimento = roundToTwoDecimals(investment * tipo.Rendimento)
	return investment, rendimento, tipo
}

func generateAccountBalance() float64 {
	return roundToTwoDecimals(100 + rand.Float64()*(50000-100))
}

func roundToTwoDecimals(value float64) float64 {
	return float64(int(value*100)) / 100
}

func generateBarcode(length int) string {
	rand.Seed(time.Now().UnixNano())
	barcode := ""
	for i := 0; i < length; i++ {
		barcode += fmt.Sprintf("%d", rand.Intn(10))
	}
	return barcode
}
