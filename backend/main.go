package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/optimus/backend/config"
	"github.com/optimus/backend/controllers"
	"github.com/optimus/backend/models"
	"github.com/optimus/backend/routes"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env file not found")
	}

	r.SetTrustedProxies(nil)
	r.Static("/uploads", "./uploads")

	config.ConnectDB()
	go controllers.StartLPSECron()

	config.DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.PokjaGroups{},
		// &models.UserPokjaGroups{},
		// &models.TenderPaket{},
		&models.RabHeader{},
		&models.RabDetail{},
		&models.ScheduleHeader{},
		&models.ScheduleItem{},
		&models.ScheduleWeek{},
		&models.RealisasiHeader{},
		&models.RealisasiDetail{},
		// &models.RealisasiWeek{},
		&models.NonTenderKontractIsb{},
		&models.NonTenderIsb{},
		&models.NonTenderSelesaiIsb{},
		&models.NonTenderTahapIsb{},
		&models.PencatatanNonTenderIsb{},
		&models.TenderIsb{},
		&models.TenderKontrakIsb{},
		&models.TenderSelesaiIsb{},
		&models.TenderTahapIsb{},
		&models.PaketPurchasing{},
		&models.RupPaketPenyedia{},
		&models.RupPaketSwakelola{},
		&models.RupPenyediaTerumumkan{},
		&models.DataEntry{},
	)

	config.DB.AutoMigrate(
		&models.Url{},
	)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(r)

	r.Run("0.0.0.0:8081")
}
