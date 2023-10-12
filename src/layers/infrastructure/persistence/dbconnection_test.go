package persistence

import (
	"gorm.io/gorm"
	"server/src/layers/domain/models"
	"testing"
)

func TestConnect(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Verificar se a tabela do usuário foi criada
	if !db.Migrator().HasTable(&models.User{}) {
		t.Fatal("A tabela do usuário não foi criada")
	}

	// Fechar conexão ao finalizar o teste
	gDB := db.Session(&gorm.Session{DryRun: true})
	sqlDB, err := gDB.DB()
	if err != nil {
		t.Fatalf("Erro ao obter *sql.DB: %v", err)
	}

	defer Close(sqlDB)
}

func TestClose(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Erro ao obter *sql.DB: %v", err)
	}

	Close(sqlDB)

	// Tentar usar a conexão após fechá-la
	err = db.Exec("SELECT 1").Error
	if err == nil {
		t.Fatal("A conexão com o banco de dados ainda está aberta")
	}
}
