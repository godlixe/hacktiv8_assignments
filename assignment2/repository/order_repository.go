package repository

import (
	"context"
	"fmt"
	"order_service/entity"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetOrders(context.Context) ([]entity.Order, error)
	CreateOrder(context.Context, entity.Order) (entity.Order, error)
	UpdateOrder(context.Context, entity.Order) (entity.Order, error)
	DeleteOrder(context.Context, uint64) error
}

type orderRepository struct {
	connection *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		connection: db,
	}
}

func (db *orderRepository) GetOrders(ctx context.Context) ([]entity.Order, error) {
	var order []entity.Order

	if err := db.connection.WithContext(ctx).Preload("Items").Find(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (db *orderRepository) CreateOrder(ctx context.Context, order entity.Order) (entity.Order, error) {

	if err := db.connection.WithContext(ctx).Create(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (db *orderRepository) UpdateOrder(ctx context.Context, order entity.Order) (entity.Order, error) {
	fmt.Println("ni dari repo", order.Items)
	if err := db.connection.Debug().WithContext(ctx).Model(&order).Association("Items").Replace(order.Items); err != nil {
		return order, err
	}
	// fmt.Println(order.Items)
	if err := db.connection.Debug().WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Preload("Item").Save(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (db *orderRepository) DeleteOrder(ctx context.Context, id uint64) error {

	if err := db.connection.WithContext(ctx).Where("id = ?", id).Delete(&entity.Order{}).Error; err != nil {
		return err
	}
	return nil
}
