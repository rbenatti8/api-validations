package handler

type UserAddressDTO struct {
	Street   string `json:"street" validate:"required,min=3"`
	Number   int    `json:"number" validate:"required,min=1"`
	Postcode string `json:"postcode" validate:"required,min=3"`
	City     string `json:"city" validate:"required,min=3"`
}
type CreateUserDTO struct {
	Name           string         `json:"name" validate:"required,min=3"`
	Age            int            `json:"age" validate:"required,min=1"`
	DocumentNumber string         `json:"document_number" validate:"required,len=11"`
	Email          string         `json:"email" validate:"email,omitempty"`
	Address        UserAddressDTO `json:"address" validate:"required"`
}
