package tomlConfig

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
	"sync"
)

var (
	conf *SL
	once sync.Once
)

func TomlConfig(configFile string) *SL {
	//检测配置文件是否存在
	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		// 说明配置文件不存在，直接报错，退出
		fmt.Println("sorry, config file", configFile, "is not exist, err is", err)
		os.Exit(0)
	}
	//说明配置文件存在，获取其绝对路径
	absPath, err1 := filepath.Abs(configFile)
	if err1 != nil {
		fmt.Println("sorry get abs path of config file", configFile, "is bad, err is", err1)
		os.Exit(0)
	}
	//单例模式读取配置文件
	once.Do(func() {
		_, err := toml.DecodeFile(absPath, &conf)
		if err != nil {
			fmt.Println("toml config file is bad, err is", err)
			os.Exit(0)
		}
	})
	return conf
}
