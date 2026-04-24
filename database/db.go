package database

import (
	"log"
	"time"
	"os"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectaComBancoDeDados() {
	host:= os.Getenv("DB_HOST")
	user:= os.Getenv("DB_USER")
	password:= os.Getenv("DB_PASSWORD")
	dbname:= os.Getenv("DB_NAME")
	port:= os.Getenv("DB_PORT")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"

	var err error

	for i := 1; i <= 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Banco conectado com sucesso!")
			break
		}

		log.Printf("Tentativa %d falhou. Aguardando banco subir...\n", i)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	_ = DB.AutoMigrate(&models.Aluno{})
}