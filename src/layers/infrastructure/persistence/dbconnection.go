package persistence

import (
	"database/sql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"server/src/layers/domain/models"
	"time"
)

func Connect() (*gorm.DB, error) {
	config := &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	}

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), config)
	if err != nil {
		log.Printf("Erro ao conectar ao banco de dados: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Printf("Erro na migração automática: %v", err)
		return nil, err
	}

	log.Println("Banco de dados aberto com sucesso")
	return db, nil
}

func Close(db *sql.DB) {
	if db != nil {
		if err := db.Close(); err != nil {
			log.Printf("falha ao fechar a conexão com o banco de dados: %v", err)
		}
	}
}
