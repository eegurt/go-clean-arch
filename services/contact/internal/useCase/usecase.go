package usecase

import (
	"eegurt.go-clean-arch/services/contact/internal/domain"
)

type UseCase interface {
	CreateContact(user *domain.Contact) (int, error)
	ReadContact(id int) (*domain.Contact, error)
	UpdateContact(id int) error
	DeleteContact(id int) error
	CreateGroup(group *domain.Group) (int, error)
	ReadGroup(id int) (*domain.Group, error)
	AddContact(id int) error
}

type useCase struct{}

func NewUseCase() *useCase {
	return &useCase{}
}
