package main

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type Config struct {
	Debug  bool
	Server struct {
		Address string
	}
	Jwt struct {
		Secret string
	}
	Context struct {
		Timeout int64
	}
	Database struct {
		Host string
		Port int64
		User string
		Pass string
		Name string
	}
}

func GetConf() *Config {
	conf := &Config{}
	err := viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	return conf
}

func main() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	a := App{}
	if err != nil {
		panic(err)
	}
	a.Initialize()
	a.RunGraphql()

}
