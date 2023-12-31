package usecase_test

import (
	"testing"

	"github.com/paulojr-eco/full-cycle-desafio/application/usecase"
	"github.com/paulojr-eco/full-cycle-desafio/domain/model"
	"github.com/paulojr-eco/full-cycle-desafio/infra/db"
	"github.com/paulojr-eco/full-cycle-desafio/infra/repository"
	"github.com/stretchr/testify/require"
)

/* Should add and return a product on success */
func TestUseCase_AddProduct(t *testing.T) {
	name := "Product 01"
	description := "Description from product 01"
	price := float32(15.99)
	database, _ := db.ConnectDB()
	productRepository := repository.ProductRepositoryDb{Db: database}
	productUseCase := usecase.ProductUseCase{ProductRepository: productRepository}

	product, _ := model.NewProduct(name, description, price)
	product, err := productUseCase.ProductRepository.AddProduct(product)

	require.Nil(t, err)
	require.NotNil(t, product)
	require.NotEmpty(t, product.ID)
	require.Equal(t, product.Name, name)
	require.Equal(t, product.Description, description)
	require.Equal(t, product.Price, price)
	require.NotNil(t, product.CreatedAt)
	require.NotNil(t, product.UpdatedAt)
}

/* Should return an empty array if there is no products */
func TestUseCase_FindZeroProducts(t *testing.T) {
	database, _ := db.ConnectDB()
	productRepository := repository.ProductRepositoryDb{Db: database}
	productUseCase := usecase.ProductUseCase{ProductRepository: productRepository}

	products := productUseCase.ProductRepository.FindProducts()

	require.NotNil(t, products)
	require.Equal(t, len(*products), 0)
}

/* Should return an array with a product on success */
func TestUseCase_FindProducts(t *testing.T) {
	name := "Product 01"
	description := "Description from product 01"
	price := float32(15.99)
	database, _ := db.ConnectDB()
	productRepository := repository.ProductRepositoryDb{Db: database}
	productUseCase := usecase.ProductUseCase{ProductRepository: productRepository}

	product, _ := model.NewProduct(name, description, price)
	productUseCase.ProductRepository.AddProduct(product)
	products := productUseCase.ProductRepository.FindProducts()

	require.NotNil(t, products)
	require.Equal(t, len(*products), 1)
}
