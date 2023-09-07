package models

type Product struct {
	Code  string
	Name  string
	Price int64
}

var Products = []Product{
	{
		Code:  "A01",
		Name:  "Fitness Gizmo",
		Price: 1869500,
	},
	{
		Code:  "A02",
		Name:  "Swimming Contraption",
		Price: 429500,
	},
	{
		Code:  "B01",
		Name:  "Cleaning Gadget",
		Price: 151990,
	},
	{
		Code:  "C01",
		Name:  "Game Item",
		Price: 43000,
	},
	{
		Code:  "A03",
		Name:  "Fitness Component",
		Price: 139000,
	},
	{
		Code:  "B02",
		Name:  "Gardening Contraption",
		Price: 138950,
	},
	{
		Code:  "C02",
		Name:  "Phone Attachment",
		Price: 13950,
	},
	{
		Code:  "C03",
		Name:  "Baby Whatchamacallit",
		Price: 62950,
	},
	{
		Code:  "C04",
		Name:  "Baby Gizmo",
		Price: 125950,
	},
	{
		Code:  "B03",
		Name:  "Furniture Tool",
		Price: 94000,
	},
}
