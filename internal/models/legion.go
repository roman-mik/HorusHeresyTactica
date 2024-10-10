package models

import (
	"context"
	"fmt"
	"os"

	"github.com/roman-mik/horus-heresy-tactica/internal/database"
)

type Legion struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Number string `json:"number"`
}

// GetLegions returns the list of legions
func GetLegions() []Legion {
	conn := database.GetConnection()

	defer database.CloseConnection(conn)

	var Legions []Legion

	rows, err := conn.Query(context.Background(), "SELECT * FROM legions")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get Legions: %v\n", err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var legion Legion

		rows.Scan(&legion.ID, &legion.Name, &legion.Number)

		Legions = append(Legions, legion)
	}

	if err := rows.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "GetLegions: Unable to read rows from legions table: %v\n", err)
		return nil
	}

	return Legions
}

func GetLegionByID(id int) *Legion {
	conn := database.GetConnection()

	defer database.CloseConnection(conn)

	var legion Legion

	err := conn.QueryRow(
		context.Background(),
		"SELECT * FROM legions WHERE id = $1",
		id,
	).Scan(&legion.ID, &legion.Name, &legion.Number)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get Legion by ID: %v\n", err)
		return nil
	}

	return &legion
}
