package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var schema = `
CREATE TABLE IF NOT EXISTS music (
	id SERIAL NOT NULL PRIMARY KEY,
	title text,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
`

type Music struct {
	ID        int64     `db:"id"`
	Title     string    `db:"title"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("betinho")
}

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	if err != nil {
		panic(err)
	}

	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPass, dbName)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	db.MustExec(schema)
	musics := []Music{}
	err = db.Select(&musics, "SELECT * FROM music")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(musics)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/posts", getPosts).Methods("GET")
	http.ListenAndServe(":8000", router)
}
