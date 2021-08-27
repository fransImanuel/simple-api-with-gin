package car2

import (
	"context"
	"encoding/json"
	"fmt"
	"go-simple-api/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Cart struct {
	Car_ID           *int32
	Name             *string
	Rent_Price_Daily *int32
	Stock            *int32
}

//Read
func ReadCar(c *gin.Context) {
	dbpool := db.ConnectToDB()
	defer dbpool.Close()
	var carts []Cart

	rows, err := dbpool.Query(context.Background(), `select * from "Cars"`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var cart Cart
		err := rows.Scan(&cart.Car_ID, &cart.Name, &cart.Rent_Price_Daily, &cart.Stock)
		if err != nil {
			log.Fatal(err)
		}
		carts = append(carts, cart)
	}
	fmt.Println("TEST")
	c.JSON(http.StatusOK, carts)
}

//Create
func CreateCar(c *gin.Context) {
	var jsonData *Cart
	jsonByte, err := c.GetRawData()
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonByte, &jsonData)
	// name := strings.Trim(&jsonData.Name, `"`)
	// fmt.Println(*jsonData.Name)

	dbpool := db.ConnectToDB()
	defer dbpool.Close()
	var carts []Cart
	id := 0

	sqlQuery := `INSERT INTO "Cars" ("Car_ID","Name","Rent_Price_Daily","Stock")
	VALUES ($1,$2,$3,$4) returning "Car_ID";`
	dbpool.QueryRow(context.Background(), sqlQuery,
		*jsonData.Car_ID,
		*jsonData.Name,
		*jsonData.Rent_Price_Daily,
		*jsonData.Stock).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, carts)
	// c.JSON(http.StatusOK, "OK")
}

//Update
func UpdateCar(c *gin.Context) {
	id := c.Param("id")
	var jsonData *Cart
	jsonByte, err := c.GetRawData()
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonByte, &jsonData)
	// name := strings.Trim(&jsonData.Name, `"`)
	// fmt.Println(*jsonData.Name)

	dbpool := db.ConnectToDB()
	defer dbpool.Close()

	sqlQuery := `
		UPDATE "Cars"
		SET "Name" = $1,
		"Rent_Price_Daily" = $2,
		"Stock" = $3
		WHERE "Car_ID" = $4;`
	commandTag, err := dbpool.Exec(context.Background(), sqlQuery,
		*jsonData.Name,
		*jsonData.Rent_Price_Daily,
		*jsonData.Stock,
		id)
	if err != nil {
		panic(err)
	}
	if commandTag.RowsAffected() >= 1 {
		c.JSON(http.StatusOK, "Update Success")
	}
}

//Delete
func DeleteCar(c *gin.Context) {
	id := c.Param("id")
	dbpool := db.ConnectToDB()
	defer dbpool.Close()

	sqlQuery := `
		DELETE FROM "Cars"
		WHERE "Car_ID" = $1;
	`
	commandTag, err := dbpool.Exec(context.Background(), sqlQuery, id)
	if err != nil {
		panic(err)
	}

	if commandTag.RowsAffected() >= 1 {
		c.JSON(http.StatusOK, "Success")
	}
}
