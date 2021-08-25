package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectToDB() (Dbpool *pgxpool.Pool) {

	Dbpool, err := pgxpool.Connect(context.Background(), "postgres://postgres:d3v3l0p8015@35.187.248.198:5432/CenturyNet?sslmode=disable&pool_max_conns=10")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// defer Dbpool.Close()

	return Dbpool

	// config, err := pgxpool.ParseConfig(os.Getenv("postgres://postgres:d3v3l0p8015@35.187.248.198:5432/CenturyNet?sslmode=disable&pool_max_conns=10"))
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
	// 	// do something with every new connection
	// }

	// pool, err := pgxpool.ConnectConfig(context.Background(), config)

	// var greeting string
	// err = Dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(greeting)

}
