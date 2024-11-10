package mocks

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

var typesOfPlaces = []string{
	"Faculdade", "Mobilia", "Loja", "Restaurante", "Supermercado", "Farmácia",
	"Clínica", "Universidade", "Escola", "Centro Médico", "Auto Peças", "Academia",
}

var nameAdjectives = []string{
	"Alfa", "Beta", "Delta", "Gama", "Omega", "Supremo", "Prime", "Max",
	"Global", "Plus", "Premium", "Avançado", "Central", "Brasil", "Nacional",
	"Metropolitano", "Comercial", "Paulista", "Carioca", "Pro", "Integral",
}

var investmentPlaceTypes = []string{
	"Banco", "Investimentos", "Corretora", "Gestora", "Fundo", "Consultoria",
	"Capital", "Holding", "Asset Management", "Financeira", "Instituição",
}

var investmentAdjectives = []string{
	"Alfa", "Beta", "Delta", "Global", "Prime", "Elite", "Premium", "Avançado",
	"Central", "Supremo", "Pro", "Primeiro", "Integral", "Brasil", "Internacional",
	"Omega", "Master", "Fortuna", "Metropolitano", "Patrimônio", "Capital",
}

func generateInvestmentPlaceName() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	placeType := investmentPlaceTypes[rand.Intn(len(investmentPlaceTypes))]
	adjective := investmentAdjectives[rand.Intn(len(investmentAdjectives))]
	return fmt.Sprintf("%s %s", placeType, adjective)
}

func generatePlaceName() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	placeType := typesOfPlaces[rand.Intn(len(typesOfPlaces))]
	nameAdjective := nameAdjectives[rand.Intn(len(nameAdjectives))]
	return fmt.Sprintf("%s %s", placeType, nameAdjective)
}
