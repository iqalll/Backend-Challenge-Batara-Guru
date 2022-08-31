package handlers

import (
	"Backend-Challenge-Batara-Guru/api/operations/gift"
	"Backend-Challenge-Batara-Guru/models"
	"Backend-Challenge-Batara-Guru/utils"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/go-openapi/strfmt"
)

func (h *handler) CreateGift(ctx context.Context, params gift.PostGiftParams) error {
	file, valid, err := params.HTTPRequest.FormFile("photo_product")
	if err != nil {
		return err
	}
	// validtion extension
	ext := filepath.Ext(valid.Filename)
	if !utils.EXTENSION_FILE[ext] {
		return errors.New("The uploaded image must be .jpg/.png/.jpeg")
	}
	// validation size
	if valid.Size > int64(utils.MAX_SIZE_FILE) {
		return errors.New("The uploaded image is too big. Please use an image less than 2MB in size")
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("assets/img", "upload-*.png")
	if err != nil {
		return err
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err

	}
	tempFile.Write(fileBytes)

	gift := &models.Gift{
		PhotoProduct: "/" + path.Base(tempFile.Name()),
		NamaProduct:  params.NamaProduct,
		Point:        params.Point,
		Stock:        params.Stock,
		Description:  params.Description,
		UpdatedAt:    strfmt.DateTime(time.Now()),
	}
	if err := h.db.Create(&gift).Error; err != nil {
		return err
	}
	return nil
}

func (h *handler) UpdateGiftById(ctx context.Context, params gift.PutGiftIDParams) error {
	gift := &models.Gift{}
	// Get Data Gift By Id
	oldData := &models.Gift{}
	model := h.db.Model(oldData)
	dataGift := model.Where("id = ?", params.ID).Where("id = ?", params.ID).Find(&oldData)
	if dataGift.RowsAffected == 0 {
		return errors.New("Gift not found")
	}
	if dataGift.Error != nil {
		return dataGift.Error
	}

	if params.PhotoProduct == nil {
		gift = &models.Gift{
			NamaProduct: params.NamaProduct,
			Point:       params.Point,
			Stock:       params.Stock,
			Description: params.Description,
			UpdatedAt:   strfmt.DateTime(time.Now()),
		}
	} else {
		file, valid, err := params.HTTPRequest.FormFile("photo_product")
		if err != nil {
			return err
		}
		defer file.Close()
		// validtion extension
		ext := filepath.Ext(valid.Filename)
		if !utils.EXTENSION_FILE[ext] {
			return errors.New("The uploaded image must be .jpg/.png/.jpeg")
		}
		// validation size
		if valid.Size > int64(utils.MAX_SIZE_FILE) {
			return errors.New("The uploaded image is too big. Please use an image less than 2MB in size")
		}

		tempFile, err := ioutil.TempFile("assets/img", "upload-*.png")
		if err != nil {
			return err
		}
		// defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			return err

		}
		tempFile.Write(fileBytes)

		gift = &models.Gift{
			PhotoProduct: "/" + path.Base(tempFile.Name()),
			NamaProduct:  params.NamaProduct,
			Point:        params.Point,
			Stock:        params.Stock,
			Description:  params.Description,
			UpdatedAt:    strfmt.DateTime(time.Now()),
		}
	}
	result := h.db.Where("id = ?", params.ID).Updates(&gift)
	if result.Error != nil {
		return result.Error
	}
	if params.PhotoProduct != nil {
		// Delete old image in folder
		err := os.Remove("assets/img/" + oldData.PhotoProduct)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *handler) DeleteGiftById(ctx context.Context, params gift.DeleteGiftIDParams) error {
	oldData := &models.Gift{}
	model := h.db.Model(oldData)
	dataGift := model.Where("id = ?", params.ID).Where("id = ?", params.ID).Find(&oldData)
	if dataGift.RowsAffected == 0 {
		return errors.New("Gift not found")
	}
	if dataGift.Error != nil {
		return dataGift.Error
	}

	result := h.db.Where("id = ?", params.ID).Delete(&models.Gift{})
	if result.Error != nil {
		return result.Error
	}
	// Delete old image in folder
	err := os.Remove("assets/img/" + oldData.PhotoProduct)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) GetAllGift(ctx context.Context, params gift.GetGiftParams) ([]*models.ResGift, *models.Metadata, error) {
	res := []*models.ResGift{}
	gift := []models.Gift{}
	offset := 0
	model := h.db.Model(gift)
	if params.Page > 0 {
		offset = (int(params.Page) - 1) * int(params.Limit)
		model = model.Limit(int(params.Limit)).Offset(offset)
	}
	if params.Sorting != "" {
		sort := fmt.Sprintf(`%s`, params.Sorting)
		model = model.Order("id " + sort)
	}
	if err := model.Find(&gift).Error; err != nil {
		return nil, nil, err
	}

	for _, v := range gift {
		res = append(res, &models.ResGift{
			ID:           v.ID,
			PhotoProduct: v.PhotoProduct,
			NamaProduct:  v.NamaProduct,
			Point:        v.Point,
			Stock:        v.Stock,
			Description:  v.Description,
			Rating:       v.Rating,
		})
	}

	var totalRow int64
	model.Count(&totalRow)

	return res, utils.GenMetadata(totalRow, params.Limit, params.Page), nil
}

func (h *handler) GetGiftById(ctx context.Context, params gift.GetGiftIDParams) (*models.ResGift, error) {
	model := models.Gift{}
	if err := h.db.WithContext(ctx).Where("id = ?", params.ID).First(&model).Error; err != nil {
		return nil, err
	}
	res := &models.ResGift{
		ID:           model.ID,
		PhotoProduct: model.PhotoProduct,
		NamaProduct:  model.NamaProduct,
		Point:        model.Point,
		Stock:        model.Stock,
		Description:  model.Description,
		Rating:       model.Rating,
	}
	return res, nil
}
