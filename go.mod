module github.com/cheolgyu/stock-write-project-trading-volume

go 1.16

require (
	github.com/cheolgyu/stock-write-common v0.0.0
	github.com/lib/pq v1.10.3
)

replace github.com/cheolgyu/stock-write-common v0.0.0 => ../stock-write-common
