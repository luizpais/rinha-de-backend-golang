package repositories

import (
	"github.com/luizpais/rinha-de-backend-go/rdbe/internal/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetContaCorrenteByID(id int64) (models.ContaCorrente, error) {
	var contaCorrente models.ContaCorrente
	result := r.db.First(&contaCorrente, id)
	return contaCorrente, result.Error
}

func (r *Repository) GetLast10MovimentosByContaCorrenteID(id int64) ([]models.Movimento, error) {
	var movimentos []models.Movimento
	result := r.db.Where("idcliente = ?", id).Order("datamovimento desc").Limit(10).Find(&movimentos)
	return movimentos, result.Error
}

func (r *Repository) SaveContaCorrente(contaCorrente models.ContaCorrente) error {
	result := r.db.Save(contaCorrente)
	return result.Error
}

func (r *Repository) SaveMovimento(movimento models.Movimento) error {
	result := r.db.Create(&movimento)
	return result.Error
}
