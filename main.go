package main

import (
	_ "github.com/cheolgyu/base/db"
	_ "github.com/cheolgyu/base/env"
	"github.com/cheolgyu/base/logging"
	"github.com/cheolgyu/stat/volume/handler"
)

func main() {
	defer logging.ElapsedTime()()
	project_run()
}

func project_run() {
	handler.GetCodeList()
}
