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
		price_arr := GetPriceList(v.Code.Id, v.LastUpdated)
		//log.Println(price_arr)
		sum_list, week, month, quarter := sum_by_unit(v.Code.Id, price_arr)
		err := dao.InsertTbSum(sum_list)
		if err != nil {
			log.Fatal("InsertCodeUnit err ===> ", err)
			log.Panic(err)
		}
		w_arr := detail_by_year(week, model.GetUnitValue("w"))
		m_arr := detail_by_year(month, model.GetUnitValue("m"))
		q_arr := detail_by_year(quarter, model.GetUnitValue("q"))

		// log.Println(w_arr)
		// log.Println(m_arr)
		// log.Println(q_arr)

		list := append(w_arr, m_arr...)
		list = append(list, q_arr...)
		cu := model.CodeUnit{
			Code: v.Code,
			List: list,
		}
		err = dao.InsertTbYear(cu)
		if err != nil {
			log.Fatal("InsertCodeUnit err ===> ", err)
			log.Panic(err)
		}
		//log.Println(cu)

	}

	return res

}

func GetPriceList(code_id int, dt int) []model.PriceInfo {
	res, err := dao.GetPriceList(code_id, dt)
	if err != nil {
		log.Panic("GetPriceList 에러")
	}
	return res
}

func sum_by_unit(code_id int, list []model.PriceInfo) ([]model.CodeSum, map[int]map[int]int, map[int]map[int]int, map[int]map[int]int) {
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

	return res, week, month, quarter

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

func detail_by_year(unit_map map[int]map[int]int, unit_type int) []model.UnitByYear {

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
