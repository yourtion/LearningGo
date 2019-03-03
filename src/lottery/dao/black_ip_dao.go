package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/lunny/log"
	"lottery/models"
)

type BlackIpDao struct {
	engine *xorm.Engine
}

func NewBlackIpDao(engine *xorm.Engine) *BlackIpDao {
	return &BlackIpDao{engine: engine}
}

func (d *BlackIpDao) Get(id int) *models.LtBlackIp {
	data := &models.LtBlackIp{Id: id}
	ok, err := d.engine.Get(data)
	if !ok || err != nil {
		log.Println("black_ip_dao.Get error=", err)
		data.Id = 0
	}
	return data
}

func (d *BlackIpDao) GetAll() []models.LtBlackIp {
	dataList := make([]models.LtBlackIp, 0)
	err := d.engine.Desc("id").Find(&dataList)
	if err != nil {
		log.Println("black_ip_dao.GetAll error=", err)
	}
	return dataList
}

func (d *BlackIpDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtBlackIp{})
	if err != nil {
		log.Println("black_ip_dao.CountAll error=", err)
		return 0
	}
	return num
}

func (d *BlackIpDao) Delete(id int) error {
	data := &models.LtBlackIp{Id: id}
	_, err := d.engine.Delete(data)
	return err
}

func (d *BlackIpDao) Update(data *models.LtBlackIp, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *BlackIpDao) Create(data *models.LtBlackIp) error {
	_, err := d.engine.Id(data.Id).Insert(data)
	return err
}

// 根据IP获取信息
func (d *BlackIpDao) GetByIp(ip string) *models.LtBlackIp {
	dataList := make([]models.LtBlackIp, 0)
	err := d.engine.
		Where("ip=?", ip).
		Desc("id").
		Limit(1).
		Find(&dataList)
	if err != nil || len(dataList) < 1 {
		return nil
	} else {
		return &dataList[0]
	}
}
