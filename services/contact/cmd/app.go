package main

import (
	"fmt"

	"eegurt.go-clean-arch/pkg/store/postgres"
	"eegurt.go-clean-arch/services/contact/internal/domain"
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

	nikita := domain.NewContact("Romanov", "Nikita", "Savvich")
	aida := domain.NewContact("Osipova", "Aida", "Artemovna")
	group1 := domain.NewGroup("Group 1")

	fmt.Println(nikita, aida, group1)
}
