package dao

import (
	"context"
	"log"

	"github.com/cheolgyu/stock-write-project-trading-volume/src/db"
	"github.com/cheolgyu/stock-write-project-trading-volume/src/model"
)

type PreprocessDao struct {
}

func (o *PreprocessDao) GetCodeList() ([]model.CodeInfo, error) {
	var res []model.CodeInfo
	query := `
select 
	mc.id,mc.code, mc.code_type
	,coalesce(max( hp.dt),19560303)::text as start_dt 
	--,to_char( now(), 'YYYYMMDD')::text as end_dt
	,'20210803' as end_dt
 from meta.code mc left join trading_vol.tb_result tv_res on mc.id = tv_res.code_id
 where mc.code_type = 2
 group by mc.id, mc.code_type
`
	log.Println(query)
	rows, err := db.Conn.QueryContext(context.Background(), query)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		item := model.CodeInfo{}

		if err := rows.Scan(&item.Id, &item.Code, &item.Code_type); err != nil {
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
