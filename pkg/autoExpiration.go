package pkg

import (
	"database/sql"
	"fmt"
	"time"
)

func RemoveLinks(db *sql.DB) {
	expiration := 180000
	time.Sleep(time.Duration(expiration))
	deletingUnused := fmt.Sprintf("DELETE FROM shorturls WHERE TIMESTAMPDIFF(SECOND, LastHit, CURRENT_TIMESTAMP) > '%v';", expiration)
	_, err2 := db.Exec(deletingUnused) // Performing the query.
	if err2 != nil {
		// fmt.Println("Error at performing deletingUnused: ", err2)
		fmt.Println("Error deleting the unused shortened links.")
		return
	}
	RemoveLinks(db)
}

func RemoveApikey(db *sql.DB) {
	expiration := 120
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
