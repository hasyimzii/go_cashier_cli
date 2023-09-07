package menus

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/hasyimzii/go_cashier_cli/models"
)

func MainMenu() {
	var input string

	fmt.Println("MENU\n",
		"1. Transaction\n",
		"2. Products\n",
		"0. Exit")

	UserInput("Type menu number: ", &input)

	switch input {
	case "2":
		sort.Slice(models.Products, func(i, j int) bool {
			return models.Products[i].Code < models.Products[j].Code
		})

		for _, product := range models.Products {
			fmt.Println(product.Code, "-", product.Name, "(Rp "+strconv.FormatInt(product.Price, 10)+")")
		}
	case "0":
		UserInput("Are you sure want to exit? [y/n]: ", &input)

		switch strings.ToLower(input) {
		case "y":
			fmt.Println("Good bye!")
		case "n":
			MainMenu()
		default:
			WrongInput()
		}
	default:
		WrongInput()
	}
}
