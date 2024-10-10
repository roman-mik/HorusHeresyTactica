package main

import (
	"fmt"

	"github.com/roman-mik/horus-heresy-tactica/internal/models"
)

func main() {
	var legions = models.GetLegions()
	fmt.Print(legions)
}
