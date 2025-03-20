package repository

import (
	"errors"
	"go-gin-sample/apps/model"
)

type IItemRepository interface {
	FindAll() (*[]model.Item, error)
	FindById(itemId uint) (*model.Item, error)
	Create(newItem model.Item) (*model.Item, error)
	Update(updateItem model.Item) (*model.Item, error)
	Delete(itemId uint) error
}

type ItemMemoryRepository struct {
	items []model.Item
}

func NewItemMemoryRepository(items []model.Item) IItemRepository {
	return &ItemMemoryRepository{items: items}
}

func (r *ItemMemoryRepository) FindAll() (*[]model.Item, error) {
	return &r.items, nil
}

func (r *ItemMemoryRepository) FindById(itemId uint) (*model.Item, error) {
	for _, v := range r.items {
		if v.ID == itemId {
			return &v, nil
		}
	}

	return nil, errors.New("Item not found")
}

func (r *ItemMemoryRepository) Create(newItem model.Item) (*model.Item, error) {
	newItem.ID = uint(len(r.items) + 1)
	r.items = append(r.items, newItem)

	return &newItem, nil
}

func (r *ItemMemoryRepository) Update(updateItem model.Item) (*model.Item, error) {
	for i, v := range r.items {
		if v.ID == updateItem.ID {
			r.items[i] = updateItem
			return &updateItem, nil
		}
	}

	return nil, errors.New("Unexpected error")
}

func (r *ItemMemoryRepository) Delete(itemId uint) error {
	for i, v := range r.items {
		if v.ID == itemId {
			r.items = append(r.items[:i], r.items[i+1:]...)
		}
		return nil
	}

	return errors.New("Item not found")
}
