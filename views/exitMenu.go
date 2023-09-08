package views

import (
	"fmt"
	"strings"

	"github.com/hasyimzii/go_cashier_cli/controllers"
)

func ExitMenu() {
	controllers.UserInput("Are you sure want to exit? [y/n]: ", &controllers.Input)

	switch strings.ToLower(controllers.Input) {
	case "y":
		fmt.Println("Good bye!")
	case "n":
		MainMenu()
	default:
		controllers.WrongInput()
		MainMenu()
	}
}
