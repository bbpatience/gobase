package yml_config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gobase/global/my_errors"
	"gobase/global/variables"
	"gobase/utils/yml_config/ymlconfig_interf"
	"log"
	"time"
)

// 由于 vipver 包本身对于文件的变化事件有一个bug，相关事件会被回调两次
// 常年未彻底解决，相关的 issue 清单：https://github.com/spf13/viper/issues?q=OnConfigChange
// 设置一个内部全局变量，记录配置文件变化时的时间点，如果两次回调事件事件差小于1秒，我们认为是第二次回调事件，而不是人工修改配置文件
// 这样就避免了 vipver 包的这个bug

var lastChangeTime time.Time

func init() {
	lastChangeTime = time.Now()
}

// 创建一个yaml配置文件工厂
// 参数设置为可变参数的文件名，这样参数就可以不需要传递，如果传递了多个，我们只取第一个参数作为配置文件名
func CreateYamlFactory(fileName ...string) ymlconfig_interf.YmlConfigInf {

	yamlConfig := viper.New()
	// 配置文件所在目录
	yamlConfig.AddConfigPath("./config")
	// 需要读取的文件名,默认为：config
	if len(fileName) == 0 {
		yamlConfig.SetConfigName("config")
	} else {
		yamlConfig.SetConfigName(fileName[0])
	}
	//设置配置文件类型(后缀)为 yml
	yamlConfig.SetConfigType("yml")

	if err := yamlConfig.ReadInConfig(); err != nil {
		log.Fatal(my_errors.ErrorsConfigInitFail + err.Error())
	}

	return &ymlConfig{
		yamlConfig,
	}
}

type ymlConfig struct {
	viper *viper.Viper
}

//监听文件变化
func (y *ymlConfig) ConfigFileChangeListen() {
	y.viper.OnConfigChange(func(changeEvent fsnotify.Event) {
		if time.Now().Sub(lastChangeTime).Seconds() >= 1 {
			if changeEvent.Op.String() == "WRITE" {
				lastChangeTime = time.Now()
			}
		}
	})
	y.viper.WatchConfig()
}

// 允许 clone 一个相同功能的结构体
func (y *ymlConfig) Clone(fileName string) ymlconfig_interf.YmlConfigInf {
	// 这里存在一个深拷贝，需要注意，避免拷贝的结构体操作对原始结构体造成影响
	var ymlC = *y
	var ymlConfViper = *(y.viper)
	(&ymlC).viper = &ymlConfViper

	(&ymlC).viper.SetConfigName(fileName)
	if err := (&ymlC).viper.ReadInConfig(); err != nil {
		variables.ZapLog.Error(my_errors.ErrorsConfigInitFail, zap.Error(err))
	}
	return &ymlC
}

// Get 一个原始值
func (y *ymlConfig) Get(keyName string) interface{} {

	return y.viper.Get(keyName)
}

// GetString
func (y *ymlConfig) GetString(keyName string) string {
	return y.viper.GetString(keyName)
}

// GetBool
func (y *ymlConfig) GetBool(keyName string) bool {
	return y.viper.GetBool(keyName)
}

// GetInt
func (y *ymlConfig) GetInt(keyName string) int {
	return y.viper.GetInt(keyName)
}

// GetInt32
func (y *ymlConfig) GetInt32(keyName string) int32 {
	return y.viper.GetInt32(keyName)
}

// GetInt64
func (y *ymlConfig) GetInt64(keyName string) int64 {
	return y.viper.GetInt64(keyName)
}

// float64
func (y *ymlConfig) GetFloat64(keyName string) float64 {
	return y.viper.GetFloat64(keyName)
}

// GetDuration
func (y *ymlConfig) GetDuration(keyName string) time.Duration {
	return y.viper.GetDuration(keyName)
}

// GetStringSlice
func (y *ymlConfig) GetStringSlice(keyName string) []string {
	return y.viper.GetStringSlice(keyName)
}
