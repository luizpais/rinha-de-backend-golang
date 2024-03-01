package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/luizpais/rinha-de-backend-go/rdbe/internal/controllers"
	"github.com/luizpais/rinha-de-backend-go/rdbe/internal/repositories"
	"github.com/luizpais/rinha-de-backend-go/rdbe/internal/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {

	// Configure dependencies
	dsn := os.Getenv("DATABASE_DSN")
	db, dberr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dberr != nil {
		fmt.Println(dberr.Error())
		return
	} // listen and serve on 0.0.0.0:8080
	repository := repositories.NewRepository(db)
	service := services.NewContaCorrenteService(repository)
	controller := controllers.NewContaCorrenteResource(service)

	r := gin.Default()

	r.GET("/clientes/:id/extrato", controller.Extrato)
	r.POST("/clientes/:id/transacoes", controller.Transacao)

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
