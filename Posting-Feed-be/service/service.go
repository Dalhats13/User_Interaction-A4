package service

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbhost = "103.157.96.115"
	dbport = "5432"
	//dbuser = "sadhelx_usr"
	//dbpass = "s4dhelx"
	//dbname = "sdx_usermgmt_db"
	dbuser = "jarijari"
	dbpass = "pinwheel"
	dbname = "db_jarijari"
)

func initDb() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

func dbConfig() map[string]string {
	conf := make(map[string]string)

	conf[dbhost] = "103.157.96.115"
	conf[dbport] = "5432"
	//conf[dbuser] = "sadhelx_usr"
	//conf[dbpass] = "s4dhelx"
	//conf[dbname] = "sdx_usermgmt_db"
	conf[dbuser] = "jarijari"
	conf[dbpass] = "pinwheel"
	conf[dbname] = "db_jarijari"
	return conf
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Create(image_feed string, caption_feed string) string {
	initDb()
	defer db.Close()
	var response string

	insertStmt := `insert into feeds values(nextval('seq_feed_id'), 987654321, $1, $2, current_timestamp)`
	_, e := db.Exec(insertStmt, image_feed, caption_feed)
	if e != nil {
		panic(e)
		response = fmt.Sprintf("There's something wrong.")
	} else {
		response = fmt.Sprintf("Feed successfully posted.")
	}

	return response
}
