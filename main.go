package main

import (
	"github.com/hasyimzii/go_cashier_cli/controllers"
	"github.com/hasyimzii/go_cashier_cli/views"
)

func init() {
	controllers.ClearScreen()
}

func main() {
	views.MainMenu()
}
