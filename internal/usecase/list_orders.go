package usecase

import (
	"github.com/mmunareto/clean-arch/internal/entity"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return []OrderOutputDTO{}, err
	}

	var ordersList []OrderOutputDTO
	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		}
		ordersList = append(ordersList, dto)
	}

	return ordersList, nil
}
