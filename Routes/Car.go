package car

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go-simple-api/db"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type NullString struct {
	sql.NullString
}

type NullInt32 struct {
	sql.NullInt32
}

func (s NullInt32) MarshalJSON() ([]byte, error) {
	fmt.Println("Masuk ke MarshalJSON() NullInt32")
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.Int32)
}

func (s *NullInt32) UnmarshalJSON(data []byte) error {
	fmt.Println("Masuk ke UnmarshalJSON() NullInt32")
	if string(data) == "null" {
		s.Int32, s.Valid = 0, false
		return nil
	}
	s.Int32, s.Valid = s.Int32, true
	return nil
}

func (s NullString) MarshalJSON() ([]byte, error) {
	fmt.Println("Masuk ke MarshalJSON() NullString")
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	fmt.Println("Masuk ke UnmarshalJSON() NullString")
	if string(data) == "null" {
		s.String, s.Valid = "", false
		return nil
	}
	s.String, s.Valid = string(data), true
	return nil
}

type Cart struct {
	Car_ID           NullString
	Name             NullString
	Rent_Price_Daily NullString
	Stock            NullInt32
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
	dbpool := db.ConnectToDB()
	defer dbpool.Close()

	var jsonData Cart
	jsonByte, err := c.GetRawData()
	if err != nil {
		panic(err)
	}

	json.Unmarshal(jsonByte, &jsonData)
	name := strings.Trim(jsonData.Name.String, `"`)
	// fmt.Println(strings.Trim(jsonData.Name.String, `"`))

	id := 0
	sqlQuery := `
		INSERT INTO "Cars" ("Car_ID","Name","Rent_Price_Daily","Stock")
		VALUES ($1,$2,$3,$4) returning "Car_ID";
	`
	dbpool.QueryRow(context.Background(), sqlQuery,
		jsonData.Car_ID.String,
		name,
		jsonData.Rent_Price_Daily.String,
		jsonData.Stock.Int32).Scan(&id)
	if err != nil {
		log.Fatal(err, " - Line 104")
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

//Update
func UpdateCar(c *gin.Context) {
	id := c.Param("id")
	dbpool := db.ConnectToDB()
	defer dbpool.Close()

	var jsonData Cart
	jsonByte, err := c.GetRawData()
	if err != nil {
		panic(err)
	}

	json.Unmarshal(jsonByte, &jsonData)
	name := strings.Trim(jsonData.Name.String, `"`)

	sqlQuery := `
		UPDATE "Cars"
		SET "Name" = $1,
		"Rent_Price_Daily" = $2,
		"Stock" = $3
		WHERE "Car_ID" = $4;
	`
	commandTag, err := dbpool.Exec(context.Background(), sqlQuery,
		name,
		jsonData.Rent_Price_Daily.String,
		jsonData.Stock.Int32,
		id)
	if err != nil {
		panic(err)
	}
	if commandTag.RowsAffected() != 1 {
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
