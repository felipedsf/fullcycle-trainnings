package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	//GetById(id string) (*Order, error)
	List() (*[]Order, error)
}
