package events

import (
	"context"
	"fmt"
	"whats/db"

	"github.com/robfig/cron"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

func (s *QuestionsService) sendWeeklyMessage() {
	var users []db.Client

	err := s.DB.Find(&users).Error
	if err != nil {
		fmt.Println("Erro ao buscar usuários:", err)
		return
	}

	for _, user := range users {
		message := "Lembrete semanal: Confira nossas novidades e promoções!"
		chatID := types.NewJID(user.Cellphone, types.DefaultUserServer)

		_, err := s.Client.SendMessage(
			context.Background(),
			chatID,
			&proto.Message{
				Conversation: &message,
			},
			whatsmeow.SendRequestExtra{},
		)
		if err != nil {
			fmt.Printf("Erro ao enviar mensagem para %s: %s\n", user.Cellphone, err)
		} else {
			fmt.Printf("Mensagem enviada para %s\n", user.Cellphone)
		}
	}
}

func (s *QuestionsService) StartWeeklyMessageScheduler() {
	c := cron.New()
	c.AddFunc("0 9 * * 1", s.sendWeeklyMessage)
	c.Start()
}
