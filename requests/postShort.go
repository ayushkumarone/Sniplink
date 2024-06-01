package requests

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ayushkumarone/Sniplinks/pkg/links"

	"github.com/gin-gonic/gin"
)

func duplicateUrlCheck(db *sql.DB, shorturl string) int {
	// query := fmt.Sprintf("SELECT Short FROM shorturls where Short='%v';", shorturl)
	var count int

	if err := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM shorturls where Short='%v';", shorturl)).Scan(&count); err != nil {
		return 1
	}

	return count
}

func linkInsertionByIP(c *gin.Context, db *sql.DB, newLink links.Link) string {
	ipAddress := c.ClientIP()

	var numberOfLinks int
	queryNumberofLinks := fmt.Sprintf("SELECT Numberoflinks FROM ipaddress WHERE IPaddress = '%v'", ipAddress)
	err := db.QueryRow(queryNumberofLinks).Scan(&numberOfLinks)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("IP address not found.")
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server Error"})
			return "Error"
		}
	}
	if numberOfLinks >= 5 {
		c.JSON(http.StatusLocked, gin.H{"message": "You have reached the maximum possible links that can be created. Please create account or login to create more links."})
		return "Error"
	}

	if err1 := insertLink(c, db, newLink, ipAddress); err1 == "Error" {
		return "Error"
	}

	var count int
	if err := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM ipaddress where IPaddress='%v';", ipAddress)).Scan(&count); err != nil {
		return "Error"
	}
	if count == 0 {
		if _, err := db.Exec(fmt.Sprintf("INSERT INTO ipaddress (IPaddress) VALUE ('%v');", ipAddress)); err != nil {
			return "Error"
		}
	}
	queryIncrementLink := fmt.Sprintf("UPDATE ipaddress SET Numberoflinks = Numberoflinks + 1 WHERE IPaddress = '%v';", ipAddress)
	_, err3 := db.Exec(queryIncrementLink)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server Error"})
		return "Error"
	}
	return ""
}

func insertLink(c *gin.Context, db *sql.DB, newLink links.Link, ipAddress string) string {
	queryInsert := fmt.Sprintf("INSERT INTO shorturls (Short, Url, CreatedBy) VALUES ('%v', '%v', '%v');", newLink.Short, newLink.Url, ipAddress)
	_, err2 := db.Exec(queryInsert)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server Error"})
		return "Error"
	}
	return ""
}

func PostShort(c *gin.Context, db *sql.DB) {
	var newLink links.Link

	if err := c.BindJSON(&newLink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error"})
		return
	}

	duplicate := duplicateUrlCheck(db, newLink.Short)

	if duplicate > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Short url already exist"})
		return
	}

	if newLink.Apikey == "" {
		if err := linkInsertionByIP(c, db, newLink); err == "Error" {
			return
		}
	}
	// var numberOfLinks int

	// queryNumberofLinks := fmt.Sprintf("SELECT Numberoflinks FROM shorturlusers WHERE Email = '%s'", email)
	// err := db.QueryRow(queryNumberofLinks).Scan(&numberOfLinks)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		c.JSON(http.StatusNotFound, gin.H{"message": "User not found."})
	// 		return
	// 	}
	// 	c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server Error"})
	// 	return
	// }

	// if numberOfLinks >= 5 {
	// 	c.JSON(http.StatusLocked, gin.H{"message": "You have reached the maximum possible links that can be created. Please create an account or log in to create more links."})
	// 	return
	// }

	// queryInsert := fmt.Sprintf("INSERT INTO shorturls (Short, Url) VALUES ('%v', '%v');", newLink.Short, newLink.Url)
	// _, err2 := db.Exec(queryInsert)
	// if err2 != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server Error"})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{"message": "Shortened link created."})
}
