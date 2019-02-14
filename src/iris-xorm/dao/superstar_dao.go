package dao

import (
	"github.com/go-xorm/xorm"
	"iris-xorm/models"
	"log"
)

type SuperstarDao struct {
	engine *xorm.Engine
}

func NewSuperstarDao(engine *xorm.Engine) *SuperstarDao {
	return &SuperstarDao{engine: engine}
}

func (d *SuperstarDao) Get(id int) *models.StarInfo {
	data := &models.StarInfo{Id: id}
	ok, err := d.engine.Get(data)
	if !ok && err != nil {
		log.Println(err)
		data.Id = 0
	}
	return data
}

func (d *SuperstarDao) GetAll() []models.StarInfo {
	var dataList []models.StarInfo
	err := d.engine.Desc("id").Find(&dataList)
	if err != nil {
		log.Println(err)
	}
	return dataList
}

func (d *SuperstarDao) Search(country string) []models.StarInfo {
	var dataList []models.StarInfo
	err := d.engine.Where("country=?", country).Desc("id").Find(&dataList)
	if err != nil {
		log.Println(err)
	}
	return dataList
}

func (d *SuperstarDao) Delete(id int) error {
	data := &models.StarInfo{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Delete(data)
	return err
}

func (d *SuperstarDao) Update(data *models.StarInfo, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *SuperstarDao) Create(data *models.StarInfo) error {
	_, err := d.engine.Insert(data)
	return err
}
