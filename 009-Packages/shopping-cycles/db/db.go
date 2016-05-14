package db

import "tutorials/009-Packages/shopping-cycles"

func LoadItem(id int) *shopping.Item {
	return &shopping.Item{
		Price: 9.001,
	}
}
