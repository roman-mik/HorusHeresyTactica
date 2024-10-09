package models

type Legion struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Legions = []Legion{
	{ID: 1, Name: "Dark Angels"},
	{ID: 3, Name: "Emperor's Children"},
	{ID: 4, Name: "Iron Warriors"},
	{ID: 5, Name: "White Scars"},
	{ID: 6, Name: "Space Wolves"},
	{ID: 7, Name: "Imperial Fists"},
	{ID: 8, Name: "Night Lords"},
	{ID: 9, Name: "Blood Angels"},
	{ID: 10, Name: "Iron Hands"},
	{ID: 11, Name: "World Eaters"},
	{ID: 13, Name: "Ultramarines"},
	{ID: 14, Name: "Death Guard"},
	{ID: 15, Name: "Thousand Sons"},
	{ID: 16, Name: "Sons of Horus"},
	{ID: 17, Name: "Word Bearers"},
	{ID: 18, Name: "Salamanders"},
}

// GetLegions returns the list of legions
func GetLegions() []Legion {
	return Legions
}

func GetLegionByID(id int) *Legion {
	for _, legion := range Legions {
		if legion.ID == id {
			return &legion
		}
	}
	return nil
}
