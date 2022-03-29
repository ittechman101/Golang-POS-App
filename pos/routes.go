package pos

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func Register(router fiber.Router, database *gorm.DB) {
	database.AutoMigrate(&Cashiers{})
	//	database.AutoMigrate(&Product{})

	cashierRepository := NewCashierRepository(database)
	cashierHandler := NewCashierHandler(cashierRepository)

	router.Get("/cashiers", cashierHandler.GetAllCashier)
	router.Get("/cashiers/:id", cashierHandler.GetCashier)
	router.Get("/cashiers/:id/passcode", cashierHandler.Passcode)
	router.Post("/cashiers/:id/login", cashierHandler.Login)
	router.Put("/cashiers/:id", cashierHandler.UpdateCashier)
	router.Post("/cashiers", cashierHandler.CreateCashier)
	router.Delete("/cashiers/:id", cashierHandler.DeleteCashier)
}
