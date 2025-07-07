package usecase

import (
	"errors"
	"log"

	"github.com/bhushan-aruto/justarquest/internal/entity"
	"github.com/bhushan-aruto/justarquest/internal/repository"
	"github.com/bhushan-aruto/justarquest/internal/utils"
)

type CreatesuersUsecase struct {
	databaserepo repository.CreateUserRepo
}

func NewUserCreate(db repository.CreateUserRepo) *CreatesuersUsecase {
	return &CreatesuersUsecase{
		databaserepo: db,
	}
}

func (u *CreatesuersUsecase) CreateUserUseacase(username, email, password string) error {
	if email == "" || password == "" || username == "" {
		log.Println("All fields are required")
		return errors.New("missing fields")
	}

	if !utils.IsValidEmail(email) {
		log.Println("Email is not valid")
		return errors.New("invalid email")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Println("Error while hashing password")
		return err
	}

	createUser := entity.NewJustRequest(username, email, hashedPassword)

	// Call the repository
	return u.databaserepo.CreateUser(createUser)
}
