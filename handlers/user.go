package handlers

import (
	"Backend-Challenge-Batara-Guru/api/operations/user"
	"Backend-Challenge-Batara-Guru/models"
	"Backend-Challenge-Batara-Guru/utils"
	"context"
	"errors"
	"time"

	"github.com/go-openapi/strfmt"
)

func (h *handler) Login(ctx context.Context, params user.PostLoginParams) (result *models.InfoToken, err error) {
	//Check username and password
	res := &models.User{}
	model := h.db.Model(res)
	if err = model.Where("username = ?", params.Body.Username).Find(&res).Error; err != nil {
		return nil, err
	}
	matchPass := utils.CheckPasswordHash(*params.Body.Password, res.Password)
	if res.Username != *params.Body.Username || !matchPass {
		return nil, errors.New("Usernam or Password Wrong")
	}

	// Create token
	token, err := utils.GenerateJWT(res.ID, res.Username, res.Role)
	if err != nil {
		return nil, err
	}

	DetailToken, err := utils.GetDetailToken(token)
	if err != nil {
		return nil, err
	}

	result = &models.InfoToken{
		Token:      token,
		Role:       DetailToken.Role,
		ExpireDate: DetailToken.ExpireDate,
	}

	return result, nil
}

func (h *handler) CreateUser(ctx context.Context, params user.PostUserParams) error {
	var err error

	// Hash Password
	passHash, err := utils.HashPassword(params.Body.Password)
	if err != nil {
		return err
	}

	// Create
	user := &models.User{
		Username:  params.Body.Username,
		Password:  string(passHash),
		Role:      params.Body.Role,
		UpdatedAt: strfmt.DateTime(time.Now()),
	}
	if err := h.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (h *handler) GetAllUser(ctx context.Context, params user.GetUserParams) ([]*models.ResUser, *models.Metadata, error) {
	res := []*models.ResUser{}
	user := []models.User{}
	offset := 0
	model := h.db.Model(user)
	if params.Page > 0 {
		offset = (int(params.Page) - 1) * int(params.Limit)
	}
	model = model.Limit(int(params.Limit)).Offset(offset)
	if err := model.Find(&user).Error; err != nil {
		return nil, nil, err
	}

	for _, v := range user {
		res = append(res, &models.ResUser{
			ID:       v.ID,
			Username: v.Username,
			Role:     v.Role,
		})
	}

	var totalRow int64
	model.Count(&totalRow)

	return res, utils.GenMetadata(totalRow, params.Limit, params.Page), nil
}

func (h *handler) UpdateUserById(ctx context.Context, params user.PutUserIDParams) error {
	var err error

	// Hash Password
	passHash, err := utils.HashPassword(params.Body.Password)
	if err != nil {
		return err
	}

	// Update
	user := &models.User{
		Username:  params.Body.Username,
		Password:  string(passHash),
		Role:      params.Body.Role,
		UpdatedAt: strfmt.DateTime(time.Now()),
	}

	result := h.db.Where("id = ?", params.ID).Updates(&user)
	if result.RowsAffected == 0 {
		return errors.New("Id not found")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (h *handler) DeleteUserById(ctx context.Context, params user.DeleteUserIDParams) error {
	result := h.db.Where("id = ?", params.ID).Delete(&models.User{})
	if result.RowsAffected == 0 {
		return errors.New("Id not found")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
