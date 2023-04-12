package config

import (
	"bytes"
	"embed"
	"github.com/spf13/viper"
)

// ConfEnv env环境变量
var ConfEnv string = "test"

type (
	Conf struct {
		App App `yaml:"app"`
		Log Log `yaml:"log"`
	}
	App struct {
		Mode            string `yaml:"mode" default:"debug"`
		DefaultPageSize int    `yaml:"defaultPageSize" default:"10"`
		MaxPageSize     int    `yaml:"maxPageSize" default:"500"`
		AppName         string `yaml:"appName" default:""`
		Version         string `yaml:"version" default:""`
		Description     string `yaml:"description" default:""`
		VersionNewMsg   string `yaml:"versionNewMsg" default:""`
		VersionOldMsg   string `yaml:"versionOldMsg" default:""`
		BtnConfirmText  string `yaml:"btnConfirmText" default:"确定"`
		BtnCancelText   string `yaml:"btnCancelText" default:"取消"`
		Width           int    `yaml:"width" default:"1024"`
		Height          int    `yaml:"height" default:"768"`
		UpgradeUrl      string `yaml:"upgradeUrl" default:""`
		HomeUrl         string `yaml:"homeUrl" default:""`
		ServerUrl       string `yaml:"serverUrl" default:""`
		DbName          string `yaml:"dbName" default:""`
		WsUrl           string `yaml:"wsUrl" default:""`
		PushUrl         string `yaml:"pushUrl" default:""`
	}
	Log struct {
		Debug    bool   `yaml:"debug" default:"true"`
		FileName string `yaml:"fileName" default:"syncstock"`
		DirPath  string `yaml:"dirPath" default:"runtime/logs"`
	}
)

//go:embed yaml
var yamlCfg embed.FS

// InitConfig 初始化配置
func InitConfig() (*Conf, error) {
	var cfg *Conf
	v := viper.New()
	v.SetConfigType("yaml")
	yamlConf, _ := yamlCfg.ReadFile("yaml/config." + ConfEnv + ".yaml")
	if err := v.ReadConfig(bytes.NewBuffer(yamlConf)); err != nil {
		return nil, err
	}
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
