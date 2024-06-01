package requests

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ayushkumarone/Sniplinks/pkg"
	"github.com/ayushkumarone/Sniplinks/pkg/users"
	"github.com/gin-gonic/gin"
)

func duplicateUserCheck(db *sql.DB, email string) int {
	var count int

	if err := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM shorturlusers where Email='%v';", email)).Scan(&count); err != nil {
		return 1
	}

	return count
}

func RegisterUser(c *gin.Context, db *sql.DB) {
	var newUser users.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error"})
		return
	}

	usercount := duplicateUserCheck(db, newUser.Email)
	if usercount > 0 {
		c.JSON(http.StatusConflict, gin.H{"message": "User with this email already exists. Please login to your account."})
		return
	}

	hashString := pkg.MD5HashGenerator(newUser.Password)

	queryInsert := fmt.Sprintf("INSERT INTO shorturlusers (Name, Email, Passwordhash) VALUES ('%v', '%v', '%v');", newUser.Name, newUser.Email, hashString)
	_, err2 := db.Exec(queryInsert)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}
