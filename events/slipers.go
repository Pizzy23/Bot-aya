package events

import (
	"fmt"
	"strconv"
	"whats/db"

	"gorm.io/gorm"
)

func Slipers(nav db.Navegation, messageText string, s *gorm.DB) (string, error) {
	var resposta string

	switch nav.Payment {
	case 1:
		slips, err := db.GetPendingSlips(s, nav.ClientID)
		if err != nil {
			fmt.Println("Erro ao obter os boletos pendentes:", err)
			return "", err
		}

		if len(slips) > 0 {
			resposta = "Olá! Aqui estão os boletos em aberto:\n"
			for i, slip := range slips {
				resposta += fmt.Sprintf("%d - %s\n", i+1, slip.Name)
			}
			resposta += "\nPor favor, digite o número do boleto que deseja agendar para pagamento."
		} else {
			resposta = "Olá! Não há boletos pendentes no momento."
			nav.Payment = 1
		}
		nav.Payment++

		if err := s.Model(&db.Navegation{}).Where("id = ?", nav.ID).Updates(map[string]interface{}{
			"payment": nav.Payment,
		}).Error; err != nil {
			return "", fmt.Errorf("erro ao salvar estado de navegação: %w", err)
		}

	case 2:
		index, err := strconv.Atoi(messageText)
		if err != nil || index < 1 {
			resposta = "Número inválido. Por favor, tente novamente digitando o número do boleto."
			return resposta, nil
		}

		slips, err := db.GetPendingSlips(s, nav.ClientID)
		if err != nil || index > len(slips) {
			resposta = "Boleto não encontrado. Por favor, tente novamente."
			return resposta, nil
		}

		selectedSlip := slips[index-1]
		resposta = fmt.Sprintf("Obrigado! Aqui estão os detalhes do boleto:\n\nNome: %s\nValor: R$%.2f\nCódigo de Barras: %s\n\nConfirma o agendamento?", selectedSlip.Name, selectedSlip.Value, selectedSlip.BarCode)
		nav.Payment++

		if err := s.Model(&db.Navegation{}).Where("id = ?", nav.ID).Updates(map[string]interface{}{
			"payment": nav.Payment,
		}).Error; err != nil {
			return "", fmt.Errorf("erro ao salvar estado de navegação: %w", err)
		}

	case 3:
		if messageText == "Sim" {
			resposta = "Perfeito! Seu pagamento foi agendado. Vou enviar uma notificação de confirmação no dia do pagamento. Algo mais em que posso ajudar?"
			nav.Payment = 1

			err := db.UpdateDebt(s, nav.ClientID)
			if err != nil {
				fmt.Println("Erro ao atualizar a dívida:", err)
			}
		} else {
			resposta = "Cancelando agendamento. Posso ajudar com mais alguma coisa?"
			nav.Payment = 1
		}

		if err := s.Model(&db.Navegation{}).Where("id = ?", nav.ID).Updates(map[string]interface{}{
			"payment": nav.Payment,
		}).Error; err != nil {
			return "", fmt.Errorf("erro ao salvar estado de navegação: %w", err)
		}
	}

	return resposta, nil
}
