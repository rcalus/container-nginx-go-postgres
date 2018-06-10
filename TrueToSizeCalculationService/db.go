package main

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func connect() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     "db:5432",
		User:     "postgres-dev",
		Password: "s3cr3tp4ssw0rd",
		Database: "dev",
	})

	return db
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*TrueToSize)(nil)} {
		fmt.Println(model)
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
