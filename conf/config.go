package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ConfigStruct struct {
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	App     App     `mapstructure:"app" yaml:"app"`
	Log     Log     `mapstructure:"log" json:"log" yaml:"log"`
	Jwt     Jwt     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Storage Storage `mapstructure:"storage" json:"storage" yaml:"storage"`
}

var config = new(ConfigStruct)

type App struct {
	Env     string `yaml:"env"`
	Port    string `yaml:"port"`
	AppName string `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
	AppUrl  string `mapstructure:"app_url" yaml:"app_url"`
}

type Mysql struct {
	Driver              string `mapstructure:"driver" json:"driver" yaml:"driver"`
	Host                string `mapstructure:"host" json:"host" yaml:"host"`
	Port                int    `mapstructure:"port" json:"port" yaml:"port"`
	Database            string `mapstructure:"database" json:"database" yaml:"database"`
	UserName            string `mapstructure:"username" json:"username" yaml:"username"`
	Password            string `mapstructure:"password" json:"password" yaml:"password"`
	Charset             string `mapstructure:"charset" json:"charset" yaml:"charset"`
	MaxIdleConns        int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns        int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
	LogMode             string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`
	EnableFileLogWriter bool   `mapstructure:"enable_file_log_writer" json:"enable_file_log_writer" yaml:"enable_file_log_writer"`
	LogFilename         string `mapstructure:"log_filename" json:"log_filename" yaml:"log_filename"`
}

type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

type Log struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`
	RootDir    string `mapstructure:"root_dir" json:"root_dir" yaml:"root_dir"`
	Filename   string `mapstructure:"filename" json:"filename" yaml:"filename"`
	Format     string `mapstructure:"format" json:"format" yaml:"format"`
	ShowLine   bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"` // MB
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`    // day
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}

func Init() *viper.Viper {
	v := viper.New()
	v.AddConfigPath("./conf")
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}

	// 将配置赋值给全局变量
	if err := v.Unmarshal(&Global.ConfigStruct); err != nil {
		fmt.Println(err)
	}

	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&config); err != nil {
			fmt.Println(err)
		}
	})

	return v
}
