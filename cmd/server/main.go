package main

import (
	db "RestApi2.0/internal/db"
	"context"
	"fmt"
)

// Run - responsible for installation and starting of app
func Run() error {
	fmt.Println("Starting app")

	db, err := db.NewDataBase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
	}

	if err = db.Ping(context.Background()); err != nil {
		fmt.Errorf("error in ping: %w", err)

	}

	if err = db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate a db")
		return err
	}

	fmt.Println("Successfully connects to database")
	return nil
}

func main() {
	fmt.Println("Rest Api ")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
