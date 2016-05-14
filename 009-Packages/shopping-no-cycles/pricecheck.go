package shopping-no-cycles

import "tutorials/009-Packages/shopping-no-cycles/db"

func PriceCheck(itemId int) (float64, bool) {
	item := db.LoadItem(itemId)
	if item == nil {
		return 0, false
	}
	return item.Price, true
}
