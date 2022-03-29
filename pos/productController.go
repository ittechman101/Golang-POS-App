package pos

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	database *gorm.DB
}

func (repository *ProductRepository) FindAllProduct(c *fiber.Ctx) []Products {
	var products []Products

	db := repository.database
	if len(c.Query("limit")) > 0 {
		db = db.Limit(c.Query("limit"))
	}
	if len(c.Query("skip")) > 0 {
		db = db.Offset(c.Query("skip"))
	}
	db.Find(&products)
	return products
}

func (repository *ProductRepository) FindProduct(id int) (Products, error) {
	var product Products
	err := repository.database.Where("product_id = ?", id).First(&product).Error
	if err != nil {
		err = errors.New("product not found")
	}
	return product, err
}

func (repository *ProductRepository) CreateProduct(product Products) (Products, error) {
	// Get Max productId
	var maxProduct Products

	repository.database.Raw(`
		SELECT COALESCE(MAX(product_id) + 1, 0) as product_id
		FROM products
		`).Scan(
		&maxProduct,
	)
	//	product.Passcode = int64(rand.Intn(899999) + 100000)
	product.ProductId = maxProduct.ProductId
	err := repository.database.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (repository *ProductRepository) SaveProduct(product Products) (Products, error) {
	err := repository.database.Table("products").Where("product_id = ?", product.ProductId).Update(product).Error
	return product, err
}

func (repository *ProductRepository) DeleteProduct(id int) int64 {
	count := repository.database.Where("product_id = ?", id).Delete(&Products{}).RowsAffected
	return count
}

func NewProductRepository(database *gorm.DB) *ProductRepository {
	return &ProductRepository{
		database: database,
	}
}
