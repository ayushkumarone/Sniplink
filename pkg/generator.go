package pkg

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"math/rand"
)

func MD5HashGenerator(password string) string {
	data := []byte(password)
	hash := md5.Sum(data)
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

func GenerateApiKey(db *sql.DB) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789#"
	length := 32
	api_key := make([]byte, length)
	for {
		for i := range api_key {
			api_key[i] = charset[rand.Intn(len(charset))]
		}
		apiKey := string(api_key)
		// Check if the generated key already exists in the database
		var count int
		if err := db.QueryRow("SELECT COUNT(*) FROM email_apikeys WHERE api_key = ?", apiKey).Scan(&count); err != nil {
			fmt.Println(err)
		}
		if count == 0 {
			// Unique key found, return it
			return apiKey
		}
		// If key exists, generate a new one
	}
}
