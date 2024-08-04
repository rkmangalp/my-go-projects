package main

import (
	"fmt"
	"log"
	"my-go-projects/go-fiber-CRM-basic/database"
	"my-go-projects/go-fiber-CRM-basic/lead"

	"github.com/gofiber/fiber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	dsn := "root:new_password@tcp(127.0.0.1:3306)/LeadsDB?charset=utf8&parseTime=True&loc=Local"
	var err error
	database.DBconn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to DB")
	}

	fmt.Println("connection opened to database")
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")

}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)

	db, err := database.DBconn.DB()
	if err != nil {
		log.Fatalf("failed to get underlying DB: %v", err)
	}
	defer db.Close()

}
