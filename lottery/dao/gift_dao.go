package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/lunny/log"
	"lottery/models"
)

type GiftDao struct {
	engine *xorm.Engine
}

func NewGiftDao(engine *xorm.Engine) *GiftDao {
	return &GiftDao{engine: engine}
}

func (d *GiftDao) Get(id int) *models.LtGift {
	data := &models.LtGift{Id: id}
	ok, err := d.engine.Get(data)
	if !ok || err != nil {
		log.Println("gift_dao.Get error=", err)
		data.Id = 0
	}
	return data
}

func (d *GiftDao) GetAll() []models.LtGift {
	dataList := make([]models.LtGift, 0)
	err := d.engine.Asc("sys_status", "display_order").Find(&dataList)
	if err != nil {
		log.Println("gift_dao.GetAll error=", err)
	}
	return dataList
}

func (d *GiftDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtGift{})
	if err != nil {
		log.Println("gift_dao.CountAll error=", err)
		return 0
	}
	return num
}

func (d *GiftDao) Delete(id int) error {
	data := &models.LtGift{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *GiftDao) Update(data *models.LtGift, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *GiftDao) Create(data *models.LtGift) error {
	_, err := d.engine.Id(data.Id).Insert(data)
	return err
}
