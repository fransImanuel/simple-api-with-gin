package main

import (
	car "go-simple-api/Routes"
	"go-simple-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//API without connection to database
	router.GET("/", handler.FirstLanding)
	router.GET("/DummyGet", handler.DummyGet)
	router.GET("/DummyGetByID/:id/Rating", handler.DummyGetByID)
	router.GET("/TestURLQuery/", handler.TestURLQuery) // ?name=Jhon&city=Jakarta
	router.POST("/DummyPost", handler.DummyPost)
	router.PUT("/DummyPut/:id", handler.DummyPut)
	router.DELETE("/DummyDelete/:id", handler.DummyDelete)

	//API with CRUD to Databases(postgres)
	router.GET("/car", car.CreateCar)

	router.Run("localhost:8080")
}

// kalau mau post/update format bodynya seperti ini
// {
//     "id" : "5",
//     "title" : "Harry Potter 5",
//     "artist" : "J.K Rowling5",
//     "price" : 50000,
// "Rating" : {
// 	"Alibaba" : "5/5"
// 	"Google" : "5/5"
// }
// }
