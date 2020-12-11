package main

import (
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	a := App{}
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	port := viper.GetString(`server.address`)
	a.Initialize(dbUser, dbPass, dbName)
	a.RunGraphql(port)

}
