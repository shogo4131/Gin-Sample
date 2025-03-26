package repository

import (
	"errors"
	"go-gin-sample/apps/model"

	"gorm.io/gorm"
)

type IItemRepository interface {
	FindAll() (*[]model.Item, error)
	FindById(itemId uint, userId uint) (*model.Item, error)
	Create(newItem model.Item) (*model.Item, error)
	Update(updateItem model.Item) (*model.Item, error)
	Delete(itemId uint, userId uint) error
}

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) Create(newItem model.Item) (*model.Item, error) {
	result := r.db.Create(&newItem)

	if result.Error != nil {
		return nil, result.Error
	}

	return &newItem, nil
}

func (r *ItemRepository) FindAll() (*[]model.Item, error) {
	var items []model.Item

	result := r.db.Find(&items)

	if result.Error != nil {
		return nil, result.Error
	}

	return &items, nil
}

func (r *ItemRepository) FindById(itemId uint, userId uint) (*model.Item, error) {
	var item model.Item

	result := r.db.First(&item, "id = ? AND user_id", itemId, userId)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("item not found")
		}
		return nil, result.Error
	}

	return &item, nil
}

func (r *ItemRepository) Update(updateItem model.Item) (*model.Item, error) {
	result := r.db.Save(&updateItem)

	if result.Error != nil {
		return nil, result.Error
	}

	return &updateItem, nil
}

func (r *ItemRepository) Delete(itemId uint, userId uint) error {
	deleteItem, err := r.FindById(itemId, userId)

	if err != nil {
		return err
	}

	result := r.db.Unscoped().Delete(&model.Item{}, deleteItem)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
