package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"whats/db"
	events "whats/events"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
	"gorm.io/gorm"
)

func main() {
	dbInstance := db.ConnectDatabaseGorm()

	migrate(dbInstance)

	dbLog := waLog.Stdout("Database", "DEBUG", true)
	container, err := sqlstore.New("sqlite3", "file:examplestore.db?_foreign_keys=on", dbLog)
	if err != nil {
		fmt.Println("Failed to open database:", err)
		return
	}

	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		fmt.Println("Failed to get first device:", err)
		return
	}

	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)

	questionsService := &events.QuestionsService{
		DB:     dbInstance,
		Client: client,
	}

	questionsService.StartWeeklyMessageScheduler()

	client.AddEventHandler(func(evt interface{}) {
		questionsService.EventHandler(client, evt)
	})

	if client.Store.ID == nil {
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			fmt.Println("Failed to connect:", err)
			return
		}
		for evt := range qrChan {
			fmt.Println("QR Code Event:", evt.Event, evt.Code)
			config := qrterminal.Config{
				Level:     qrterminal.L,
				Writer:    os.Stdout,
				BlackChar: qrterminal.BLACK,
				WhiteChar: qrterminal.WHITE,
				QuietZone: 1,
			}
			qrterminal.GenerateWithConfig(evt.Code, config)
		}
	} else {
		err = client.Connect()
		if err != nil {
			fmt.Println("Failed to reconnect:", err)
			return
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
}

func migrate(dbInstance *gorm.DB) {
	if err := db.Migrate(dbInstance); err != nil {
		log.Fatal("MIGRATION: error", err)
	}
	fmt.Println("MIGRATION: successfully migrated")
}
