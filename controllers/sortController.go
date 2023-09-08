package controllers

import (
	"sort"

	"github.com/hasyimzii/go_cashier_cli/models"
)

func SortProductAsc() {
	sort.Slice(models.Products, func(i, j int) bool {
		return models.Products[i].Code < models.Products[j].Code
	})
}
