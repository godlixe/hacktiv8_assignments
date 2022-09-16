package service

import (
	"context"
	"order_service/dto"
	"order_service/entity"
	"order_service/repository"

	"github.com/mashingan/smapping"
)

type OrderService interface {
	GetOrders(context.Context) ([]entity.Order, error)
	GetOrderByID(context.Context, uint64) (entity.OrderWithPerson, error)
	CreateOrder(context.Context, dto.OrderCreateDTO) (entity.Order, error)
	UpdateOrder(context.Context, dto.OrderUpdateDTO) (entity.Order, error)
	DeleteOrder(context.Context, uint64) error
}

type orderService struct {
	orderRepository  repository.OrderRepository
	itemRepository   repository.ItemRepository
	personRepository repository.PersonRepository
}

func NewOrderService(or repository.OrderRepository, ir repository.ItemRepository, pr repository.PersonRepository) OrderService {
	return &orderService{
		orderRepository:  or,
		itemRepository:   ir,
		personRepository: pr,
	}
}

func (s *orderService) GetOrders(ctx context.Context) ([]entity.Order, error) {
	result, err := s.orderRepository.GetOrders(ctx)
	return result, err
}

func (s *orderService) GetOrderByID(ctx context.Context, id uint64) (entity.OrderWithPerson, error) {
	var orderWithPerson entity.OrderWithPerson
	order, err := s.orderRepository.GetOrderByID(ctx, id)
	if err != nil {
		return orderWithPerson, err
	}

	person, err := s.personRepository.GetPerson()
	if err != nil {
		return orderWithPerson, err
	}

	err = smapping.FillStruct(&orderWithPerson, smapping.MapFields(&order))
	if err != nil {
		return orderWithPerson, err
	}

	orderWithPerson.Person = person

	return orderWithPerson, nil
}

func (s *orderService) CreateOrder(ctx context.Context, orderDTO dto.OrderCreateDTO) (entity.Order, error) {
	var createdOrder entity.Order
	err := smapping.FillStruct(&createdOrder, smapping.MapFields(&orderDTO))
	if err != nil {
		return createdOrder, err
	}
	result, err := s.orderRepository.CreateOrder(ctx, createdOrder)
	return result, err
}

func (s *orderService) UpdateOrder(ctx context.Context, orderDTO dto.OrderUpdateDTO) (entity.Order, error) {
	var updatedOrder entity.Order
	err := smapping.FillStruct(&updatedOrder, smapping.MapFields(&orderDTO))
	if err != nil {
		return updatedOrder, err
	}
	result, err := s.orderRepository.UpdateOrder(ctx, updatedOrder)
	return result, err
}

func (s *orderService) DeleteOrder(ctx context.Context, id uint64) error {
	return s.orderRepository.DeleteOrder(ctx, id)
}
