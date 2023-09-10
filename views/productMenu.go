package views

import (
	"fmt"
	"strconv"

	"github.com/hasyimzii/go_cashier_cli/controllers"
	"github.com/hasyimzii/go_cashier_cli/models"
)

func ProductMenu() {
	fmt.Println("[1] All Products")
	fmt.Println("[2] Create Product")
	fmt.Println("[3] Update Product")
	fmt.Println("[4] Remove Product")
	fmt.Println("[0] Back")

	controllers.UserInput("Type menu number: ", &controllers.Input)

	switch controllers.Input {
	case "1":
		AllProducts()
	case "2":
		CreateProduct()
	case "3":
		UpdateProduct()
	case "4":
		RemoveProduct()
	case "0":
		MainMenu()
	default:
		controllers.WrongInput()
		ProductMenu()
	}
}

func AllProducts() {
	controllers.GetProduct()

	controllers.UserInput("[Press enter to go back] ", &controllers.Input)

	ProductMenu()
}

func CreateProduct() {
	var code, name, price, stock string

	fmt.Println("Type on every field")
	fmt.Print("Product code: ")
	fmt.Scanln(&code)
	fmt.Print("Product name: ")
	fmt.Scanln(&name)
	fmt.Print("Product price: ")
	fmt.Scanln(&price)
	fmt.Print("Product stock: ")
	fmt.Scanln(&stock)

	controllers.ClearScreen()

	if code == "" || name == "" || price == "" || stock == "" {
		fmt.Println("[Alert: Field cannot be empty!]")
	} else {
		err := controllers.PostProduct(code, name, price, stock)
		if err != nil {
			fmt.Println(err)
		}
	}

	ProductMenu()
}

func UpdateProduct() {
	controllers.GetProduct()

	controllers.UserInput("Type product code to edit: ", &controllers.Input)

	var index int
	var err error
	index, err = controllers.FindProductIndex(controllers.Input)
	if err != nil {
		fmt.Println(err)
		ProductMenu()
		return
	}

	var product models.Product = controllers.FindProductModel(index)
	var name, price, stock string

	fmt.Println("Type on the field you want to edit (Press enter to skip)")
	fmt.Print("Product name (" + product.Name + "): ")
	fmt.Scanln(&name)
	fmt.Printf("Price (Rp " + strconv.Itoa(product.Price) + "): ")
	fmt.Scanln(&price)
	fmt.Printf("Stock (" + strconv.Itoa(product.Stock) + "): ")
	fmt.Scanln(&stock)

	controllers.ClearScreen()

	err = controllers.PutProduct(index, name, price, stock)
	if err != nil {
		fmt.Println(err)
	}

	ProductMenu()
}

func RemoveProduct() {
	controllers.GetProduct()

	controllers.UserInput("Type product code to remove: ", &controllers.Input)

	index, err := controllers.FindProductIndex(controllers.Input)
	if err != nil {
		fmt.Println(err)
		ProductMenu()
	}

	controllers.DeleteProduct(index)

	ProductMenu()
}
