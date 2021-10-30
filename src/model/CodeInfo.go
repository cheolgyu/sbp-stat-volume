package model

import (
	"fmt"
)

var UnitType map[string]int

func init() {
	UnitType = map[string]int{
		"w": 1,
		"m": 2,
		"q": 3,
	}
}

type CodeInfo struct {
	Code
	LastUpdated int
	Opening
}

type Code struct {
	Id        int
	Code      string
	Code_type int
}

type PriceInfo struct {
	Price
	Opening
}

type Price struct {
	Dt     int
	Volume int
}

type Opening struct {
	YY      int
	MM      int
	DD      int
	Week    int
	Quarter int
}
type CodeSum struct {
	Row_pk string
	Code
	UnitType int
	Year     int
	Unit     int
	Sum      int
}

func (o *CodeSum) SetRowPk() {
	o.Row_pk = fmt.Sprintf("%v_%v_%v_%v", o.Code.Id, o.UnitType, o.Year, o.Unit)
}

func GetUnitValue(unit string) int {
	v := -1
	if unit == "w" {
		v = 1
	} else if unit == "m" {
		v = 2
	} else if unit == "q" {
		v = 3
	}
	return v
}

type CodeUnit struct {
	Code

	List []UnitByYear
}

type UnitByYear struct {
	// 1 :week ,2 : month, 3: q
	Unit int
	Year int
	Max  int
	Min  int
	Avg  int
	Up   []int
	Down []int
	Rate map[int]float64
}

type CodeYear struct {
	Code
	UnitType int
	UnitByYear
}
type CodeTotal struct {
	Code
	UnitType int
	UnitByTotal
}

type UnitByTotal struct {
	// 1 :week ,2 : month, 3: q
	Unit    int
	YearCnt int

	MaxUnit    int
	MaxPercent float64
	MinUnit    int
	MinPercent float64
	MaxRate    map[int]float64
	MinRate    map[int]float64
	MaxArrRate map[int]float64
	MinArrRate map[int]float64
	Avg        int
}
