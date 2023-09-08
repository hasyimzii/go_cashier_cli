package controllers

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hasyimzii/go_cashier_cli/models"
)

func GetProduct() {
	SortProductAsc()

	for _, product := range models.Products {
		fmt.Println(product.Code+"|", product.Name, "(Rp "+strconv.FormatInt(product.Price, 10)+")", "Stock:", product.Stock)
	}
}

func FindProductIndex(code string) (int, error) {
	for index, product := range models.Products {
		if product.Code == code {
			return index, nil
		}
	}
	return -100, errors.New("[Product code \"" + code + "\" not found!]")
}

func PostProduct() {

}

func PutProduct() {

}

func DeleteProduct(index int) {
	models.Products = append(models.Products[:index], models.Products[index+1:]...)
}
