package repository

import (
	"context"
	"order_service/entity"

	"gorm.io/gorm"
)

type ItemRepository interface {
	InsertItem(context.Context, entity.Item) (entity.Item, error)
}

type itemRepository struct {
	connection *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{
		connection: db,
	}
}
func (db *itemRepository) InsertItem(ctx context.Context, item entity.Item) (entity.Item, error) {

	if err := db.connection.WithContext(ctx).Create(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}
