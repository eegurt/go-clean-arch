package main

import (
	"fmt"
	"net/http"

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

	defer db.Pool.Close()

	repo := repository.New(db.Pool)
	delivery := delivery.New()
	usecase := usecase.New(repo)

	_ = usecase

	fmt.Println("app started")
	http.ListenAndServe("localhost:4000", delivery.Mux)
}
