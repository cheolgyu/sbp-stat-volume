package handler

import (
	"github.com/cheolgyu/base/logging"
	cmm_model "github.com/cheolgyu/model"
	"github.com/cheolgyu/stat/price/dao"
	"github.com/cheolgyu/stat/price/model"
)

var TimeFrames []model.TimeFrame
var Configs []cmm_model.Config

var CONFIG_OP = -1
var CONFIG_CP = -1
var CONFIG_LP = -1
var CONFIG_HP = -1

func init() {
	setTimeFrames()
	configs, _ := dao.GetConfigListByUpperCode()
	Configs = configs

	for _, v := range configs {
		switch v.Code {
		case "open":
			CONFIG_OP = v.Id
		case "close":
			CONFIG_CP = v.Id
		case "low":
			CONFIG_LP = v.Id
		case "high":
			CONFIG_HP = v.Id
		}
	}

	dao.Initail_table()
}

func Handler() {

	codes, err := dao.GetCodeAll()

	if err != nil {
		logging.Log.Panic(err)
	}
	for _, v := range codes {

		code_info, err := dao.SelectList(v.Id)
		if err != nil {
			logging.Log.Panic(err)
		}
		list := find(code_info)
		err = dao.Insert(list)
		if err != nil {
			logging.Log.Panic(err)
		}
	}

}

func find(item model.CodeInfo) []cmm_model.Tb52Weeks {

	var res []cmm_model.Tb52Weeks
	res = append(res, findPointInfo(item.Code.Id, item.OP, CONFIG_OP)...)
	res = append(res, findPointInfo(item.Code.Id, item.CP, CONFIG_CP)...)
	res = append(res, findPointInfo(item.Code.Id, item.LP, CONFIG_LP)...)
	res = append(res, findPointInfo(item.Code.Id, item.HP, CONFIG_HP)...)

	return res
}

func findPointInfo(code_id int, arr []model.PointInfo, price_type int) []cmm_model.Tb52Weeks {
	var tmp model.PriceInfo

	var res []cmm_model.Tb52Weeks

	if len(arr) > 0 {
		tmp.Cur.X = arr[0].Point.X
		tmp.Cur.Y = arr[0].Point.Y
		tmp.Min.X = arr[0].Point.X
		tmp.Min.Y = arr[0].Point.Y
		tmp.Max.X = arr[0].Point.X
		tmp.Max.Y = arr[0].Point.Y
	}
	var break_timeframes int = 0

	for _, v := range arr {

		if v.Point.Y >= tmp.Max.Y {
			tmp.Max.Y = v.Point.Y
			tmp.Max.X = v.Point.X
		} else {
			tmp.Min.Y = v.Point.Y
			tmp.Min.X = v.Point.X
		}

		for _, t := range TimeFrames {
			if v.Xcnt > t.Day && break_timeframes < t.Day {
				break_timeframes = t.Day

				max_item := cmm_model.Tb52Weeks{
					Code_id:    code_id,
					Price_type: price_type,
					P3_type:    cmm_model.P3_type_HIGH,
					P1x_Left:   t.Day,
					P1: cmm_model.P{
						X: v.Point.X,
						Y: v.Point.Y,
					},
					P2: cmm_model.P{
						X: tmp.Cur.X,
						Y: tmp.Cur.Y,
					},
					P3: cmm_model.P{
						X: tmp.Max.X,
						Y: tmp.Max.Y,
					},
					P32y_percent: cmm_model.Get_percent(tmp.Max.Y, tmp.Cur.Y),
				}
				res = append(res, max_item)

				min_item := cmm_model.Tb52Weeks{
					Code_id:    code_id,
					Price_type: price_type,
					P3_type:    cmm_model.P3_type_LOW,
					P1x_Left:   t.Day,
					P1: cmm_model.P{
						X: v.Point.X,
						Y: v.Point.Y,
					},
					P2: cmm_model.P{
						X: tmp.Cur.X,
						Y: tmp.Cur.Y,
					},
					P3: cmm_model.P{
						X: tmp.Min.X,
						Y: tmp.Min.Y,
					},

					P32y_percent: cmm_model.Get_percent(tmp.Min.Y, tmp.Cur.Y),
				}
				res = append(res, min_item)
			}
		}

	}

	return res
}

//일자 목록 구하기
func setTimeFrames() {

	for i := 1; i <= 52; i++ {
		TimeFrames = append(TimeFrames, model.TimeFrame{Day: 7 * i, UnitType: 1, UnitVal: i})
	}

}
