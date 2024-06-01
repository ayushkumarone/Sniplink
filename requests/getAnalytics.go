package requests

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ayushkumarone/Sniplinks/pkg/links"
	"github.com/gin-gonic/gin"
)

func GetUserAnalytics(c *gin.Context, db *sql.DB) {
	var newLink links.Link // getting apikey and short
	var jsonlink links.JsonLink

	var links []links.JsonLink

	if err := c.BindJSON(&newLink); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error"})
		return
	}

	apikey := newLink.Apikey

	if apikey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "This feature is only available for accounts."})
	}
	var email string
	if err := db.QueryRow(fmt.Sprintf("SELECT Email FROM email_apikeys where Api_key='%v';", apikey)).Scan(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The API Key is invalid."})
		return
	}

	query := fmt.Sprintf("SELECT Short, Url, HitCount, LastHit, CreatedBy FROM shorturls WHERE CreatedBy = '%s'", email)

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		err := rows.Scan(&jsonlink.Short, &jsonlink.Url, &jsonlink.HitCount, &jsonlink.LastHit, &jsonlink.CreatedBy)
		if err != nil {
			fmt.Println(err)
			return
		}
		links = append(links, jsonlink)
		count++
	}

	jsonData, err := json.Marshal(links)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Links count": count, "Links": json.RawMessage(jsonData)})
}
