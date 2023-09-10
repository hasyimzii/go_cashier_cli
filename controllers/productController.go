package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hasyimzii/go_cashier_cli/models"
)

func GetProduct() {
	SortProductAsc()

	for _, product := range models.ProductList {
		fmt.Println(product.Code+"|", product.Name, "(Rp "+strconv.Itoa(product.Price)+")", "Stock:", product.Stock)
	}
}

func FindProductIndex(code string) (int, error) {
	for index, product := range models.ProductList {
		if product.Code == code {
			return index, nil
		}
	}

	return -100, errors.New("[Product code not found!]")
}

func FindProductModel(index int) models.Product {
	return models.ProductList[index]
}

func PostProduct(code string, name string, price string, stock string) error {
	// convert to int
	var intPrice, intStock int
	var err error

	intPrice, err = strconv.Atoi(price)
	if err != nil {
		return err
	}
	intStock, err = strconv.Atoi(stock)
	if err != nil {
		return err
	}

	models.ProductList = append(models.ProductList, models.Product{
		Code:  code,
		Name:  name,
		Price: intPrice,
		Stock: intStock,
	})

	fmt.Println("[Create product success]")
	return nil
}

func PutProduct(index int, name string, price string, stock string) error {
	product := FindProductModel(index)

	// check if nil
	if name != "" {
		product.Name = name
	}
	if price != "" {
		intPrice, err := strconv.Atoi(price)
		if err != nil {
			return err
		}
		product.Price = intPrice
	}
	if stock != "" {
		intStock, err := strconv.Atoi(stock)
		if err != nil {
			return err
		}
		product.Stock = intStock
	}

	models.ProductList[index] = product

	fmt.Println("[Update product success]")
	return nil
}

func DeleteProduct(index int) {
	models.ProductList = append(models.ProductList[:index], models.ProductList[index+1:]...)

	fmt.Println("[Delete product success]")
}
