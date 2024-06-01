package requests

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ayushkumarone/Sniplinks/pkg"
	"github.com/ayushkumarone/Sniplinks/pkg/users"
	"github.com/gin-gonic/gin"
)

func validateUser(db *sql.DB, inputEmail string) (int, string) {
	query := fmt.Sprintf("SELECT Email, Passwordhash FROM shorturlusers where Email='%v';", inputEmail)

	var email string
	var passwordhash string
	if err := db.QueryRow(query).Scan(&email, &passwordhash); err != nil {
		fmt.Print(err)
		return 0, ""
	}

	usercount := 0
	usercount++

	return usercount, passwordhash
}
func LoginUser(c *gin.Context, db *sql.DB) {
	var newUser users.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error"})
		return
	}

	usercount, passwordhash := validateUser(db, newUser.Email)
	if usercount == 0 {
		c.JSON(http.StatusConflict, gin.H{"message": "User does not exist. Please make a account to login."})
		return
	}

	hashString := pkg.MD5HashGenerator(newUser.Password)

	if passwordhash != hashString { // Check for mismatch in password
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Email and Password incorrect."})
		return
	}

	// Correct password
	queryDelete := fmt.Sprintf("DELETE FROM email_apikeys WHERE Email = '%v';", newUser.Email)
	_, err2 := db.Exec(queryDelete)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server Error"})
		return
	}

	apiKey := pkg.GenerateApiKey(db)

	queryInsert := fmt.Sprintf("INSERT INTO email_apikeys (Email, Api_key) VALUES ('%v', '%v');", newUser.Email, apiKey)
	_, err3 := db.Exec(queryInsert)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged in", "api_key": apiKey})
}
