package repositories

import (
	"os"
	"fmt"
	"strings"

	repositoriesDatabaseInterfaces "my-app/internal/domain/repositories/database_interfaces"
	createDTO "my-app/internal/domain/input_dtos/products/create"
	listDTO "my-app/internal/domain/input_dtos/products/list"
	entities "my-app/internal/domain/entities/products"
)

type ProductRepositoryInterface interface {
	Insert(product createDTO.CreateProductInputDTO) error
	ListBy(filters listDTO.ListProductInputDTO) ([]entities.Product, error)
}

type ProductRepository struct {
	db repositoriesDatabaseInterfaces.RepositoryDatabase
}

func (productRepository ProductRepository) Insert(productInputDTO createDTO.CreateProductInputDTO) error {
	_, err := productRepository.db.GetDB().Exec(
		fmt.Sprintf("INSERT INTO products(identifier, full_name, state_name) VALUES (%s)", generateNPreparedStatementBindings(3)),
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

func (productRepository ProductRepository) ListBy(listInputDTO listDTO.ListProductInputDTO) ([]entities.Product, error) {
	rows, err := productRepository.db.GetDB().Query("SELECT identifier, full_name, state_name FROM products")

	if err != nil {
		fmt.Println("Error in ProductRepository::ListBy. Error %v", err)
		return nil, err
	}

	products := []entities.Product{}

	for rows.Next() {
		var product entities.Product

		err = rows.Scan(&product.Identifier, &product.FullName, &product.StateName)

		if err != nil {
			msg := fmt.Sprintf("Error in ListBy::Scan. Error %v", err)
			fmt.Println(msg)
			return nil, err
		}

		products = append(products, product)
	}

	fmt.Println("ProductRepository::ListBy executed successfully")

	return products, nil
}

func NewProductRepository(db repositoriesDatabaseInterfaces.RepositoryDatabase) ProductRepositoryInterface {
	return &ProductRepository{
		db: db,
	}
}

func generateNPreparedStatementBindings(n int) string {
	switch os.Getenv("APP_ADAPTER_DATABASE_DRIVER") {
		case "mysql":
			var statements []string

			for i := 0; i < n; i++ {
				statements = append(statements, "?")
			}

			return strings.Join(statements[:], ", ")
		case "postgres":
			var statements []string

			for i := 0; i < n; i++ {
				statements = append(statements, fmt.Sprintf("$%d", i + 1))
			}

			return strings.Join(statements[:], ", ")
		default:
			panic("Must be defined a database driver, it is not possible to generate prepared statements bindings")
	}
}
