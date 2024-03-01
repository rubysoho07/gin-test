package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// struct people
type People struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	State string `json:"state"`
}

func ConnectDB() {
	var err error
	db, err = sql.Open("mysql", "root:example@tcp(localhost:3306)/test_data")
	if err != nil {
		log.Fatal(err)
	}
}

func GetData(c *gin.Context) {

	id := c.Param("id")

	query, err := db.Query("SELECT id, name, age, state FROM people WHERE id = ?;", id)

	if err != nil {
		log.Fatal(err)
	}

	// convert query result to struct people
	var p People
	for query.Next() {
		err := query.Scan(&p.ID, &p.Name, &p.Age, &p.State)
		if err != nil {
			log.Fatal(err)
		}
	}

	c.JSON(http.StatusOK, p)

	defer query.Close()

}

func InsertData(c *gin.Context) {
	// Insert a row to MySQL database

	// request body in JSON to struct people
	var p People
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	query, err := db.Exec("INSERT INTO people (`name`, `age`, `state`) VALUES (?, ?, ?);", &p.Name, &p.Age, &p.State)

	if err != nil {
		log.Fatal(err)
	}

	// Get insert query result
	id, err := query.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusOK, "Inserted data with id: %d", id)
}

func DeleteData(c *gin.Context) {
	// Delete a row from MySQL database
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Fatal(err)
	}

	delete, err := db.Exec("DELETE FROM people WHERE id = ?;", id)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := delete.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, rows)
}

func UpdateData(c *gin.Context) {
	// Update a row in MySQL database
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Fatal(err)
	}

	// request body in JSON to struct people
	var p People
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	update, err := db.Exec("UPDATE people SET name = ?, age = ?, state = ? WHERE id = ?;", &p.Name, &p.Age, &p.State, id)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := update.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, rows)
}
