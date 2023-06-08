package main

import (
	"fmt"

	"eegurt.go-clean-arch/pkg/store/postgres"
)

func main() {
	dcp := &postgres.DbConnParams{
		Host:     "localhost",
		Port:     5432,
		User:     "architect",
		Password: "clean123",
		DbName:   "go_clean_arch",
	}

	db, err := postgres.OpenDB(dcp)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	fmt.Println("app started")
}
