package services

import (
	"iris-xorm/dao"
	"iris-xorm/datasource"
	"iris-xorm/models"
)

type SuperstarService interface {
	GetAll() []models.StarInfo
	Get(id int) *models.StarInfo
	Delete(id int) error
	Update(star *models.StarInfo, columns []string) error
	Create(star *models.StarInfo) error

	Search(country string) []models.StarInfo
}

type superstarService struct {
	dao *dao.SuperstarDao
}

func NewSuperstarService() SuperstarService {
	return &superstarService{dao: dao.NewSuperstarDao(datasource.InstanceMaster())}
}

func (s *superstarService) GetAll() []models.StarInfo {
	return s.dao.GetAll()
}

func (s *superstarService) Get(id int) *models.StarInfo {
	return s.dao.Get(id)
}

func (s *superstarService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *superstarService) Update(star *models.StarInfo, columns []string) error {
	return s.dao.Update(star, columns)
}

func (s *superstarService) Create(star *models.StarInfo) error {
	return s.dao.Create(star)
}

func (s *superstarService) Search(country string) []models.StarInfo {
	return s.dao.Search(country)
}
