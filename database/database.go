package database

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/pharmacity-xyz/server/ent"
)

var EntClient *ent.Client

func init() {
	Client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=pharmacity-db password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database successfully")
	defer Client.Close()

	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	EntClient = Client
}
