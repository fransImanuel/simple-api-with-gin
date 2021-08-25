package car

import (
	"fmt"
	"go-simple-api/db"

	"github.com/gin-gonic/gin"
)

var Carvar = "variable"

//Create
func CreateCar(c *gin.Context) {
	db.ConnectToDB()
	fmt.Println("MASUK")

	// var greeting string
	// err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }
}

//Read
//Update
//Delete
