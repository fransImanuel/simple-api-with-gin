package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-simple-api/data"
)

func FirstLanding(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"Isi Header": c.Request.Header,
	})
	c.String(http.StatusOK, "\n\nHello This Is First Landing Page")
}


func DummyGet(c *gin.Context){
	c.IndentedJSON(http.StatusOK, data.Albums)
}

func DummyGetByID(c *gin.Context){
	id := c.Param("id")

	for i, v := range data.Albums {
		if v.ID == id {
			c.IndentedJSON(http.StatusOK, data.Albums[i].Rating)
			return
		}
	}
}

func TestURLQuery(c *gin.Context){
	name := c.Request.URL.Query().Get("name")
	city := c.Request.URL.Query().Get("city")
	// c.IndentedJSON(http.StatusOK, c.Request.URL.Query())
	c.IndentedJSON(http.StatusOK, map[string]string{"name" : name, "city": city})
}

func DummyPost(c *gin.Context){
	var newAlbum data.AlbumStruct

	err := c.BindJSON(&newAlbum);
	if err != nil{
		return
	}

	data.Albums = append(data.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func DummyPut(c *gin.Context){
	//this way I got fron stackoverflow
	// buf := make([]byte, 1024)
	// num,_ := c.Request.Body.Read(buf)
	// reqBody := string(buf[0:num])
	// fmt.Println(reqBody)
	// var newAlbum album
	id := c.Param("id")

	//get body data and insert it on newAlbum
	var newAlbum data.AlbumStruct
	err := c.BindJSON(&newAlbum);
	if err != nil{
		return
	}

	// fmt.Println(newAlbum.ID)
	// fmt.Println(newAlbum.Title)
	// fmt.Println(newAlbum.Artist)
	// fmt.Println(newAlbum.Price)
	// for i, a := range albums {
	// 	if (a.ID == id) {
	// 		temp := &albums[i]
	// 		temp.ID = newAlbum.ID
	// 		temp.Title = newAlbum.Title
	// 		temp.Artist = newAlbum.Artist
	// 		temp.Price = newAlbum.Price
	// 	}
	// 	fmt.Println("Dalam for")
	// 	fmt.Println(a)
	// }

	for i, a := range data.Albums {
		if (a.ID == id) {
			data.Albums[i].ID = newAlbum.ID
			data.Albums[i].Title = newAlbum.Title
			data.Albums[i].Artist = newAlbum.Artist
			data.Albums[i].Price = newAlbum.Price
		}
		// fmt.Println("Dalam for")
		// fmt.Println(a)
	}
	
	// for _, a := range albums {
	// 	if (a.ID == id) {
	// 		a.ID = newAlbum.ID
	// 		a.Title = newAlbum.Title
	// 		a.Artist = newAlbum.Artist
	// 		a.Price = newAlbum.Price
	// 	}
	// 	fmt.Println("Dalam for")
	// 	fmt.Println(a)
	// }
	// fmt.Println("Luar for")
	// fmt.Println(data.Albums)
	c.IndentedJSON(http.StatusOK, data.Albums)
}

func DummyDelete(c *gin.Context){
	id := c.Param("id")
	// var newSlice []album
	for i, v := range data.Albums {
		if v.ID == id {
			data.Albums = append(data.Albums[:i], data.Albums[i+1:]...)
		}
	}

	// fmt.Println(newSlice)
	c.IndentedJSON(http.StatusOK, data.Albums)
}