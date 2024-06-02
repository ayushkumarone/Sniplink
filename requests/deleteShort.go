package requests

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ayushkumarone/Sniplinks/pkg/links"
	"github.com/gin-gonic/gin"
)

func DeleteShortByID(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var newLink links.Link // getting apikey and short

	if err := c.BindJSON(&newLink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error"})
		return
	}

	apikey := newLink.Apikey

	if apikey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "This feature is only available for accounts."})
		return
	}
	var email string
	if err := db.QueryRow(fmt.Sprintf("SELECT Email FROM email_apikeys where Api_key='%v';", apikey)).Scan(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The API Key is invalid."})
		return
	}

	var count int
	if err1 := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM shorturls WHERE CreatedBy = '%v' AND Short = '%v';", email, id)).Scan(&count); err1 != nil {
		return
	}
	if count == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No such shotened link was created by you."})
		return
	}

	query := fmt.Sprintf("DELETE FROM shorturls WHERE CreatedBy = '%v' AND Short = '%v';", email, id)
	_, err := db.Exec(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	updateUsers := fmt.Sprintf("UPDATE shorturlusers SET Numberoflinks = Numberoflinks - 1 WHERE Email='%v';", email)
	_, err1 := db.Exec(updateUsers) // Performing the query.link/a1
	if err1 != nil {
		// fmt.Println("Error at performing deletingUnused: ", err)
		fmt.Println("Error reducing the number of available links.")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "The shortened link deleted successfully."})
}
