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
	session.Query(`CREATE INDEX IF NOT EXISTS age_user ON dev.users (  age  )`).WithContext(ctx).Exec()

	fmt.Println("Create Index OK!")
}
