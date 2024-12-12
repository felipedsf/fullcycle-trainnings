package usecase

import (
	"github.com/felipedsf/go-samples/clean-architecture/internal/entity"
)

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrderUseCase) Get() ([]OrderOutputDTO, error) {
	var dto []OrderOutputDTO

	orders, err := c.OrderRepository.List()
	if err != nil {
		return nil, err
	}
	for _, order := range *orders {
		dto = append(dto, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		})
	}

	return dto, nil
}
