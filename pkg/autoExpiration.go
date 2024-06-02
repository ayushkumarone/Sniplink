package pkg

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

func RemoveLinks(db *sql.DB) {
	expiration := 300
	time.Sleep(time.Duration(expiration))

	_, err := db.Exec("UPDATE timetable SET time = CURRENT_TIME WHERE place = 0;")
	if err != nil {
		fmt.Println("Error in Setting time query.")
		return
	}

	// Construct the SQL queries
	deletingUnused := fmt.Sprintf("DELETE FROM shorturls WHERE TIMESTAMPDIFF(SECOND, LastHit, (select time from timetable where place = 0)) > %v;", expiration)
	reducingLinks := fmt.Sprintf("SELECT CreatedBy, count(*) FROM shorturls WHERE TIMESTAMPDIFF(SECOND, LastHit, (select time from timetable where place = 0)) > %v GROUP BY CreatedBy;", expiration)

	// Execute the reducingLinks query
	rows, err1 := db.Query(reducingLinks)
	if err1 != nil {
		fmt.Println("Error in reducing links query.", err1)
		return
	}

	// Execute the deletingUnused query
	_, err2 := db.Exec(deletingUnused)
	if err2 != nil {
		fmt.Println("Error in deleting query.")
		return
	}

	for rows.Next() {
		var count int
		var CreatedBy string
		err := rows.Scan(&CreatedBy, &count)
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.Contains(CreatedBy, "@") { // This is an Email
			fmt.Println(count, CreatedBy)
			updateUsers := fmt.Sprintf("UPDATE shorturlusers SET Numberoflinks = Numberoflinks - %v WHERE Email='%v';", count, CreatedBy)
			_, err := db.Exec(updateUsers) // Performing the query.link/a1
			if err != nil {
				// fmt.Println("Error at performing deletingUnused: ", err)
				fmt.Println("Error reducing the number of available links.")
				return
			}
			fmt.Println("User number of links updated.")
			RemoveLinks(db)
			return

		} else {
			fmt.Println(count, CreatedBy)
			updateIPs := fmt.Sprintf("UPDATE ipaddress SET Numberoflinks = Numberoflinks - %v WHERE IPaddress='%v';", count, CreatedBy)
			_, err2 := db.Exec(updateIPs) // Performing the query.
			if err2 != nil {
				// fmt.Println("Error at performing deletingUnused: ", err)
				fmt.Println("Error reducing the number of available links.")
				return
			}
			fmt.Println("IP number of links updated.")

			deletingIPwithzerolinks := "DELETE FROM ipaddress WHERE Numberoflinks = 0;"
			_, err3 := db.Exec(deletingIPwithzerolinks) // Performing the query.
			if err3 != nil {
				// fmt.Println("Error at performing delete: ", err2)
				fmt.Println("Error deleting the unused shortened links.")
				return
			}
		}
	}

	RemoveLinks(db)
}

func RemoveApikey(db *sql.DB) {
	expiration := 730
	time.Sleep(time.Duration(expiration))
	deletingUnused := fmt.Sprintf("DELETE FROM email_apikeys WHERE TIMESTAMPDIFF(SECOND, Creationtime, CURRENT_TIMESTAMP) > '%v';", expiration)
	_, err2 := db.Exec(deletingUnused) // Performing the query.
	if err2 != nil {
		// fmt.Println("Error at performing deletingUnused: ", err2)
		fmt.Println("Error deleting the Api key.")
		return
	}
	RemoveApikey(db)
}
