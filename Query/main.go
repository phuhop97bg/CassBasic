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

	scanner := session.Query(`SELECT name, age, addr FROM dev.users WHERE age > ? ALLOW FILTERING`, 25).WithContext(ctx).Iter().Scanner()
	for scanner.Next() {
		var (
			name string
			age  int
			addr string
		)
		err = scanner.Scan(&name, &age, &addr)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Users:", name, age, addr)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
