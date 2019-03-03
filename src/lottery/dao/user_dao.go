package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/lunny/log"
	"lottery/models"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{engine: engine}
}

func (d *UserDao) Get(id int) *models.LtUser {
	data := &models.LtUser{Id: id}
	ok, err := d.engine.Get(data)
	if !ok || err != nil {
		log.Println("user_dao.Get error=", err)
		data.Id = 0
	}
	return data
}

func (d *UserDao) GetAll() []models.LtUser {
	dataList := make([]models.LtUser, 0)
	err := d.engine.Asc("sys_status", "display_order").Find(&dataList)
	if err != nil {
		log.Println("user_dao.GetAll error=", err)
	}
	return dataList
}

func (d *UserDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtUser{})
	if err != nil {
		log.Println("user_dao.CountAll error=", err)
		return 0
	}
	return num
}

func (d *UserDao) Delete(id int) error {
	data := &models.LtUser{Id: id}
	_, err := d.engine.Id(data.Id).Delete(data)
	return err
}

func (d *UserDao) Update(data *models.LtUser, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *UserDao) Create(data *models.LtUser) error {
	_, err := d.engine.Id(data.Id).Insert(data)
	return err
}
