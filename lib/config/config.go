package configmanager

import (
	"io/ioutil"
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

// Global GlobalConfig 全Service可使用的參數
var Global *Configuration

// LogSetting 是 log 的設定
type LogSetting struct {
	Name             string `yaml:"name"`
	Type             string `yaml:"type"`
	MinLevel         string `yaml:"min_level"`
	ConnectionString string `yaml:"connection_string"`
}

// Database 用來提供連線的資料庫數據
type Database struct {
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	DBName   string `yaml:"dbname"`
}

type Mysql struct {
	Address  string `yaml:"address"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
type Redis struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

//Configuration 參數結構
type Configuration struct {
	//WSServer jiface.IServer
	Env   string `yaml:"env"`
	Debug bool   `yaml:"debug"`

	Api struct {
		HTTPBind string `yaml:"http_bind"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Mysql    Mysql  `yaml:"mysql"`
		Redis    Redis  `yaml:"redis"`
	} `yaml:"api_groups"`

	Job struct {
		CrontabSetting string `yaml:"crontab_setting"`
		Restart        bool   `yaml:"restart"`
	} `yaml:"trigger_groups"`
}

//Reload 重新載入參數
func Reload() *Configuration {
	data, err := ioutil.ReadFile("configs/config.yml")
	if err != nil {
		log.Panic().Msgf("%v", err)
	}

	tempPara := &Configuration{}

	err = yaml.Unmarshal(data, &tempPara)
	if err != nil {
		log.Panic().Msgf("%v", err)
	}

	if value, ok := os.LookupEnv("MYSQL_ADDRESS"); ok {
		tempPara.Api.Mysql.Address = value
	}
	if value, ok := os.LookupEnv("MYSQL_USER"); ok {
		tempPara.Api.Mysql.Username = value
	}
	if value, ok := os.LookupEnv("MYSQL_PASSWORD"); ok {
		tempPara.Api.Mysql.Password = value
	}

	if value, ok := os.LookupEnv("REDIS_ADDRESS"); ok {
		tempPara.Api.Redis.Address = value
	}
	if value, ok := os.LookupEnv("REDIS_PASSWORD"); ok {
		tempPara.Api.Redis.Password = value
	}
	if value, ok := os.LookupEnv("REDIS_DATABASE"); ok {
		redisDB, err := strconv.Atoi(value)
		if err == nil {
			tempPara.Api.Redis.Database = redisDB
		}
	}

	if value, ok := os.LookupEnv("CRONTAB_SETTING"); ok {
		tempPara.Job.CrontabSetting = value
	}

	log.Info().Msgf("%v", tempPara)

	return tempPara
}
