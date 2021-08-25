package car

import (
	"context"
	"fmt"
	"go-simple-api/db"
	"os"

	"github.com/gin-gonic/gin"
)

var Carvar = "variable"

//Create
func CreateCar(c *gin.Context) {
	dbpool := db.ConnectToDB()

	var greeting string
	err := dbpool.QueryRow(context.Background(), "select 'This is From Create Car'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}

//Read
//Update
//Delete
