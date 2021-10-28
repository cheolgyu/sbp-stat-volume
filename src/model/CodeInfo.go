package model

import (
	"fmt"
)

type CodeInfo struct {
	Code
	LastUpdated int
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
	Unit    int
	Year    int
	Max     int
	Min     int
	Avg     int
	Up      []int
	Down    []int
	Percent map[int]float64
}
