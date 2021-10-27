package model

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
	Code
	UnitType int
	Year     int
	Unit     int
	Sum      int
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
