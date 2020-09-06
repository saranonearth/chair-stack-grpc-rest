package store

var itemCount int

// Item type for a garment- [id, garment ]
type Item struct {
	ID      int    `json:"id"`
	Garment string `json:"garment"`
}

// ItemList - list of items
type ItemList []Item

var list ItemList

func init() {
	itemCount = 0
	list = make(ItemList, 0)

}

// AddToList to add to store list
func AddToList(garment string) ItemList {
	itemCount++
	newGarment := Item{
		ID:      itemCount,
		Garment: garment,
	}

	list = append(list, newGarment)

	return list
}

// GetAllItems - return all list items
func GetAllItems() ItemList {
	return list
}

//Clear the list
func Clear() {
	itemCount = 0
	list = make(ItemList, 0)
}
