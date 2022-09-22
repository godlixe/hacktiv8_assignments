package repository

import (
	"context"
	"jwt-hacktiv8/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductByID(ctx context.Context, id uint) (entity.Product, error)
	InsertProduct(ctx context.Context, product entity.Product) (entity.Product, error)
	UpdateProduct(ctx context.Context, product entity.Product) (entity.Product, error)
}

type productRepository struct {
	connection *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		connection: db,
	}
}

func (db *productRepository) GetProductByID(ctx context.Context, id uint) (entity.Product, error) {
	var product entity.Product
	tx := db.connection.Where(("id = ?"), id).Find(&product)
	if tx.Error != nil {
		return entity.Product{}, tx.Error
	}
	return product, nil
}

func (db *productRepository) InsertProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	tx := db.connection.Create(&product)
	if tx.Error != nil {
		return entity.Product{}, tx.Error
	}
	return product, nil
}

func (db *productRepository) UpdateProduct(ctx context.Context, product entity.Product) (entity.Product, error) {
	tx := db.connection.Debug().Where(("id = ?"), product.ID).Updates(&product)
	if tx.Error != nil {
		return entity.Product{}, tx.Error
	}
	return product, nil
}
