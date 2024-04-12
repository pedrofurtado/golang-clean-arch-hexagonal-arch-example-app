package repositories

import (
	"fmt"
	repositoriesDatabaseInterfaces "my-app/src/domain/repositories/database_interfaces"
	productInputDTO "my-app/src/domain/input_dtos/products"
)

type ProductRepositoryInterface interface {
	Insert(product productInputDTO.ProductInputDTO) error
}

type ProductRepository struct {
	db repositoriesDatabaseInterfaces.RepositoryDatabase
}

func (productRepository ProductRepository) Insert(productInputDTO productInputDTO.ProductInputDTO) error {
	_, err := productRepository.db.GetDB().Exec(
		"INSERT INTO products (identifier, full_name, state_name) VALUES (?, ?, ?)",
		productInputDTO.Identifier,
		productInputDTO.FullName,
		productInputDTO.StateName,
	)

	if err != nil {
		fmt.Println("Error in ProductRepository::Insert. Error %v", err)
		return err
	}

	fmt.Println("ProductRepository::Insert executed successfully")
	return nil
}

func NewProductRepository(db repositoriesDatabaseInterfaces.RepositoryDatabase) ProductRepositoryInterface {
	return &ProductRepository{
		db: db,
	}
}
