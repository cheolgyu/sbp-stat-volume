module github.com/cheolgyu/stat

go 1.16

require (
	github.com/cheolgyu/tb v0.0.0
	github.com/cheolgyu/base v0.0.0
	github.com/cheolgyu/model v0.0.0
	github.com/lib/pq v1.10.3
)

replace (
	github.com/cheolgyu/tb v0.0.0 => ../tb
	github.com/cheolgyu/base v0.0.0 => ../base
	github.com/cheolgyu/model v0.0.0 => ../model
)
