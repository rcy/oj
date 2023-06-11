package db

import (
	"log"
	"os"
)

const schemaFile = "./db/schema.sql"

func MustDump() {
	rows, err := DB.Queryx("SELECT sql FROM sqlite_master")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Create a file to store the schema
	file, err := os.Create(schemaFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Iterate over the rows and write the SQL statements to the file
	for rows.Next() {
		var sqlStatement string
		err := rows.Scan(&sqlStatement)
		if err != nil {
			panic(err)
		}
		_, err = file.WriteString(sqlStatement + ";\n")
		if err != nil {
			panic(err)
		}
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	log.Printf("Dumped schema to %s", schemaFile)
}
