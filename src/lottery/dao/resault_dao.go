package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/lunny/log"
	"lottery/models"
)

type ResaultDao struct {
	engine *xorm.Engine
}

func NewResaultDao(engine *xorm.Engine) *ResaultDao {
	return &ResaultDao{engine: engine}
}

func (d *ResaultDao) Get(id int) *models.LtResult {
	data := &models.LtResult{Id: id}
	ok, err := d.engine.Get(data)
	if !ok || err != nil {
		log.Println("result_dao.Get error=", err)
		data.Id = 0
	}
	return data
}

func (d *ResaultDao) GetAll() []models.LtResult {
	dataList := make([]models.LtResult, 0)
	err := d.engine.Asc("sys_status", "display_order").Find(&dataList)
	if err != nil {
		log.Println("result_dao.GetAll error=", err)
	}
	return dataList
}

func (d *ResaultDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtResult{})
	if err != nil {
		log.Println("result_dao.CountAll error=", err)
		return 0
	}
	return num
}

func (d *ResaultDao) Delete(id int) error {
	data := &models.LtResult{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *ResaultDao) Update(data *models.LtResult, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *ResaultDao) Create(data *models.LtResult) error {
	_, err := d.engine.Id(data.Id).Insert(data)
	return err
}
