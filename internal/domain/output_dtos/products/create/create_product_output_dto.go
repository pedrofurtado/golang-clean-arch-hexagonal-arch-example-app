package output_dtos

import (
	entities "my-app/internal/domain/entities/products"
)

type CreateProductOutputDTO struct {
	Products []entities.Product `json:"products"`
}
