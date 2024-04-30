package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

const (
	dataSourceName="postgresql://postgres:root@localhost:5432/simple_bank?sslmode=disable"
)
var testQueries *Queries

func TestMain(m *testing.M) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dataSourceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer conn.Close(ctx)

	testQueries = New(conn)
	
	os.Exit(m.Run())
}