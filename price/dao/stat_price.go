package dao

import (
	"fmt"

	"github.com/cheolgyu/base/db"
	"github.com/cheolgyu/base/logging"
	cmm_model "github.com/cheolgyu/model"
	"github.com/cheolgyu/stat/price/model"
)

const query_insert = `INSERT INTO stat.price( ` +
	` code_id, price_type, p1x_Left, p1x, p1y, p2x, p2y, p3x, p3y, p3_type, p32y_percent) ` +
	` VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11); `

const query_select_list = `
with tmp as (select TO_DATE(max(dt::text),'YYYYMMDD') - 365  as before_year
			 , TO_DATE(max(dt::text),'YYYYMMDD') as max_dt  from hist.price where code_id = $1)
select  code_id, dt ,op, hp, lp, cp
	,tmp.max_dt - TO_DATE(dt::text,'YYYYMMDD') as day_cnt
from hist.price hp , tmp
where code_id = $1 and TO_DATE(dt::text,'YYYYMMDD')  > tmp.before_year 
order by dt desc
`

/*
	테이블 초기화
*/
func Initail_table() {
	query := `TRUNCATE table stat.price`

	_, err := db.Conn.Exec(query)
	if err != nil {
		logging.Log.Fatalln(err, query)
		panic(err)
	}
}

/*
	조회 및 가격별 분류 및 반환
*/
func SelectList(code_id int) (model.CodeInfo, error) {

	var res model.CodeInfo
	rows, err := db.Conn.Query(query_select_list, code_id)
	if err != nil {
		logging.Log.Fatalln(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var item model.Res
		// hp.code_id, hp.dt,  hp.op, hp.hp, hp.lp, hp.cp
		if err := rows.Scan(&item.Code.Id,
			&item.PriceMarket.Dt, &item.PriceMarket.OpenPrice, &item.PriceMarket.HighPrice, &item.PriceMarket.LowPrice, &item.PriceMarket.ClosePrice, &item.DayCnt); err != nil {
			logging.Log.Fatal(err)
			panic(err)
		}

		convert := item.ByPrice()
		res.Code.Id = item.Code.Id
		res.OP = append(res.OP, convert[0])
		res.CP = append(res.CP, convert[1])
		res.LP = append(res.LP, convert[2])
		res.HP = append(res.HP, convert[3])

	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		logging.Log.Fatal(err)
		panic(err)
	}

	return res, err
}

func Insert(list []cmm_model.TbStatPrice) error {

	client := db.Conn
	stmt, err := client.Prepare(query_insert)
	if err != nil {
		logging.Log.Println("쿼리:Prepare 오류: ")
		logging.Log.Fatal(err)
		panic(err)
	}
	defer stmt.Close()

	for _, item := range list {

		//code_id, price_type, p1x_unit_type, p1x_unit, p1x, p1y, p2x, p2y, p3x, p3y, p3_type, p32y_percent
		_, err := stmt.Exec(
			item.Code_id, item.Price_type, item.P1x_Left,
			item.P1.X, item.P1.Y, item.P2.X, item.P2.Y, item.P3.X, item.P3.Y,
			item.P3_type, item.P32y_percent,
		)
		if err != nil {
			err_item := fmt.Sprintf("%+v", item)
			logging.Log.Println("쿼리:InsertCompanyState:stmt.Exec 오류: ", err_item)
			logging.Log.Fatal(err)
			panic(err)
		}
	}
	return err
}
