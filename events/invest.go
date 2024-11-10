package events

import (
	"fmt"
	"strconv"
	"whats/db"

	"gorm.io/gorm"
)

func InvestSummary(nav *db.Navegation, messageText string, s *gorm.DB) (string, error) {
	var resposta string

	switch nav.Invest {
	case 1:
		// Exibe os tipos de investimento disponíveis para escolha
		resposta = "Olá! Aqui estão os tipos de investimento disponíveis:\n" +
			"1 - Renda Fixa\n" +
			"2 - Renda Variável\n" +
			"3 - Tesouro Direto\n" +
			"4 - CDB PagBank\n" +
			"5 - Fundos de Investimento\n\n" +
			"Por favor, digite o número do investimento que deseja visualizar."

		// Atualiza o estado de navegação para o próximo passo
		nav.Invest++
		if err := s.Model(&db.Navegation{}).Where("id = ?", nav.ID).Update("invest", nav.Invest).Error; err != nil {
			return "", fmt.Errorf("erro ao salvar estado de navegação: %w", err)
		}

	case 2:
		index, err := strconv.Atoi(messageText)
		if err != nil || index < 1 || index > 5 {
			resposta = "Número inválido. Por favor, tente novamente digitando o número do tipo de investimento."
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
			resposta = "Fico feliz em ajudar! Caso precise de atualizações ou qualquer outra assistência, é só chamar. Boa semana! 📈"
			nav.Invest = 1
		} else {
			index, err := strconv.Atoi(messageText)
			if err != nil || index < 1 || index > 5 {
				resposta = "Número inválido. Por favor, tente novamente digitando o número do investimento específico. \n \n Digite 'não' para sair"
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
		resposta = fmt.Sprintf("Nenhum investimento encontrado para o tipo: %s.", investmentType)
	} else {
		totalValue := 0.0
		for _, inv := range investments {
			totalValue += inv.Value
		}

		resposta = fmt.Sprintf("Resumo do investimento em %s:\n\nTotal Investido: R$%.2f\nRendimentos: +R$%.2f\n", investmentType, totalValue, totalValue*0.015)
		resposta += "Deseja mais detalhes sobre algum investimento específico? (Digite o número ou 'não' para encerrar.)"
	}

	return resposta, nil
}
