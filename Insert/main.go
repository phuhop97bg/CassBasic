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
	userid := 9
	name := "hoang anh"
	age := "22"
	addr := "bac ninh"

	query := session.Query(`INSERT INTO dev.users (userid, name, age, addr, last_update_timestamp)
	VALUES (?, ? , ? ,? , toTimeStamp(now()));`, userid, name, age, addr)

	if err = query.WithContext(ctx).Exec(); err != nil {
		log.Fatalf("-------------------------------------------------------%v", err)
	}

	fmt.Println("INSERT DATA OK!")
}
