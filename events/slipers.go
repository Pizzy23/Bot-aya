package events

import (
	"fmt"
	"strconv"
	"whats/db"
	"whats/mocks"

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
			resposta = mocks.SlipsIntro
			for i, slip := range slips {
				resposta += fmt.Sprintf("%d - %s\n", i+1, slip.Name)
			}
			resposta += mocks.SlipsPrompt
		} else {
			resposta = mocks.SlipsNoPending
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
			resposta = mocks.InvalidSlipNumber
			return resposta, nil
		}

		slips, err := db.GetPendingSlips(s, nav.ClientID)
		if err != nil || index > len(slips) {
			resposta = mocks.SlipNotFound
			return resposta, nil
		}

		selectedSlip := slips[index-1]
		resposta = fmt.Sprintf(mocks.SlipDetailsTemplate, selectedSlip.Name, selectedSlip.Value, selectedSlip.BarCode)
		nav.Payment++

		if err := s.Model(&db.Navegation{}).Where("id = ?", nav.ID).Updates(map[string]interface{}{
			"payment": nav.Payment,
		}).Error; err != nil {
			return "", fmt.Errorf("erro ao salvar estado de navegação: %w", err)
		}

	case 3:
		if messageText == "Sim" || messageText == "sim" || messageText == "SIM" {
			resposta = mocks.PaymentConfirmed
			nav.Payment = 1

			err := db.UpdateDebt(s, nav.ClientID)
			if err != nil {
				fmt.Println("Erro ao atualizar a dívida:", err)
			}
		} else {
			resposta = mocks.PaymentCancelled
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
