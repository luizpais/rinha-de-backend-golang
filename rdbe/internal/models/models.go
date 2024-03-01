package models

import (
	"gorm.io/gorm"
	"time"
)

//DTOs

type ExtratoResponse struct {
	Saldo             SaldoAtual  `json:"saldo"`
	UltimasTransacoes []Transacao `json:"ultimas_transacoes"`
}

type SaldoAtual struct {
	Total       int64     `json:"total"`
	DataExtrato time.Time `json:"data_extrato"`
	Limite      int64     `json:"limite"`
}

type Transacao struct {
	Valor       int64     `json:"valor"`
	Tipo        string    `json:"tipo"`
	Descricao   string    `json:"descricao"`
	RealizadaEm time.Time `json:"realizada_em"`
}

type TransacaoRequest struct {
	Valor     int    `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

type SaldoResponse struct {
	Total       int    `json:"total"`
	DataExtrato string `json:"data_extrato"`
	Limite      int    `json:"limite"`
}
type TransacaoResponse struct {
	Limite int `json:"limite"`
	Saldo  int `json:"saldo"`
}

// Entidades
type ContaCorrente struct {
	gorm.Model
	Nome   string `gorm:"column:nome"`
	Saldo  int64  `gorm:"column:saldo"`
	Limite int64  `gorm:"column:limite"`
}

func (c *ContaCorrente) TableName() string {
	return "contacorrente"
}

type Movimento struct {
	gorm.Model
	IDCliente     int64     `gorm:"column:idcliente"`
	Valor         int64     `gorm:"column:valor"`
	Tipo          string    `gorm:"column:tipo"`
	Descricao     string    `gorm:"column:descricao"`
	DataMovimento time.Time `gorm:"column:datamovimento"`
}

func (m *Movimento) TableName() string {
	return "movimento"
}

func FindAteDezMovimentosByIdCliente(db *gorm.DB, idCliente int64) ([]Movimento, error) {
	var movimentos []Movimento
	result := db.Where("id_cliente = ?", idCliente).Order("data_movimento desc").Limit(10).Find(&movimentos)
	return movimentos, result.Error
}
