package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Repo *gorm.DB

func ConnectDatabaseGorm() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	if user == "" || password == "" || host == "" || port == "" || dbName == "" {
		log.Fatal("Variáveis de ambiente do banco de dados não configuradas corretamente")
	}

	// DSN para PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbName, password)

	engine, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados: ", err)
	}

	sqlDB, err := engine.DB()
	if err != nil {
		log.Fatal("Falha ao obter a conexão do banco de dados: ", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Falha ao pingar o banco de dados: ", err)
	}

	Repo = engine

	fmt.Println("Conexão com o banco de dados PostgreSQL estabelecida com sucesso")

	return Repo
}

func Migrate(engine *gorm.DB) error {
	tables := []interface{}{
		&Client{},
		&DataClient{},
		&Balance{},
		&Investment{},
		&Slip{},
		&Recharge{},
		&Navegation{},
	}

	for _, table := range tables {
		if err := engine.AutoMigrate(table); err != nil {
			return fmt.Errorf("falha ao migrar a tabela %T: %v", table, err)
		}
	}
	return nil
}
