package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/zkgogreen/bisago/domain"
	"github.com/zkgogreen/bisago/dto"
)

type customerService struct {
	customerRepository domain.CustomerRepository
}

func NewCustomer(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

func (c customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := c.customerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var customerDatas []dto.CustomerData
	for _, customer := range customers {
		customerDatas = append(customerDatas, dto.CustomerData{
			ID:   customer.ID,
			Code: customer.Code,
			Name: customer.Name,
		})
	}
	return customerDatas, nil
}

func (c customerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	customer := domain.Customer{
		Code:      req.Code,
		Name:      req.Name,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}
	return c.customerRepository.Save(ctx, &customer)
}
