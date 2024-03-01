package repositories

import (
	"github.com/luizpais/rinha-de-backend-go/rdbe/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"time"
)

var repo *Repository

func setup() {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	repo = NewRepository(db)
}

func TestGetContaCorrenteByID(t *testing.T) {
	setup()

	// Create a ContaCorrente for testing
	contaCorrente := models.ContaCorrente{
		Nome:   "Test",
		Saldo:  1000,
		Limite: 500,
	}
	repo.db.Create(&contaCorrente)

	// Test the method
	result, err := repo.GetContaCorrenteByID(int64(contaCorrente.ID))
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if result.ID != contaCorrente.ID {
		t.Errorf("Expected ID %d, got %d", contaCorrente.ID, result.ID)
	}
}

func TestGetLast10MovimentosByContaCorrenteID(t *testing.T) {
	setup()

	// Create a Movimento for testing
	movimento := models.Movimento{
		IDCliente:     1,
		Valor:         100,
		Tipo:          "d",
		Descricao:     "Test",
		DataMovimento: time.Now(),
	}
	repo.db.Create(&movimento)

	// Test the method
	result, err := repo.GetLast10MovimentosByContaCorrenteID(movimento.IDCliente)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	if len(result) != 1 {
		t.Errorf("Expected 1, got %d", len(result))
	}
}

func TestSaveContaCorrente(t *testing.T) {
	setup()

	// Create a ContaCorrente for testing
	contaCorrente := models.ContaCorrente{
		Nome:   "Test",
		Saldo:  1000,
		Limite: 500,
	}

	// Test the method
	err := repo.SaveContaCorrente(contaCorrente)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestSaveMovimento(t *testing.T) {
	setup()

	// Create a Movimento for testing
	movimento := models.Movimento{
		IDCliente:     1,
		Valor:         100,
		Tipo:          "d",
		Descricao:     "Test",
		DataMovimento: time.Now(),
	}

	// Test the method
	err := repo.SaveMovimento(movimento)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
