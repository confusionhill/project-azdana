package application

import (
	"fmt"
	"log"
	"os"

	"com.github/confusionhill-aqw-ps/application/consumer"
	"com.github/confusionhill-aqw-ps/internal/config"
)

func Setup(cfg *config.Config, rsc *consumer.Resources) {
	// db. execute sql query from folder private/db_template.sql
	sqlBytes, err := os.ReadFile("private/aqw_db2.sql")
	if err != nil {
		log.Fatalf("failed to read SQL file: %v", err)
	}

	sqlQuery := string(sqlBytes)

	// Execute the SQL query
	_, err = rsc.Db.Exec(sqlQuery)
	if err != nil {
		log.Fatalf("failed to execute SQL query: %v", err)
	}

	fmt.Println("SQL query executed successfully")
}
