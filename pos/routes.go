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

	router.Get("/cashiers", cashierHandler.GetAll)
	router.Get("/cashiers/:id", cashierHandler.Get)
	router.Get("/cashiers/:id/passcode", cashierHandler.Passcode)
	router.Put("/cashiers/:id", cashierHandler.Update)
	router.Post("/cashiers", cashierHandler.Create)
	router.Delete("/cashiers/:id", cashierHandler.Delete)
}
