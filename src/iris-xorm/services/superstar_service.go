package services

import (
	"iris-xorm/models"
)

type SuperstarService interface {
	GetAll() []models.StarInfo
	Get(id int) models.StarInfo
	Delete(id int) bool
	Update(star *models.StarInfo) error
	Create(star *models.StarInfo) error

	Search(country string) []models.StarInfo
}
