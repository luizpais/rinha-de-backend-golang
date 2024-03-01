package services

import (
	"errors"
	"github.com/luizpais/rinha-de-backend-go/rdbe/internal/models"
	"github.com/luizpais/rinha-de-backend-go/rdbe/internal/repositories"
	"time"
)

type ContaCorrenteService struct {
	repository *repositories.Repository
}

func NewContaCorrenteService(repository *repositories.Repository) *ContaCorrenteService {
	return &ContaCorrenteService{
		repository: repository,
	}
}

func (s *ContaCorrenteService) Extrato(id int64) (models.ExtratoResponse, error) {
	contaCorrente, err := s.repository.GetContaCorrenteByID(id)
	if err != nil {
		return models.ExtratoResponse{}, err
	}

	extrato := models.ExtratoResponse{
		Saldo: models.SaldoAtual{
			Total:       contaCorrente.Saldo,
			DataExtrato: time.Now(),
			Limite:      contaCorrente.Limite,
		},
	}

	movimentos, err := s.repository.GetLast10MovimentosByContaCorrenteID(id)
	if err != nil {
		return models.ExtratoResponse{}, err
	}

	for _, movimento := range movimentos {
		extrato.UltimasTransacoes = append(extrato.UltimasTransacoes, models.Transacao{
			Valor:       movimento.Valor,
			Tipo:        movimento.Tipo,
			Descricao:   movimento.Descricao,
			RealizadaEm: movimento.DataMovimento,
		})
	}

	return extrato, nil
}

func (s *ContaCorrenteService) Transacao(id int64, request models.TransacaoRequest) (models.TransacaoResponse, error) {
	contaCorrente, err := s.repository.GetContaCorrenteByID(id)
	if err != nil {
		return models.TransacaoResponse{}, err
	}

	if request.Tipo == "d" {
		if contaCorrente.Saldo-int64(request.Valor) < -contaCorrente.Limite {
			return models.TransacaoResponse{}, errors.New("saldo insuficiente")
		}
		contaCorrente.Saldo -= int64(request.Valor)
	} else if request.Tipo == "c" {
		contaCorrente.Saldo += int64(request.Valor)
	}

	if s.repository.SaveContaCorrente(contaCorrente) != nil {
		return models.TransacaoResponse{}, errors.New("erro ao salvar conta corrente")
	}

	movimento := models.Movimento{
		DataMovimento: time.Now(),
		Descricao:     request.Descricao,
		IDCliente:     int64(contaCorrente.ID),
		Tipo:          request.Tipo,
		Valor:         int64(request.Valor),
	}

	// Persist the Movimento
	// This is a placeholder and should be replaced with actual data persistence
	// Persist the Movimento
	err = s.repository.SaveMovimento(movimento)
	if err != nil {
		return models.TransacaoResponse{}, err
	}

	response := models.TransacaoResponse{
		Limite: int(contaCorrente.Limite),
		Saldo:  int(contaCorrente.Saldo),
	}

	return response, nil
}
