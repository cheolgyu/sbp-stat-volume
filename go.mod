module github.com/cheolgyu/stock-write-project-trading-volume

go 1.16

require (
	github.com/cheolgyu/stock-write-common v0.0.0
	github.com/gchaincl/dotsql v1.0.0 // indirect
	github.com/jmoiron/sqlx v1.3.4
	github.com/joho/godotenv v1.4.0
	github.com/lib/pq v1.10.3
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/swithek/dotsqlx v1.0.0 // indirect
)

replace github.com/cheolgyu/stock-write-common v0.0.0 => ../stock-write-common
