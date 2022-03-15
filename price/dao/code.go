package dao

import (
	"github.com/cheolgyu/base/db"
	"github.com/cheolgyu/model"
	"github.com/cheolgyu/tb/code"
	"github.com/cheolgyu/tb/config"
	"github.com/cheolgyu/tb/info"
)

func Update_info() {
	info.UpdateNow(db.Conn, info.NAME_UPDATE_STAT_PRICE)

}

func GetCodeAll() ([]model.Code, error) {
	res, err := code.GetCodeList(db.Conn)
	return res, err
}

func GetConfigListByUpperCode() ([]model.Config, error) {
	res, err := config.GetConfigListByUpperCode(db.Conn, config.CONFIG_UPPER_CODE_PRICE_TYPE)
	return res, err
}
