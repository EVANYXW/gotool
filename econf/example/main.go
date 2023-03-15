package main

import (
	"fmt"
	"github.com/evanyxw/gotool/econf"
)

func main() {

	//file, _ := exec.LookPath(os.Args[0])
	//path, _ := filepath.Abs(file)
	//index := strings.LastIndex(path, string(os.PathSeparator))
	//
	//fmt.Println(path[:index])
	//fmt.Println(os.Getwd())

	conf := econf.GetInstance("../conf/dev")
	conf.InitViperConf()
	time_location := conf.GetStringConf("base.base.time_location")
	app_model := conf.GetStringConf("my.app_model")
	fmt.Println(time_location)
	fmt.Println(app_model)
}
