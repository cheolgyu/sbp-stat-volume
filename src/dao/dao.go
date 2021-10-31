package dao

import (
	"context"
	"encoding/json"
	"log"

	"github.com/cheolgyu/stock-write-project-trading-volume/src/db"
	"github.com/cheolgyu/stock-write-project-trading-volume/src/model"
	"github.com/lib/pq"
)

func GetCodeList() ([]model.CodeInfo, error) {
	var res []model.CodeInfo
	query := `
SELECT MC.ID,
	MC.CODE,
	MC.CODE_TYPE,
	COALESCE(TR.LAST_UPDATED,0) AS LAST_UPDATED,
	COALESCE(MO.YYYY,0 ) as yyyy,
	COALESCE(MO.MM,0 ) as mm,
	COALESCE(MO.WEEK,0 ) as week,
	COALESCE(MO.QUARTER,0 ) as quarter
FROM META.CODE MC
LEFT JOIN PROJECT_TRADING_VOLUME.TB_TOTAL TR ON MC.ID = TR.CODE_ID
LEFT JOIN META.OPENING MO ON TR.LAST_UPDATED = MO.DT
WHERE MC.CODE_TYPE = 1
	
`
	//log.Println(query)
	rows, err := db.Conn.QueryContext(context.Background(), query)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := model.CodeInfo{}

		if err := rows.Scan(&item.Code.Id, &item.Code.Code, &item.Code.Code_type, &item.LastUpdated,
			&item.Opening.YY, &item.Opening.MM, &item.Opening.Week, &item.Opening.Quarter); err != nil {
			log.Fatal(err)
			panic(err)
		}
		res = append(res, item)

	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	return res, err
}

func GetPriceList(code_id int, dt int) ([]model.PriceInfo, error) {
	var res []model.PriceInfo
	query := `
	SELECT
		p.dt,
		p.vol,
		o.yyyy,
		o.mm,
		o.dd,
		o.week,
		o.quarter
   FROM HIST.PRICE p left join meta.opening o on p.dt = o.dt
   where p.code_id = $1
   and p.dt > $2
   order by p.dt asc
   
	
`
	//log.Println(query)
	rows, err := db.Conn.QueryContext(context.Background(), query, code_id, dt)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := model.PriceInfo{}

		if err := rows.Scan(&item.Price.Dt, &item.Price.Volume, &item.Opening.YY, &item.Opening.MM, &item.Opening.DD, &item.Opening.Week, &item.Opening.Quarter); err != nil {
			log.Fatal(err)
			panic(err)
		}
		res = append(res, item)

	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	return res, err
}

const query_insert_tb_sum = `INSERT INTO  project_trading_volume.tb_sum (` +
	`row_pk, code_id, unit_type, yyyy, unit, sum_val)` +
	`VALUES ($1, $2, $3, $4, $5, $6)` +
	` ON CONFLICT (row_pk) DO UPDATE SET ` +
	`   sum_val = tb_sum.sum_val+$6  ; `

func InsertTbSum(list []model.CodeSum) error {

	client := db.Conn
	stmt, err := client.Prepare(query_insert_tb_sum)
	if err != nil {
		log.Println("쿼리:Prepare 오류: ")
		log.Fatal(err)
		panic(err)
	}
	defer stmt.Close()
	//code_id, unit_type,  max_unit, sum_val
	for _, v := range list {

		_, err = stmt.Exec(
			v.Row_pk, v.Code.Id, v.UnitType, v.Year, v.Unit, v.Sum,
		)
		if err != nil {
			log.Println("쿼리:stmt.Exec 오류: ")
			log.Printf("%+v ", v)
			log.Fatal(err)
			panic(err)
		}
	}

	return err
}

func SelectTbSumByUnitType(item model.CodeInfo, unit_type int) ([]model.CodeSum, error) {
	var res []model.CodeSum
	query := `
	SELECT row_pk, code_id, unit_type, yyyy, unit, sum_val
	FROM project_trading_volume.tb_sum
	WHERE	code_id =  $1 and unit_type = $2 
		and yyyy >= $3
		 
`
	//log.Println(query)
	rows, err := db.Conn.QueryContext(context.Background(), query, item.Code.Id, unit_type, item.Opening.YY)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := model.CodeSum{}

		if err := rows.Scan(&item.Row_pk, &item.Code.Id, &item.UnitType, &item.Year, &item.Unit, &item.Sum); err != nil {
			log.Fatal(err)
			panic(err)
		}
		res = append(res, item)

	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	return res, err
}

const query_insert_tb_year = `INSERT INTO  project_trading_volume.tb_year (` +
	`code_id, unit_type, yyyy, max_unit, max_unit_arr, min_unit, min_unit_arr, avg_vol, rate)` +
	`VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)` +
	` ON CONFLICT (code_id, unit_type, yyyy) DO UPDATE SET ` +
	`  max_unit=$4, max_unit_arr=$5, min_unit=$6, min_unit_arr=$7, avg_vol=$8, rate=$9`

func InsertTbYear(item model.CodeUnit) error {

	client := db.Conn
	stmt, err := client.Prepare(query_insert_tb_year)
	if err != nil {
		log.Println("쿼리:Prepare 오류: ", item)
		log.Fatal(err)
		panic(err)
	}
	defer stmt.Close()
	//code_id, unit_type, yyyy,  max_unit, max_unit_arr, min_unit, min_unit_arr, avg_vol, rate
	for _, v := range item.List {
		rate_json, jerr := json.Marshal(v.Rate)
		if jerr != nil {

			log.Printf("josn 변환 오류 %+v ", item.Code)
			log.Printf("%+v ", v)
			log.Fatal(" josn 변환 오류", v.Rate)
			log.Panic(jerr)
		}

		_, err = stmt.Exec(
			item.Code.Id, v.Unit, v.Year, v.Max, pq.Array(v.Up), v.Min, pq.Array(v.Down), v.Avg, rate_json,
		)
		if err != nil {
			log.Println("쿼리:stmt.Exec 오류: ")
			log.Printf("%+v ", v)
			log.Fatal(err)
			panic(err)
		}
	}

	return err
}

func SelectTbYear(item model.CodeInfo, unit_type int) ([]model.CodeYear, error) {
	var res []model.CodeYear
	query := `
	SELECT code_id, unit_type, yyyy, max_unit, max_unit_arr, min_unit, min_unit_arr, avg_vol, rate
	FROM project_trading_volume.tb_year
	WHERE	code_id =  $1 and unit_type = $2 
	order by yyyy asc
		
`
	//log.Println(query)
	rows, err := db.Conn.QueryContext(context.Background(), query, item.Code.Id, unit_type)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := model.CodeYear{}
		var max_arr pq.Int64Array
		var min_arr pq.Int64Array
		var jsonData string

		//code_id, unit_type, yyyy, max_unit, max_unit_arr, min_unit, min_unit_arr, avg_vol, rate
		if err := rows.Scan(&item.Code.Id, &item.UnitType,
			&item.UnitByYear.Year, &item.UnitByYear.Max, &max_arr,
			&item.UnitByYear.Min, &min_arr,
			&item.UnitByYear.Avg, &jsonData); err != nil {
			log.Fatal(err)
			panic(err)
		}
		for _, v := range max_arr {
			item.UnitByYear.Up = append(item.UnitByYear.Up, int(v))
		}
		for _, v := range min_arr {
			item.UnitByYear.Up = append(item.UnitByYear.Down, int(v))
		}
		json.Unmarshal([]byte(jsonData), &item.UnitByYear.Rate)
		//log.Printf("%+v ", item.UnitByYear.Rate)

		res = append(res, item)

	}
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	return res, err
}

const query_insert_tb_total = `INSERT INTO  project_trading_volume.tb_total (` +
	`code_id, unit_type, yyyy_cnt, max_unit, max_percent, min_unit, min_percent, max_rate, min_rate, max_arr_rate, min_arr_rate, avg_vol, last_updated)` +
	`VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13 )` +
	` ON CONFLICT (code_id, unit_type) DO UPDATE SET ` +
	`  yyyy_cnt=$3, max_unit=$4, max_percent=$5, min_unit=$6, min_percent=$7, max_rate=$8, min_rate=$9, max_arr_rate=$10, min_arr_rate=$11, avg_vol=$12, last_updated=$13 `

func InsertTbTotal(list []model.CodeTotal) error {

	client := db.Conn
	stmt, err := client.Prepare(query_insert_tb_total)
	if err != nil {
		log.Println("쿼리:Prepare 오류: ")
		log.Fatal(err)
		panic(err)
	}
	defer stmt.Close()
	//code_id, unit_type,  max_unit, sum_val
	for _, v := range list {

		max_rate, jerr := json.Marshal(v.UnitByTotal.MaxRate)
		if jerr != nil {

			log.Printf("josn 변환 오류 %+v ", v.Code)
			log.Printf("%+v ", v)
			log.Fatal(" josn 변환 오류", v.UnitByTotal.MaxRate)
			log.Panic(jerr)
		}
		min_rate, jerr := json.Marshal(v.UnitByTotal.MinRate)
		if jerr != nil {

			log.Printf("josn 변환 오류 %+v ", v.Code)
			log.Printf("%+v ", v)
			log.Fatal(" josn 변환 오류", v.UnitByTotal.MinRate)
			log.Panic(jerr)
		}
		max_arr_rate, jerr := json.Marshal(v.UnitByTotal.MaxArrRate)
		if jerr != nil {

			log.Printf("josn 변환 오류 %+v ", v.Code)
			log.Printf("%+v ", v)
			log.Fatal(" josn 변환 오류", v.UnitByTotal.MaxArrRate)
			log.Panic(jerr)
		}
		min_arr_rate, jerr := json.Marshal(v.UnitByTotal.MinArrRate)
		if jerr != nil {

			log.Printf("josn 변환 오류 %+v ", v.Code)
			log.Printf("%+v ", v)
			log.Fatal(" josn 변환 오류", v.UnitByTotal.MinArrRate)
			log.Panic(jerr)
		}
		//code_id, unit_type, yyyy_cnt, max_unit, max_percent,
		//min_unit, min_percent, max_rate, min_rate, max_arr_rate,
		// min_arr_rate, avg_vol, last_updated)` +
		_, err = stmt.Exec(
			v.Code.Id, v.UnitType, v.UnitByTotal.YearCnt, v.UnitByTotal.MaxUnit, v.UnitByTotal.MaxPercent,
			v.UnitByTotal.MinUnit, v.UnitByTotal.MinPercent, max_rate, min_rate, max_arr_rate,
			min_arr_rate, v.UnitByTotal.Avg, v.UnitByTotal.LastUpdated,
		)
		if err != nil {
			log.Println("쿼리:stmt.Exec 오류: ")
			log.Printf("%+v ", v)
			log.Fatal(err)
			panic(err)
		}
	}

	return err
}
