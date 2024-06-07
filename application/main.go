package main

import (
	"database/sql"
	"fmt"
	"log"
	"main/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
	// _ "embed"
	// "fmt"
	// "hash/crc32"
	// "strconv"
	// "strings"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// -------------------------------------------------------
// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// ----------------------------------------

func main() {

	cfg := config.Get()

	cfg.DB.Host = "localhost"
	cfg.DB.Port = "8080"
	cfg.DB.Username = "postgres"
	cfg.DB.Password = ""
	cfg.DB.Name = "MY_DB"

	connectionSettings := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.Name)

	DBConnection, err := sql.Open("postgres", connectionSettings)

	if err != nil {
		log.Fatal("MY_DB database connect fails")
	}
	err = DBConnection.Ping()

	if err != nil {
		log.Fatal("MY_DB database ping fail")
	}

	fmt.Println("Successfully connected!")

	DBConnection.Ping()

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("albums", postAlbums)

	router.Run("localhost:8080")

}

// package main

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/lib/pq"
// )

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "new_pass"
// 	dbname   = "my_test"
// )

// func main() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	CheckFatal(err)
// 	defer db.Close()

// 	err = db.Ping()
// 	CheckFatal(err)

// 	fmt.Println("Successfully connected!")

// 	CreateTableQuery(db)
// 	InsertIntoTableQuery(db)
// 	SelectTableQuery(db)
// }

// func CreateTableQuery(db *sql.DB) {
// 	q := `CREATE TABLE IF NOT EXISTS my_users (
// 		id SERIAL PRIMARY KEY,
// 		name VARCHAR(100) NOT NULL
// 		)`

// 	_, err := db.Exec(q)
// 	CheckFatal(err)
// 	fmt.Println("Successfully created table!")
// }

// func InsertIntoTableQuery(db *sql.DB) {
// 	q := `INSERT INTO my_users (id, name) values(1, 'Jonh')`
// 	_, err := db.Exec(q)
// 	CheckFatal(err)
// 	fmt.Println("Successfully inserted data into table!")
// }

// func SelectTableQuery(db *sql.DB) {
// 	q := `SELECT * FROM my_users`
// 	_, err := db.Exec(q)
// 	CheckFatal(err)
// 	fmt.Println("Successfully selected data!")
// }
// func CheckFatal(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
