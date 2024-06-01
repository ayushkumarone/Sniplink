package requests

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/ayushkumarone/Sniplinks/pkg/users"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context, db *sql.DB) {
	var newUser users.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error"})
		return
	}

	query := fmt.Sprintf("SELECT Email FROM shorturlusers where Email='%v';", newUser.Email)

	rows, err := db.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	var usercount int
	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			fmt.Print(err)
			return
		}
		usercount++
	}
	if usercount > 0 {
		c.JSON(http.StatusConflict, gin.H{"message": "User with this email already exists. Please login to your account."})
		return
	}

	data := []byte(newUser.Password)
	hash := md5.Sum(data)
	hashString := hex.EncodeToString(hash[:])

	queryInsert := fmt.Sprintf("INSERT INTO shorturlusers (Name, Email, Passwordhash) VALUES ('%v', '%v', '%v');", newUser.Name, newUser.Email, hashString)
	_, err2 := db.Exec(queryInsert)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}
