package main

import (
	"bootcamp-api-hmsi/connectdb"
	"bootcamp-api-hmsi/moduls/customers/customerHandler"
	"bootcamp-api-hmsi/moduls/customers/customerRepository"
	"bootcamp-api-hmsi/moduls/customers/customerUseCase"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_DRIVER := os.Getenv("DB_DRIVER")
	PORT := os.Getenv("PORT")

	log.Info().Msg(DB_HOST)
	log.Info().Msg(DB_NAME)
	log.Info().Msg(DB_PORT)
	log.Info().Msg(DB_USER)
	log.Info().Msg(DB_PASSWORD)
	log.Info().Msg(DB_DRIVER)
	log.Info().Msg(PORT)

	db, errConn := connectdb.GetConnPostgress(DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_DRIVER)

	if errConn != nil {
		log.Error().Msg(errConn.Error())
		os.Exit(1)
	}
	log.Info().Msg("Connected to database")

	// Inisisalisasi router
	router := gin.Default()

	// inisisalisasi modules
	customerRepo := customerRepository.NewCustomerRepository(db)
	customerUC := customerUseCase.NewCustomerUseCase(customerRepo)
	customerHandler.NewCustomerHandler(router, customerUC)

	log.Info().Msg("Server running on port" + PORT)
	router.Run(":" + PORT)
}
