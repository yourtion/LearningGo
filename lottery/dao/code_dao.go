package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/lunny/log"
	"lottery/models"
)

type CodeDao struct {
	engine *xorm.Engine
}

func NewCodeDao(engine *xorm.Engine) *CodeDao {
	return &CodeDao{engine: engine}
}

func (d *CodeDao) Get(id int) *models.LtCode {
	data := &models.LtCode{Id: id}
	ok, err := d.engine.Get(data)
	if !ok || err != nil {
		log.Println("code_dao.Get error=", err)
		data.Id = 0
	}
	return data
}

func (d *CodeDao) GetAll() []models.LtCode {
	dataList := make([]models.LtCode, 0)
	err := d.engine.Desc("id").Find(&dataList)
	if err != nil {
		log.Println("code_dao.GetAll error=", err)
	}
	return dataList
}

func (d *CodeDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtCode{})
	if err != nil {
		log.Println("code_dao.CountAll error=", err)
		return 0
	}
	return num
}

func (d *CodeDao) Delete(id int) error {
	data := &models.LtCode{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *CodeDao) Update(data *models.LtCode, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *CodeDao) Create(data *models.LtCode) error {
	_, err := d.engine.Id(data.Id).Insert(data)
	return err
}
