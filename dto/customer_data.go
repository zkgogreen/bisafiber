package dto

type CustomerData struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type CreateCustomerRequest struct {
	Code string `json:"code" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateCustomerData struct {
	ID   int    `json:"-"`
	Code string `json:"code" validate:"required"`
	Name string `json:"name" validate:"required"`
}
