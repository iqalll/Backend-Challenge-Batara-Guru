package fakers

import (
	"Backend-Challenge-Batara-Guru/models"
	"strings"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/go-openapi/strfmt"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) []*models.User {
	res := []*models.User{}
	for i := 0; i < 5; i++ {
		res = append(res, &models.User{
			Username:  strings.ToLower(faker.FirstName()),
			Password:  "$2a$14$LeSrpFVNXAQciRBIE4XqxerrWJHqyF1NS8PMdFWWOz0nJBxpE5dHW", // 123
			Role:      "admin",
			UpdatedAt: strfmt.DateTime(time.Now()),
		})
	}
	return res
}
