package db

import "tutorials/009-Packages/shopping-no-cycles/models"

func LoadItem(id int) *models.Item {
	return &models.Item{
		Price: 9.001,
	}
}
