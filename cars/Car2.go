package car2

import (
	"context"
	"fmt"
	"go-simple-api/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

	
type Cart struct{
	Car_ID	*int32
	Name 	*string
	Rent_Price_Daily *int32
	Stock 	*int32
}

//Read
func ReadCar(c *gin.Context) {
	dbpool := db.ConnectToDB()
	defer dbpool.Close()
	var carts []Cart

	
	rows, err := dbpool.Query(context.Background(),`select * from "Cars"`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	for rows.Next() {
		var cart Cart
		err:=rows.Scan(&cart.Car_ID, &cart.Name, &cart.Rent_Price_Daily, &cart.Stock)
		if err!=nil {
			log.Fatal(err)
		}
		carts = append(carts,cart)
	}
	fmt.Println("TEST")
	c.JSON(http.StatusOK, carts)
}

//Create
//Update
//Delete
