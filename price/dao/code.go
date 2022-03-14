package dao

import (
	"log"

	"github.com/cheolgyu/base/db"
	"github.com/cheolgyu/model"
	"github.com/cheolgyu/stat/price/c"
	mod_code "github.com/cheolgyu/tb/code"
	mod_config "github.com/cheolgyu/tb/config"
)

func Update_info() {
	query := `INSERT INTO public.info( name, updated) VALUES ('`
	query += c.INFO_NAME_UPDATED
	query += `', now()) ON CONFLICT ("name") DO UPDATE SET  updated= now()  `

	_, err := db.Conn.Exec(query)
	if err != nil {
		log.Fatalln(err, query)
		panic(err)
	}
}

func GetCodeAll() ([]model.Code, error) {
	res, err := mod_code.GetCodeList(db.Conn)
	return res, err
}

func GetConfigListByUpperCode() ([]model.Config, error) {
	res, err := mod_config.GetConfigListByUpperCode(db.Conn, mod_config.CONFIG_UPPER_CODE_PRICE_TYPE)
	return res, err
}
