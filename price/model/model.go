package model

import (
	cmm_model "github.com/cheolgyu/model"
)

/*
DayCnt : 해당 일자가 365일에서 얼마나 경과된 일자인지
*/
type Res struct {
	cmm_model.Code
	cmm_model.PriceMarket
	DayCnt int
}

type CodeInfo struct {
	cmm_model.Code

	OP []PointInfo
	HP []PointInfo
	LP []PointInfo
	CP []PointInfo
}

func (o *Res) ByPrice() [4]PointInfo {
	var list [4]PointInfo

	dt := o.PriceMarket.Dt

	list[0] = PointInfo{
		Point: Point{
			dt, o.PriceMarket.OpenPrice,
		},
		Xcnt: o.DayCnt,
	}
	list[1] = PointInfo{
		Point: Point{
			dt, o.PriceMarket.ClosePrice,
		},
		Xcnt: o.DayCnt,
	}
	list[2] = PointInfo{
		Point: Point{
			dt, o.PriceMarket.LowPrice,
		},
		Xcnt: o.DayCnt,
	}
	list[3] = PointInfo{
		Point: Point{
			dt, o.PriceMarket.HighPrice,
		},
		Xcnt: o.DayCnt,
	}

	return list
}

/*
X : DATE
Y : PRICE
*/
type Point struct {
	X int
	Y float32
}

/*

Xcnt : int

Xcnt :: X 값이 365일 부터 몇일째인지 sql로 계산된 일수
*/
type PointInfo struct {
	Point Point
	Xcnt  int
}

// type PriceArr struct {
// 	PointInfos []PointInfo
// }

// // 0:op 1:cp 2:lp 3:hp
// type PriceCode struct {
// 	PriceArr []PriceArr
// }
/*
 0:op 1:cp 2:lp 3:hp
*/
type PriceInfo struct {
	Cur Point
	Min Point
	Max Point
}

type PriceInfoItemRes struct {
	cmm_model.Code
	PriceType int
	Arr       []PriceInfoItem
}

type PriceInfoItem struct {
	PriceInfo
	TimeFrame
}

func (o *PriceInfoItem) to_comm_model() [2]cmm_model.Tb52Weeks {
	var res [2]cmm_model.Tb52Weeks

	return res
}

type TimeFrame struct {
	Day      int
	UnitType int
	UnitVal  int
}
