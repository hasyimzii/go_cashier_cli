package views

import (
	"fmt"

	"github.com/hasyimzii/go_cashier_cli/controllers"
)

func MainMenu() {
	fmt.Println("MENU")
	fmt.Println("[1] Transaction")
	fmt.Println("[2] Products")
	fmt.Println("[0] Exit")

	controllers.UserInput("Type menu number: ", &controllers.Input)

	switch controllers.Input {
	case "2":
		ProductMenu()
	case "0":
		ExitMenu()
	default:
		controllers.WrongInput()
		MainMenu()
	}
}
