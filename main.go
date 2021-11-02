package main

import (
	_ "github.com/cheolgyu/stock-write-common/src/db"
	_ "github.com/cheolgyu/stock-write-common/src/env"
	"github.com/cheolgyu/stock-write-common/src/logging"
	"github.com/cheolgyu/stock-write-project-trading-volume/src/handler"
)

func main() {
	defer logging.ElapsedTime()()
	project_run()
}

func project_run() {
	handler.GetCodeList()
}
