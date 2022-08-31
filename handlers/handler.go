package handlers

import (
	"Backend-Challenge-Batara-Guru/api/operations/gift"
	"Backend-Challenge-Batara-Guru/api/operations/user"
	"Backend-Challenge-Batara-Guru/db"
	"Backend-Challenge-Batara-Guru/models"
	"context"

	"gorm.io/gorm"
)

type handler struct {
	db *gorm.DB
}

type Handlers interface {
	// User
	Login(ctx context.Context, params user.PostLoginParams) (res *models.InfoToken, err error)
	CreateUser(ctx context.Context, params user.PostUserParams) error
	GetAllUser(ctx context.Context, params user.GetUserParams) ([]*models.ResUser, *models.Metadata, error)
	UpdateUserById(ctx context.Context, params user.PutUserIDParams) error
	DeleteUserById(ctx context.Context, params user.DeleteUserIDParams) error
	// Gift
	CreateGift(ctx context.Context, params gift.PostGiftParams) error
	GetAllGift(ctx context.Context, params gift.GetGiftParams) ([]*models.ResGift, *models.Metadata, error)
	GetGiftById(ctx context.Context, params gift.GetGiftIDParams) (*models.ResGift, error)
	UpdateGiftById(ctx context.Context, params gift.PutGiftIDParams) error
	DeleteGiftById(ctx context.Context, params gift.DeleteGiftIDParams) error
}

func NewHandler() Handlers {
	return &handler{db: db.NewGorm()}
}
