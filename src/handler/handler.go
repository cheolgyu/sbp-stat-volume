package handler

import (
	"log"
	"math"
	"sort"

	"github.com/cheolgyu/stock-write-project-trading-volume/src/dao"
	"github.com/cheolgyu/stock-write-project-trading-volume/src/model"
)

func GetCodeList() []model.CodeInfo {

	res, err := dao.GetCodeList()
	if err != nil {
		log.Panic("GetCodeList 에러")
	}

	//log.Println(res)
	for _, v := range res {
		DoSumByUnit(v)
		DoCalculateUnitDataByYear(v)
		DoYearOfTotal(v)
	}

	return res
}

/*
코드의 마지막 업데이트일 이후의 가격목록을 unit별 합계를 구해서 TB_SUM에 저장하기
*/
func DoSumByUnit(code_info model.CodeInfo) {
	new_price_arr, err := dao.GetPriceList(code_info.Code.Id, code_info.LastUpdated)
	if err != nil {
		log.Panic("GetPriceList 에러")
	}

	//log.Println(price_arr)
	sum_list := sum_by_unit(code_info.Code.Id, new_price_arr)
	err = dao.InsertTbSum(sum_list)
	if err != nil {
		log.Fatal("InsertCodeUnit err ===> ", err)
		log.Panic(err)
	}
}

/*
TB_SUM에서 코드의 마지막 업데이트일자의 연도보다 같거나 큰 TB_SUM 목록을 단위별로 조회 후 TB_YEAR에 저장하기
*/
func DoCalculateUnitDataByYear(code_info model.CodeInfo) {
	var res []model.UnitByYear

	for _, v := range model.UnitType {

		list, err := dao.SelectTbSumByUnitType(code_info, v)
		if err != nil {
			log.Fatal("InsertCodeUnit err ===> ", err)
			log.Panic(err)
		}
		unit_map := convert_codesum_to_map(list)
		arr := agg_by_year(unit_map, v)
		res = append(res, arr...)
	}

	//insert year
	cu := model.CodeUnit{
		Code: code_info.Code,
		List: res,
	}
	err := dao.InsertTbYear(cu)
	if err != nil {
		log.Fatal("InsertCodeUnit err ===> ", err)
		log.Panic(err)
	}

}

/*
TB_YEAR목록을 조회 후 계산하여 저장 하기.
*/
func DoYearOfTotal(code_info model.CodeInfo) {
	// 조회
	// 계산
	// 저장
}

func sum_by_unit(code_id int, list []model.PriceInfo) []model.CodeSum {
	//([]model.UnitByYear, []model.UnitByYear, []model.UnitByYear) {

	week := make(map[int]map[int]int)
	month := make(map[int]map[int]int)
	quarter := make(map[int]map[int]int)

	for _, v := range list {

		if _, exist := week[v.Opening.YY]; !exist {
			week[v.Opening.YY] = map[int]int{}
		}
		week[v.Opening.YY][v.Opening.Week] += v.Price.Volume

		if _, exist := month[v.Opening.YY]; !exist {
			month[v.Opening.YY] = map[int]int{}
		}
		month[v.Opening.YY][v.Opening.MM] += v.Price.Volume

		if _, exist := quarter[v.Opening.YY]; !exist {
			quarter[v.Opening.YY] = map[int]int{}
		}
		quarter[v.Opening.YY][v.Opening.Quarter] += v.Price.Volume

	}
	var res []model.CodeSum
	res = append(res, sum(code_id, model.GetUnitValue("w"), week)...)
	res = append(res, sum(code_id, model.GetUnitValue("m"), month)...)
	res = append(res, sum(code_id, model.GetUnitValue("q"), quarter)...)
	// log.Println(week)
	// log.Println(month)
	// log.Println(quarter)

	return res

}

func sum(code_id int, unit_type int, unit_map map[int]map[int]int) []model.CodeSum {
	var res []model.CodeSum
	for year, year_map := range unit_map {
		for k, v := range year_map {
			item := model.CodeSum{}
			item.Code.Id = code_id
			item.UnitType = unit_type
			item.Year = year
			item.Unit = k
			item.Sum = v
			item.SetRowPk()
			res = append(res, item)
		}
	}
	return res
}

/*
TB_CODESUM 목록을 연도의 UNIT단위의 합값 형태로 변환한다.
*/
func convert_codesum_to_map(list []model.CodeSum) map[int]map[int]int {
	unit_datas := make(map[int]map[int]int)

	for _, v := range list {
		if _, exist := unit_datas[v.Year]; !exist {
			unit_datas[v.Year] = map[int]int{}
		}
		unit_datas[v.Year][v.Unit] = v.Sum
	}

	return unit_datas
}

/*
TB_CODESUM의 연도별 unit별 맵데이터로 연도별 unit의 집계 데이터를 구한다.
*/
func agg_by_year(unit_map map[int]map[int]int, unit_type int) []model.UnitByYear {

	res := []model.UnitByYear{}

	sort_year := make([]int, 0, len(unit_map))
	for year := range unit_map {
		sort_year = append(sort_year, year)
	}
	sort.Ints(sort_year)

	for _, year := range sort_year {
		item := model.UnitByYear{}
		//log.Println(year, "==================len=", len(unit_map[year]))

		sort_keys := make([]int, 0, len(unit_map[year]))
		for unit := range unit_map[year] {
			sort_keys = append(sort_keys, unit)
		}
		sort.Ints(sort_keys)

		// for _, v := range sort_unit {
		// 	log.Println(v, week[year][v])
		// }
		// max
		max_k := -1
		max_v := -1

		for k, v := range unit_map[year] {
			if max_v < v {
				max_k = k
				max_v = v
			}
		}
		min_k := -1
		min_v := max_v

		for k, v := range unit_map[year] {
			if min_v > v {
				min_k = k
				min_v = v
			}
		}

		//log.Println("max k,v=", max_k, max_v)
		//log.Println("min k,v=", min_k, min_v)

		// avg
		sum := -1
		for _, v := range unit_map[year] {
			sum += v
		}
		avg_v := sum / len(unit_map[year])

		//log.Println("avg_v =", avg_v)

		// split avg
		var up_arr []int
		var down_arr []int
		for k, v := range unit_map[year] {
			if avg_v < v {
				up_arr = append(up_arr, k)
			} else {
				down_arr = append(down_arr, k)
			}
		}
		sort.Ints(up_arr)
		sort.Ints(down_arr)
		//log.Println("split avg  up_arr,down_arr=", up_arr, down_arr)

		//percent
		percent := make(map[int]float64)
		for _, k := range sort_keys {
			//log.Println("k,v=", k, float32(week[year][k])/float32(sum)*100)
			per := float64(unit_map[year][k]) / float64(sum) * 100
			if !math.IsNaN(per) {
				percent[k] = math.Round(per*100) / 100
			}

		}
		//log.Println("percent=", percent)

		item.Unit = unit_type
		item.Year = year
		item.Max = max_k
		item.Min = min_k
		item.Up = up_arr
		item.Down = down_arr
		item.Percent = percent
		item.Avg = avg_v

		res = append(res, item)
	}

	return res

}
