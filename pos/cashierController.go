package pos

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type CashierRepository struct {
	database *gorm.DB
}

func (repository *CashierRepository) FindAll(c *fiber.Ctx) []Cashiers {
	var cashiers []Cashiers

	db := repository.database
	if len(c.Query("limit")) > 0 {
		db = db.Limit(c.Query("limit"))
	}
	if len(c.Query("skip")) > 0 {
		db = db.Offset(c.Query("skip"))
	}
	db.Find(&cashiers)
	return cashiers
}

func (repository *CashierRepository) Find(id int) (Cashiers, error) {
	var cashier Cashiers
	err := repository.database.Where("cashier_id = ?", id).First(&cashier).Error
	if err != nil {
		err = errors.New("cashier not found")
	}
	return cashier, err
}

func (repository *CashierRepository) Passcode(id int) (int, error) {
	var cashier Cashiers
	err := repository.database.Where("cashier_id = ?", id).First(&cashier).Error
	if err != nil {
		err = errors.New("cashier not found")
	}
	return int(cashier.Passcode), err
}

func (repository *CashierRepository) Create(cashier Cashiers) (Cashiers, error) {

	err := repository.database.Create(&cashier).Error
	if err != nil {
		return cashier, err
	}

	return cashier, nil
}

func (repository *CashierRepository) Save(cashier Cashiers) (Cashiers, error) {
	err := repository.database.Table("cashiers").Where("cashier_id = ?", cashier.ID).Update(cashier).Error
	return cashier, err
}

func (repository *CashierRepository) Delete(id int) int64 {
	count := repository.database.Where("cashier_id = ?", id).Delete(&Cashiers{}).RowsAffected
	return count
}

func NewCashierRepository(database *gorm.DB) *CashierRepository {
	return &CashierRepository{
		database: database,
	}
}
