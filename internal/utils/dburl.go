package utils

import (
	"fmt"

	"github.com/roman-mik/horus-heresy-tactica/internal/config"
	"github.com/roman-mik/horus-heresy-tactica/internal/constants"
)

// urlExample := "postgres://username:password@host:port/database_name"

func GetDBUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.GetEnvVar(constants.DB_USER),
		config.GetEnvVar(constants.DB_PASSWORD),
		config.GetEnvVar(constants.DB_HOST),
		config.GetEnvVar(constants.DB_PORT),
		config.GetEnvVar(constants.DB_NAME))
}
