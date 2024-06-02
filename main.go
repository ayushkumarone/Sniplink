package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ayushkumarone/Sniplinks/pkg"
	"github.com/ayushkumarone/Sniplinks/requests"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default() // Defining the router using gin framework.

	// Accessing .env file
	godotenv.Load(".env")
	// Verify the access
	// fmt.Println(os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("Address"), os.Getenv("DB_Name"))

	// ----------------	 START : Defining the configurations  ----------------

	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("Address"),
		DBName:               os.Getenv("DB_Name"),
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	//  ----------------  END : Defining the configurations  ----------------

	//  ----------------  START : Verify connection  ----------------

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println(pingErr)
		return
	}
	fmt.Println("Successfully connected to the MariaDB database!")

	//  ----------------  END : Verify connection  ----------------

	//  ----------------  START : Automated unused URL removal  ----------------

	go pkg.RemoveLinks(db)
	go pkg.RemoveApikey(db)

	//  ----------------  END : Automated unused URL removal  ----------------

	//  ----------------  START : Routes defined here  ----------------

	router.POST("/shorten", func(c *gin.Context) {
		requests.PostShort(c, db)
	})

	router.GET("/link/:id", func(c *gin.Context) {
		requests.GetLinkByID(c, db)
	})

	router.POST("/register", func(c *gin.Context) {
		requests.RegisterUser(c, db)
	})

	router.POST("/login", func(c *gin.Context) {
		requests.LoginUser(c, db)
	})

	router.GET("/analytics", func(c *gin.Context) {
		requests.GetUserAnalytics(c, db)
	})
	//  ----------------  END : Routes defined here  ----------------

	router.Run("localhost:8080") // Running router on localhost port 8080
}
