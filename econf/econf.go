package econf

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

var once sync.Once
var config *econf

type econf struct {
	viperConfMap map[string]*viper.Viper
	confEnvPath  string
}

func GetInstance(path string) *econf {
	once.Do(func() {
		config = &econf{
			confEnvPath: path,
		}
	})
	return config
}

func (e *econf) ParseConfig(path string, conf interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Open config %v fail, %v", path, err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Read config fail, %v", err)
	}

	v := viper.New()
	v.SetConfigType("toml")
	v.ReadConfig(bytes.NewBuffer(data))
	if err := v.Unmarshal(conf); err != nil {
		return fmt.Errorf("Parse config fail, config:%v, err:%v", string(data), err)
	}
	fmt.Println(v.GetString("base.debug_mode"))
	return nil
}

// InitViperConf 初始化配置文件
func (e *econf) InitViperConf() error {
	f, err := os.Open(e.confEnvPath + "/")
	if err != nil {
		return err
	}

	fileList, err := f.Readdir(1024)
	if err != nil {
		return err
	}
	for _, f0 := range fileList {
		if !f0.IsDir() {
			bts, err := ioutil.ReadFile(e.confEnvPath + "/" + f0.Name())
			if err != nil {
				return err
			}
			v := viper.New()
			pathArr := strings.Split(f0.Name(), ".")
			v.SetConfigType(pathArr[1])
			v.ReadConfig(bytes.NewBuffer(bts))

			if e.viperConfMap == nil {
				e.viperConfMap = make(map[string]*viper.Viper)
			}
			//debug_mode := v.GetString("debug_mode")
			e.viperConfMap[pathArr[0]] = v
		}
	}
	return nil
}

func (e *econf) getViper(key string) (*viper.Viper, []string, error) {
	keys := strings.Split(key, ".")
	if len(keys) < 2 {
		return nil, keys, errors.New("incorrect incoming parameters")
	}
	v, ok := e.viperConfMap[keys[0]]
	if !ok {
		return nil, keys, errors.New(fmt.Sprintf("no configuration file corresponding to '%s' was found", keys[0]))
	}
	return v, keys, nil
}

// 获取get配置信息
func (e *econf) GetStringSliceConf(key string) []string {
	//keys := strings.Split(key, ".")
	//if len(keys) < 2 {
	//	return nil
	//}
	//v := e.viperConfMap[keys[0]]
	v, keys, err := e.getViper(key)
	if err != nil {
		return nil
	}
	conf := v.GetStringSlice(strings.Join(keys[1:len(keys)], "."))
	return conf
}

// 获取get配置信息
func (e *econf) GetStringConf(key string) string {
	v, keys, err := e.getViper(key)
	if err != nil {
		return ""
	}

	confString := v.GetString(strings.Join(keys[1:len(keys)], "."))
	return confString
}
