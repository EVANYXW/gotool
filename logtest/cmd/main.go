package main

import (
	"demo/tools/elog/zlog"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	//elog.GetInstance()
	//log := elog.GetLogger().GetCtx(context.Background())

	//log.Error("START")

	if _, err := kingpin.CommandLine.Parse([]string{
		"--log.level", "debug",
		"--log.format", "logfmt",
		"--log.path", "./logs",
		"--log.filename", "error.log",
		"--log.file-max-size", "3",
		"--log.file-max-backups", "2"}); err != nil {
		fmt.Println(err)
	}
	if err := zlog.InitLogger(); err != nil {
		fmt.Println(err)
	}
	zlog.Logger.Debug("debug")
}
