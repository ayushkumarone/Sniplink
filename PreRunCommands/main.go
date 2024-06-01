package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
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

	fmt.Println("Execution of Tables creation command completed")
}
