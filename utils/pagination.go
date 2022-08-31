package utils

import (
	"Backend-Challenge-Batara-Guru/models"
	"math"
)

func GenMetadata(totalRow, limit, page int64) *models.Metadata {
	totalPage := math.Ceil(float64(totalRow) / float64(limit))
	meta := &models.Metadata{
		Page:      int64(page),
		PerPage:   int64(limit),
		TotalPage: int64(totalPage),
		TotalRow:  totalRow,
	}
	return meta
}
