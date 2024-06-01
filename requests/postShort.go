package requests

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ayushkumarone/Sniplinks/pkg/links"

	"github.com/gin-gonic/gin"
)

func PostShort(c *gin.Context, db *sql.DB) {
	var newLink links.Link

	if err := c.BindJSON(&newLink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error"})
		return
	}

	query := fmt.Sprintf("SELECT Short FROM shorturls where Short='%v';", newLink.Short)

	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	var duplicate int
	for rows.Next() {
		var short string
		if err := rows.Scan(&short); err != nil {
			fmt.Print(err)
			return
		}
		duplicate++
	}
	if duplicate > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Short url already exist"})
		return
	}

	queryInsert := fmt.Sprintf("INSERT INTO shorturls (Short, Url) VALUES ('%v', '%v');", newLink.Short, newLink.Url)
	_, err2 := db.Exec(queryInsert)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Shortened link created."})
}
