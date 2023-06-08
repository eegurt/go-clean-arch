package main

import (
	"fmt"

	"eegurt.go-clean-arch/pkg/store/postgres"
	"eegurt.go-clean-arch/services/contact/internal/delivery"
	"eegurt.go-clean-arch/services/contact/internal/repository"
	usecase "eegurt.go-clean-arch/services/contact/internal/useCase"
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

	repo := repository.New()
	delivery := delivery.New()
	usecase := usecase.New()

	_, _, _ = repo, delivery, usecase

	fmt.Println("app started")
	for {

	}
}
