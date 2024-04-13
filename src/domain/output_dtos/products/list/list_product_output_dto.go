package output_dtos

import (
	entities "my-app/src/domain/entities/products"
)

type ListProductOutputDTO struct {
	Products []entities.Product
}
