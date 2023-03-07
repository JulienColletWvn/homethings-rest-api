package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	_ "github.com/lib/pq"
)

var Database *pgxpool.Pool

func Connect() error {
	var err error
	Database, err = pgxpool.Connect(context.Background(), fmt.Sprintf("postgresql://%v:%v@%v:5432/%v?sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_URL"), os.Getenv("POSTGRES_DB")))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return nil
}
