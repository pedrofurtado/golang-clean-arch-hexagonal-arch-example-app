package repositories

import (
	"os"
	"fmt"
	"strings"

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
