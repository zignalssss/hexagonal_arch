package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/zignalssss/hexagonal_arch/adapters"
	"github.com/zignalssss/hexagonal_arch/core"

	// "github.com/zignalssss/hexagonal_arch/adapters"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

func main() {
	app := fiber.New()

	// Configure PostgreSQL database
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	//Secondary to Primary Port

	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := core.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	app.Post("/order", orderHandler.CreateOrder)

	// Migrate the schema
	db.AutoMigrate(&core.Order{})

	app.Listen((":8000"))
}
