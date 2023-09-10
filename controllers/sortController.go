package controllers

import (
	"sort"

	"github.com/hasyimzii/go_cashier_cli/models"
)

func SortProductAsc() {
	sort.Slice(models.ProductList, func(i, j int) bool {
		return models.ProductList[i].Code < models.ProductList[j].Code
	})
}
