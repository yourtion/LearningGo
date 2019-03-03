package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/lunny/log"
	"lottery/models"
)

type UserDay struct {
	engine *xorm.Engine
}

func NewUserDay(engine *xorm.Engine) *UserDay {
	return &UserDay{engine: engine}
}

func (d *UserDay) Get(id int) *models.LtUserDay {
	data := &models.LtUserDay{Id: id}
	ok, err := d.engine.Get(data)
	if !ok || err != nil {
		log.Println("user_day_dao.Get error=", err)
		data.Id = 0
	}
	return data
}

func (d *UserDay) GetAll() []models.LtUserDay {
	dataList := make([]models.LtUserDay, 0)
	err := d.engine.Asc("sys_status", "display_order").Find(&dataList)
	if err != nil {
		log.Println("user_day_dao.GetAll error=", err)
	}
	return dataList
}

func (d *UserDay) CountAll() int64 {
	num, err := d.engine.Count(&models.LtUserDay{})
	if err != nil {
		log.Println("user_day_dao.CountAll error=", err)
		return 0
	}
	return num
}

func (d *UserDay) Delete(id int) error {
	data := &models.LtUserDay{Id: id}
	_, err := d.engine.Id(data.Id).Delete(data)
	return err
}

func (d *UserDay) Update(data *models.LtUserDay, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *UserDay) Create(data *models.LtUserDay) error {
	_, err := d.engine.Id(data.Id).Insert(data)
	return err
}
