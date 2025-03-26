package services

import (
	"go-gin-sample/apps/dto"
	"go-gin-sample/apps/model"
	"go-gin-sample/apps/repository"
)

type IItemService interface {
	FindAll() (*[]model.Item, error)
	FindById(itemId uint, userId uint) (*model.Item, error)
	Create(newItem dto.CreateItemInput, userId uint) (*model.Item, error)
	Update(itemId uint, userId uint, updateItemInput dto.UpdateItemInput) (*model.Item, error)
	Delete(itemId uint, userId uint) error
}

type ItemService struct {
	repository repository.IItemRepository
}

func NewItemService(repository repository.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) FindAll() (*[]model.Item, error) {
	return s.repository.FindAll()
}

func (s *ItemService) FindById(itemId uint, userId uint) (*model.Item, error) {
	return s.repository.FindById(itemId, userId)
}

func (s *ItemService) Create(createItemInput dto.CreateItemInput, userId uint) (*model.Item, error) {
	newItem := model.Item{
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		SoldOut:     false,
		UserID:      userId,
	}

	return s.repository.Create(newItem)
}

func (s *ItemService) Update(itemId uint, userId uint, updateItemInput dto.UpdateItemInput) (*model.Item, error) {
	targetItem, err := s.FindById(itemId, userId)

	if err != nil {
		return nil, err
	}

	if updateItemInput.Name != nil {
		targetItem.Name = *updateItemInput.Name
	}

	if updateItemInput.Price != nil {
		targetItem.Price = *updateItemInput.Price
	}

	if updateItemInput.Description != nil {
		targetItem.Description = *updateItemInput.Description
	}

	if updateItemInput.SoldOut != nil {
		targetItem.SoldOut = *updateItemInput.SoldOut
	}

	return s.repository.Update(*targetItem)
}

func (s *ItemService) Delete(itemId uint, userId uint) error {
	return s.repository.Delete(itemId, userId)
}
