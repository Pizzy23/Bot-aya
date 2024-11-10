package events

import (
	"fmt"
	"strconv"
	"whats/db"
	"whats/mocks"

	"gorm.io/gorm"
)

func InvestSummary(nav *db.Navegation, messageText string, s *gorm.DB) (string, error) {
	var resposta string

	switch nav.Invest {
	case 1:
		resposta = mocks.InvestIntro
		nav.Invest++
		if err := s.Model(&db.Navegation{}).Where("id = ?", nav.ID).Update("invest", nav.Invest).Error; err != nil {
			return "", fmt.Errorf("erro ao salvar estado de navegação: %w", err)
		}

	case 2:
		index, err := strconv.Atoi(messageText)
		if err != nil || index < 1 || index > 5 {
			resposta = mocks.InvestInvalidOption
			return resposta, nil
		}

		resposta, err = invest(nav, index, s)
		if err != nil {
			return "", err
		}

		nav.Invest++
		if err := s.Model(&db.Navegation{}).Where("id = ?", nav.ID).Update("invest", nav.Invest).Error; err != nil {
			return "", fmt.Errorf("erro ao salvar estado de navegação: %w", err)
		}

	case 3:
		if messageText == "não" || messageText == "Não" || messageText == "NÃO" ||
			messageText == "nao" || messageText == "Nao" || messageText == "NAO" {
			resposta = mocks.InvestExitPrompt
			nav.Invest = 1
		} else {
			index, err := strconv.Atoi(messageText)
			if err != nil || index < 1 || index > 5 {
				resposta = mocks.InvestExitOption
				return resposta, nil
			}

			resposta, err = invest(nav, index, s)
			if err != nil {
				return "", err
			}
		}

		if err := s.Model(&db.Navegation{}).Where("id = ?", nav.ID).Update("invest", nav.Invest).Error; err != nil {
			return "", fmt.Errorf("erro ao salvar estado de navegação: %w", err)
		}
	}

	return resposta, nil
}

func invest(nav *db.Navegation, index int, s *gorm.DB) (string, error) {
	var resposta string
	var investments []db.Investment
	var investmentType string

	switch index {
	case 1:
		investmentType = "Renda Fixa"
	case 2:
		investmentType = "Renda Variável"
	case 3:
		investmentType = "Tesouro Direto"
	case 4:
		investmentType = "CDB PagBank"
	case 5:
		investmentType = "Fundos de Investimento"
	}

	if err := s.Where("client_id = ? AND type_of_invest = ?", nav.ClientID, investmentType).Find(&investments).Error; err != nil {
		return "", fmt.Errorf("erro ao buscar investimentos: %w", err)
	}

	if len(investments) == 0 {
		resposta = fmt.Sprintf(mocks.InvestNotFound, investmentType)
	} else {
		totalValue := 0.0
		for _, inv := range investments {
			totalValue += inv.Value
		}

		resposta = fmt.Sprintf(mocks.InvestSummaryFormat, investmentType, totalValue, totalValue*0.015)
		resposta += mocks.InvestDetailsPrompt
	}

	return resposta, nil
}
