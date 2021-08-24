package data



// album represents data about a record album.
type Rating struct{
	Alibaba string
	Google string
}
type AlbumStruct struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
	Rating Rating
}

// albums slice to seed record album data.
var Albums = []AlbumStruct{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99, Rating : Rating{ Alibaba : "3/5", Google : "3.5/5" }},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99 ,Rating : Rating{ Alibaba : "2/5", Google : "2.5/5" }},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99,Rating : Rating{ Alibaba : "5/5", Google : "5/5" }},
}
