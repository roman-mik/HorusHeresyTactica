package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/roman-mik/horus-heresy-tactica/internal/utils"
)

func GetConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), utils.GetDBUrl())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func CloseConnection(conn *pgx.Conn) {
	conn.Close(context.Background())
}
