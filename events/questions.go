package events

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"whats/db"
	"whats/mocks"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
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

func (s *QuestionsService) sendWelcomeMessage(client *whatsmeow.Client, chat types.JID) {
	welcomeMessage := mocks.WelcomeMessage

	_, err := client.SendMessage(
		context.Background(),
		chat,
		&proto.Message{
			Conversation: &welcomeMessage,
		},
		whatsmeow.SendRequestExtra{},
	)
	if err != nil {
		fmt.Println("Erro ao enviar a mensagem de boas-vindas:", err)
	}
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
						ClientID: user.ID,
						Payment:  1,
						Recharge: 1,
						Invest:   1,
					}
					err = db.Create(s.DB, &nav)
					if err != nil {
						fmt.Printf("Erro ao criar a navegação inicial: %s\n", err)
						return
					}
				}

				if messageText == "Agenda Integrada" || messageText == "agenda integrada" ||
					messageText == "AGENDA INTEGRADA" || messageText == "Agenda integrada" || nav.Payment >= 2 {
					resposta, err = Slipers(nav, messageText, s.DB)
					if err != nil {
						fmt.Printf("Erro na navegação de pagamentos: %s\n", err)
					}
				}
				if messageText == "investimento" || messageText == "Investimento" || messageText == "INVESTIMENTO" || nav.Invest >= 2 {
					resposta, err = InvestSummary(&nav, messageText, s.DB)
					if err != nil {
						fmt.Printf("Erro na navegação de investimento: %s\n", err)
					}
				}
				if messageText == "recarga" || messageText == "Recarga" || messageText == "RECARGA" || nav.Recharge >= 2 {
					resposta, err = Recharge(&nav, messageText, s.DB)
					if err != nil {
						fmt.Printf("Erro na navegação de Recarga: %s\n", err)
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

			if isEmail(messageText) {
				newUser := db.Client{
					Email:     messageText,
					Cellphone: userWhats,
				}
				err := db.Create(s.DB, &newUser)
				if err != nil {
					fmt.Printf("Erro ao salvar o user: %s\n", err)
					return
				}

				invest := mocks.CreateInvestments(newUser.ID)
				slip := mocks.CreateSlips(newUser.ID)
				balance := mocks.GenerateBalance(newUser.ID)

				data := db.DataClient{
					Investments: invest,
					Slips:       slip,
					Balances:    balance,
				}

				err = db.Create(s.DB, &data)
				if err != nil {
					fmt.Printf("Erro ao salvar o Data: %s\n", err)
					return
				}

				resposta = mocks.UserCreationSuccess
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

			resposta = mocks.EmailPrompt
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

func isEmail(message string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(message)
}
