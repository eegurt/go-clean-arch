package repository

import (
	"eegurt.go-clean-arch/services/contact/internal/domain"
)

type ContactRepo interface {
	Create(user *domain.Contact) (int, error)
	Read(id int) (*domain.Contact, error)
	Update(id int) error
	Delete(id int) error
}

type GroupRepo interface {
	Create(group *domain.Group) (int, error)
	Read(id int) (*domain.Group, error)
	AddContact(id int) error
}
