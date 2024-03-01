package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/luizpais/rinha-de-backend-go/rdbe/internal/models"
	"github.com/luizpais/rinha-de-backend-go/rdbe/internal/services"
	"net/http"
	"strconv"
)

type ContaCorrenteResource struct {
	contaCorrenteService *services.ContaCorrenteService
}

func NewContaCorrenteResource(service *services.ContaCorrenteService) *ContaCorrenteResource {
	return &ContaCorrenteResource{
		contaCorrenteService: service,
	}
}

func (c *ContaCorrenteResource) Extrato(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	extrato, _ := c.contaCorrenteService.Extrato(id)
	ctx.JSON(http.StatusOK, extrato)
}

func (c *ContaCorrenteResource) Transacao(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var request models.TransacaoRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if request.Descricao == "" || len(request.Descricao) > 10 || request.Tipo == "" || (request.Tipo != "d" && request.Tipo != "c") || request.Valor <= 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request parameters"})
		return
	}

	transacao, err := c.contaCorrenteService.Transacao(id, request)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Transaction failed"})
		return
	}

	ctx.JSON(http.StatusOK, transacao)
}
