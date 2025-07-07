package repository

import "github.com/bhushan-aruto/justarquest/internal/entity"

type CreateUserRepo interface {
	CreateUser(justrequeest *entity.Jsutrequest) error
}
