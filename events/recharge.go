package events

import (
	"fmt"
	"strconv"
	"strings"
	"whats/db"
	"whats/mocks"

	"gorm.io/gorm"
)

func Recharge(nav *db.Navegation, messageText string, s *gorm.DB) (string, error) {
	var resposta string
	var recharge db.Recharge

	if err := s.Where("client_id = ? AND recharge_value = 0", nav.ClientID).FirstOrInit(&recharge).Error; err != nil {
		return "", fmt.Errorf("erro ao carregar recarga: %w", err)
	}

	recharge.ClientID = nav.ClientID

	switch nav.Recharge {
	case 1:
		resposta = mocks.RechargePrompt
		nav.Recharge++

		if err := s.Model(&db.Navegation{}).Where("id = ?", nav.ID).Update("recharge", nav.Recharge).Error; err != nil {
			return "", fmt.Errorf("erro ao salvar estado de navegação: %w", err)
		}

	case 2:
		number := strings.ReplaceAll(messageText, " ", "")
		if len(number) == 10 {
			recharge.RechargeType = "Bilhete Único"
		} else if len(number) == 11 {
			recharge.RechargeType = "Celular"
		} else {
			resposta = mocks.InvalidNumber
			return resposta, nil
		}

		recharge.RechargeNumber = number
		resposta = fmt.Sprintf(mocks.RechargeValuePrompt, recharge.RechargeType)
		nav.Recharge++

		if err := s.Save(&recharge).Error; err != nil {
			return "", fmt.Errorf("erro ao salvar recarga temporária: %w", err)
		}
		if err := s.Model(&db.Navegation{}).Where("id = ?", nav.ID).Update("recharge", nav.Recharge).Error; err != nil {
			return "", fmt.Errorf("erro ao salvar estado de navegação: %w", err)
		}

	case 3:
		rechargeValue, err := strconv.ParseFloat(strings.ReplaceAll(messageText, "R$", ""), 64)
		if err != nil || rechargeValue <= 0 {
			resposta = mocks.InvalidRechargeValue
			return resposta, nil
		}

		recharge.RechargeValue = rechargeValue
		resposta = fmt.Sprintf(mocks.RechargeConfirmation, recharge.RechargeValue, recharge.RechargeType, recharge.RechargeNumber)
		nav.Recharge++

		if err := s.Save(&recharge).Error; err != nil {
			return "", fmt.Errorf("erro ao salvar valor de recarga: %w", err)
		}
		if err := s.Model(&db.Navegation{}).Where("id = ?", nav.ID).Update("recharge", nav.Recharge).Error; err != nil {
			return "", fmt.Errorf("erro ao salvar estado de navegação: %w", err)
		}

	case 4:
		if err := s.Where("client_id = ? AND recharge_value > 0", nav.ClientID).First(&recharge).Error; err != nil {
			return "", fmt.Errorf("erro ao carregar recarga confirmada: %w", err)
		}

		if strings.EqualFold(messageText, "Sim") {
			resposta = fmt.Sprintf(mocks.RechargeSuccess, recharge.RechargeValue, recharge.RechargeType, recharge.RechargeNumber)
		} else {
			resposta = mocks.RechargeCancellation
		}
		nav.Recharge = 1

		if err := s.Model(&db.Navegation{}).Where("id = ?", nav.ID).Update("recharge", nav.Recharge).Error; err != nil {
			return "", fmt.Errorf("erro ao resetar estado de navegação: %w", err)
		}
	}

	return resposta, nil
}
