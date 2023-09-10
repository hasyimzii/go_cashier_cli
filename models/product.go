package models

type Product struct {
	Code  string
	Name  string
	Price int
	Stock int
}

var ProductList = []Product{
	{
		Code:  "A01",
		Name:  "Fitness Gizmo",
		Price: 986900,
		Stock: 15,
	},
	{
		Code:  "A02",
		Name:  "Swimming Contraption",
		Price: 429500,
		Stock: 11,
	},
	{
		Code:  "B01",
		Name:  "Cleaning Gadget",
		Price: 151990,
		Stock: 20,
	},
	{
		Code:  "C01",
		Name:  "Game Item",
		Price: 43000,
		Stock: 12,
	},
	{
		Code:  "A03",
		Name:  "Fitness Component",
		Price: 139000,
		Stock: 22,
	},
	{
		Code:  "B02",
		Name:  "Gardening Contraption",
		Price: 138950,
		Stock: 8,
	},
	{
		Code:  "C02",
		Name:  "Phone Attachment",
		Price: 13950,
		Stock: 17,
	},
	{
		Code:  "C03",
		Name:  "Baby Whatchamacallit",
		Price: 62950,
		Stock: 5,
	},
	{
		Code:  "C04",
		Name:  "Baby Gizmo",
		Price: 125950,
		Stock: 19,
	},
	{
		Code:  "B03",
		Name:  "Furniture Tool",
		Price: 94000,
		Stock: 25,
	},
}
