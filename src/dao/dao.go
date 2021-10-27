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
	select
		mc.id,
		mc.code,
		mc.code_type,
	
		coalesce(tr.last_updated, 00000000) as last_updated
	from meta.code mc left join project_trading_volume.tb_total tr on mc.id = tr.code_id 
	where 
		mc.code_type=1
	
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

		if err := rows.Scan(&item.Code.Id, &item.Code.Code, &item.Code.Code_type, &item.LastUpdated); err != nil {
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
	`code_id, unit_type, yyyy, unit, sum_val)` +
	`VALUES ($1, $2, $3, $4, $5)` +
	` ON CONFLICT (code_id, unit_type, yyyy, unit) DO UPDATE SET ` +
	`   sum_val=$5`

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
			v.Code.Id, v.UnitType, v.Year, v.Unit, v.Sum,
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

const query_insert_tb_year = `INSERT INTO  project_trading_volume.tb_year (` +
	`code_id, unit_type, yyyy, max_unit, max_unit_arr, min_unit, min_unit_arr, avg_vol, percent)` +
	`VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)` +
	` ON CONFLICT (code_id, unit_type, yyyy) DO UPDATE SET ` +
	`  max_unit=$4, max_unit_arr=$5, min_unit=$6, min_unit_arr=$7, avg_vol=$8, percent=$9`

func InsertTbYear(item model.CodeUnit) error {

	client := db.Conn
	stmt, err := client.Prepare(query_insert_tb_year)
	if err != nil {
		log.Println("쿼리:Prepare 오류: ", item)
		log.Fatal(err)
		panic(err)
	}
	defer stmt.Close()
	//code_id, unit_type, yyyy,  max_unit, max_unit_arr, min_unit, min_unit_arr, avg_vol, percent
	for _, v := range item.List {
		percent_json, jerr := json.Marshal(v.Percent)
		if jerr != nil {

			log.Printf("josn 변환 오류 %+v ", item.Code)
			log.Printf("%+v ", v)
			log.Fatal(" josn 변환 오류", v.Percent)
			log.Panic(jerr)
		}

		_, err = stmt.Exec(
			item.Code.Id, v.Unit, v.Year, v.Max, pq.Array(v.Up), v.Min, pq.Array(v.Down), v.Avg, percent_json,
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
