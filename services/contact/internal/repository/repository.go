package repository

import (
	"eegurt.go-clean-arch/services/contact/internal/domain"
)

type Repo interface {
	CreateContact(user *domain.Contact) (int, error)
	ReadContact(id int) (*domain.Contact, error)
	UpdateContact(id int) error
	DeleteContact(id int) error
	CreateGroup(group *domain.Group) (int, error)
	ReadGroup(id int) (*domain.Group, error)
	AddContact(id int) error
}

type repo struct{}

func NewRepo() *repo {
	return &repo{}
}
