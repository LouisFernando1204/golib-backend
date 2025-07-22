package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/LouisFernando1204/golang-restapi.git/domain"
	"github.com/LouisFernando1204/golang-restapi.git/dto"
	"github.com/google/uuid"
)

type customerService struct {
	customerRepository domain.CustomerRepository
}

func NewCustomer(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

// Index implements domain.CustomerService.
func (c *customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := c.customerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var customerData []dto.CustomerData
	for _, v := range customers {
		customerData = append(customerData, dto.CustomerData{
			ID:   v.ID,
			Code: v.Code,
			Name: v.Name,
		})
	}

	return customerData, nil
}

// Show implements domain.CustomerService.
func (c *customerService) Show(ctx context.Context, id string) (dto.CustomerData, error) {
	persisted, err := c.customerRepository.FindById(ctx, id)
	if err != nil {
		return dto.CustomerData{}, err
	}

	if persisted.ID == "" {
		return dto.CustomerData{}, errors.New("customer data not found")
	}

	return dto.CustomerData{
		ID:   persisted.ID,
		Code: persisted.Code,
		Name: persisted.Name,
	}, nil
}

// Create implements domain.CustomerService.
func (c *customerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	customer := domain.Customer{
		ID:        uuid.NewString(),
		Code:      req.Code,
		Name:      req.Name,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}

	return c.customerRepository.Save(ctx, &customer)
}

// Update implements domain.CustomerService.
func (c *customerService) Update(ctx context.Context, req dto.UpdateCustomerRequest) error {
	persisted, err := c.customerRepository.FindById(ctx, req.ID)
	if err != nil {
		return err
	}

	if persisted.ID == "" {
		return errors.New("customer data not found")
	}

	persisted.Code = req.Code
	persisted.Name = req.Name
	persisted.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}

	return c.customerRepository.Update(ctx, &persisted)
}

// Delete implements domain.CustomerService.
func (c *customerService) Delete(ctx context.Context, id string) error {
	exist, err := c.customerRepository.FindById(ctx, id)
	if err != nil {
		return err
	}

	if exist.ID == "" {
		return errors.New("customer data not found")
	}

	return c.customerRepository.Delete(ctx, exist.ID)
}
