package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/cheolgyu/stock-write-project-trading-volume/src/c"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Conn *sqlx.DB

func init() {
	pg := PQ{}
	Conn = pg.conn_sqlx()
	Conn.SetMaxIdleConns(c.DB_MAX_CONN)
	Conn.SetMaxOpenConns(c.DB_MAX_CONN)
}

type PQ struct {
}

func (o *PQ) conn() *sql.DB {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Panic("Error loading .env file")
	}
	DB_URL := os.Getenv("DB_URL")

	log.Println("============================")
	log.Println(DB_URL)

	db, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Println("커넥션 오류:", err)
		panic(err)
	}
	return db
}

func (o *PQ) conn_sqlx() *sqlx.DB {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Panic("Error loading .env file")
	}
	DB_URL := os.Getenv("DB_URL")

	log.Println("============================")
	log.Println(DB_URL)

	db, err := sqlx.Connect("postgres", DB_URL)
	if err != nil {
		log.Println("커넥션 오류:", err)
		panic(err)
	}
	return db
}
