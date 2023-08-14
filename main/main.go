package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lcslima45/desafio-padawan-go/models"

	"github.com/lcslima45/desafio-padawan-go/controllers"
	"github.com/lcslima45/desafio-padawan-go/converter"
	"github.com/lcslima45/desafio-padawan-go/repository"
	"github.com/lcslima45/desafio-padawan-go/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Carregar variáveis de ambiente do arquivo "environment"

	router := gin.Default()

	dsn := "admin:admin@tcp(localhost:3306)/conversions_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// AutoMigrate
	db.AutoMigrate(&models.ConversionModel{})
	db.Migrator().DropColumn(&models.ConversionModel{}, "contime")

	// Inicializar o repositório
	conversionRepo := repository.NewRepository(db)

	// Inicializar o serviço de conversão
	converterService := converter.NewConverterService(conversionRepo)

	// Inicializar o controlador de conversão
	conversionController := controllers.NewConversionController(converterService)

	// Definir as rotas
	routes.ConvertRoute(router, conversionController)

	// Iniciar o servidor Gin na porta 8000
	portAPI := ":8000"
	log.Println("Server is running on port", portAPI)
	router.Run(portAPI)
}
