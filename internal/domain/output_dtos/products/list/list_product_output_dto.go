package output_dtos

import (
	entities "my-app/internal/domain/entities/products"
)

type ListProductOutputDTO struct {
	Products []entities.Product `json:"products"`
}
