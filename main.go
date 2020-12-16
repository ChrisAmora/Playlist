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
	jwtSecret := viper.GetString(`jwt.secret`)
	a.Initialize(dbUser, dbPass, dbName, jwtSecret)
	a.RunGraphql(port)

}
