package main

import (
	car "go-simple-api/Routes"
	car2 "go-simple-api/Routes2/cars"
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
	//null with interface marshaljson
	router.GET("/car", car.ReadCar)
	router.POST("/car", car.CreateCar)
	router.PUT("/car/:id", car.UpdateCar)
	router.DELETE("/car/:id", car.DeleteCar)

	//null with pointer
	router.GET("/car2", car2.ReadCar)
	router.POST("/car2", car2.CreateCar)
	router.PUT("/car2/:id", car2.UpdateCar)
	router.DELETE("/car2/:id", car2.DeleteCar)

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
