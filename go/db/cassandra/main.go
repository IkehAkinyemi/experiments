package main

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

func main() {
	// Create a cluster configuration
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "mykeyspace"
	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer session.Close()

	// Basic query
	var id gocql.UUID
	var name string
	if err := session.Query("SELECT id, name FROM users WHERE id = ?", gocql.UUIDFromTime(time.Now())).Scan(&id, &name); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User name:", name)

	// Prepared statement
	stmt, err := session.Prepare("INSERT INTO users (id, name, email) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := stmt.Exec(gocql.TimeUUID(), "John Doe", "johndoe@example.com"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User created successfully")

	// Transaction
	if err := session.Query("BEGIN").Exec(); err != nil {
		fmt.Println(err)
		return
	}
	if err := session.Query("UPDATE account SET balance = balance - 10 WHERE user_id = ?", 123).Exec(); err != nil {
		fmt.Println(err)
		return
	}
	if err := session.Query("UPDATE account SET balance = balance + 10 WHERE user_id = ?", 456).Exec(); err != nil {
		fmt.Println(err)
		return
	}
	if err := session.Query("COMMIT").Exec(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Transaction completed successfully")
}
