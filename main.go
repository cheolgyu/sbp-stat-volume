package main

import (
	_ "github.com/cheolgyu/sbm-base/db"
	_ "github.com/cheolgyu/sbm-base/env"
	"github.com/cheolgyu/sbm-base/logging"
	"github.com/cheolgyu/sbp-stat-volume/src/handler"
)

func main() {
	defer logging.ElapsedTime()()
	project_run()
}

func project_run() {
	handler.GetCodeList()
}
