package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {
	cluster := gocql.NewCluster("10.8.14.226:9042")
	cluster.Keyspace = "dev"
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "dev",
		Password: "Zh9JYnQYtaCVDVRy",
	}
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("-------------------------------------------------------%v", err)
	}
	defer session.Close()
	ctx := context.Background()
	session.Query(`CREATE TABLE IF NOT EXISTS dev.users3 (userid int PRIMARY KEY,
		name text,
		age int,
		addr text,
		last_update_timestamp timestamp
		);`).WithContext(ctx).Exec()

	fmt.Println("Create Table OK!")
}
