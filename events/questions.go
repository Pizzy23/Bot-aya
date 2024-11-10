package events

import (
	"context"
	"fmt"
	"os"
	"strings"
	"whats/db"
	"whats/mocks"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"

	"gorm.io/gorm"
)

type QuestionsService struct {
	DB     *gorm.DB
	Client *whatsmeow.Client
}

const forceSendCode = "SemanalSend"

func getMessageText(msg *proto.Message) string {
	if msg.Conversation != nil {
		return *msg.Conversation
	}
	if msg.ExtendedTextMessage != nil {
		return *msg.ExtendedTextMessage.Text
	}
	return ""
}

func (s *QuestionsService) EventHandler(client *whatsmeow.Client, evt interface{}) {
	var user db.Client
	var nav db.Navegation
	var resposta string

	switch v := evt.(type) {
	case *events.Message:
		messageText := getMessageText(v.Message)
		fmt.Printf("Mensagem recebida: %s\n", messageText)
		admin := os.Getenv("ADMIN")
		if messageText == forceSendCode && v.Info.Chat.User == admin {
			fmt.Println(mocks.WeeklySendCodeReceived)
			s.sendWeeklyMessage()
			return
		}

		if messageText != "" && !v.Info.IsFromMe {
			userWhats := v.Info.Chat.User
			userDb, err := db.GetByCell(s.DB, &user, map[string]interface{}{"cellphone": userWhats})
			if err != nil {
				fmt.Printf("Erro ao pegar o user: %s\n", err)
				return
			}

			if userDb {
				found, err := db.GetByID(s.DB, &nav, int64(user.ID))
				if err != nil {
					fmt.Printf("Erro ao pegar a navegação: %s\n", err)
					return
				}

				if !found {
					nav = db.Navegation{
						ClientID:  user.ID,
						Payment:   1,
						Recharge:  1,
						Invest:    1,
						Treatment: false,
					}
					err = db.Create(s.DB, &nav)
					if err != nil {
						fmt.Printf("Erro ao criar a navegação inicial: %s\n", err)
						return
					}
				}

				if nav.Treatment && nav.Payment == 1 && nav.Recharge == 1 && nav.Invest == 1 {
					resposta = mocks.WelcomeMessage
					nav.Treatment = false
					if err := s.DB.Model(&db.Navegation{}).Where("id = ?", nav.ID).Updates(map[string]interface{}{
						"treatment": nav.Treatment,
						"payment":   1,
						"recharge":  1,
						"invest":    1,
					}).Error; err != nil {
						fmt.Printf("Erro ao salvar estado de navegação: %s", err)
					}
				} else {
					switch {
					case strings.Contains(strings.ToLower(messageText), "agenda integrada") || strings.Contains(strings.ToLower(messageText), "agenda") || nav.Payment >= 2:
						resposta, err = Slipers(nav, messageText, s.DB)
					case strings.Contains(strings.ToLower(messageText), "investimento") || nav.Invest >= 2:
						resposta, err = InvestSummary(&nav, messageText, s.DB)
					case strings.Contains(strings.ToLower(messageText), "recargas") || nav.Recharge >= 2:
						resposta, err = Recharge(&nav, messageText, s.DB)
					}

					if err != nil {
						fmt.Printf("Erro na navegação: %s\n", err)
					}

					if nav.Payment == 1 && nav.Recharge == 1 && nav.Invest == 1 {
						nav.Treatment = true
						if err := s.DB.Model(&db.Navegation{}).Where("id = ?", nav.ID).Update("treatment", true).Error; err != nil {
							fmt.Printf("Erro ao salvar estado de navegação: %s", err)
						}
					}
				}

				if resposta == "" {
					resposta = mocks.UnrecognizedCommand
				}

				_, err = client.SendMessage(
					context.Background(),
					v.Info.Chat,
					&proto.Message{
						Conversation: &resposta,
					},
					whatsmeow.SendRequestExtra{},
				)

				if err != nil {
					fmt.Println("Erro ao enviar a mensagem:", err)
				}
				return
			}

			newUser := db.Client{Cellphone: userWhats}
			if err := db.Create(s.DB, &newUser); err != nil {
				fmt.Printf("Erro ao salvar o usuário: %s\n", err)
				return
			}

			data := db.DataClient{
				ClientID:    newUser.ID,
				Investments: mocks.CreateInvestments(newUser.ID),
				Slips:       mocks.CreateSlips(newUser.ID),
				Balances:    mocks.GenerateBalance(newUser.ID),
			}
			if err := db.Create(s.DB, &data); err != nil {
				fmt.Printf("Erro ao salvar o Data: %s\n", err)
				return
			}

			resposta = mocks.WelcomeMessage
			_, err = client.SendMessage(
				context.Background(),
				v.Info.Chat,
				&proto.Message{
					Conversation: &resposta,
				},
				whatsmeow.SendRequestExtra{},
			)
			if err != nil {
				fmt.Println("Erro ao enviar a mensagem:", err)
			}
		}
	}
}
