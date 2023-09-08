package views

import (
	"fmt"
	"strconv"

	"github.com/hasyimzii/go_cashier_cli/controllers"
	"github.com/hasyimzii/go_cashier_cli/models"
)

func ProductMenu() {
	fmt.Println("[1] Product List")
	fmt.Println("[2] Add Product")
	fmt.Println("[3] Edit Product")
	fmt.Println("[4] Remove Product")
	fmt.Println("[0] Back")

	controllers.UserInput("Type menu number: ", &controllers.Input)

	switch controllers.Input {
	case "1":
		ProductList()
	case "2":
		AddProduct()
	case "3":
		EditProduct()
	case "4":
		RemoveProduct()
	case "0":
		MainMenu()
	default:
		controllers.WrongInput()
		ProductMenu()
	}
}

func ProductList() {
	controllers.GetProduct()

	controllers.UserInput("[Press enter to go back] ", &controllers.Input)
}

func AddProduct() {
	// TODO: Bikin add
}

func EditProduct() {
	controllers.GetProduct()

	controllers.UserInput("Type product code to edit: ", &controllers.Input)

	index, err := controllers.FindProductIndex(controllers.Input)
	if err != nil {
		fmt.Println(err)
		ProductMenu()
	}

	var name string = models.Products[index].Name
	var price string = strconv.FormatInt(models.Products[index].Price, 10)
	var stock string = strconv.FormatInt(models.Products[index].Stock, 10)

	var nameInput string
	var priceInput int64
	var stockInput int64

	// FIXME: update input benerin
	fmt.Println("Type on the field you want to edit (Press enter to skip)")
	fmt.Printf("Product name (%v): ", name)
	fmt.Scan(nameInput)
	fmt.Printf("Price (Rp %v): ", price)
	fmt.Scan(priceInput)
	fmt.Printf("Stock (%v): ", stock)
	fmt.Scan(stockInput)

	fmt.Println(models.Product{
		Name:  nameInput,
		Price: priceInput,
		Stock: stockInput,
	})
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
}
