package main

import (
	_ "expvar"
	"net/http"
	"os"

	"Posting-Feed-be/transport"

	"github.com/go-kit/kit/log"

	_ "github.com/lib/pq"
)

func main() {

	logger := log.NewLogfmtLogger(os.Stdout)

	transport.RegisterHttpsServicesAndStartListener()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}

	logger.Log("listening-on", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Log("listen.error", err)
	}
	/*
		// test insert
		insertStmt := `insert into feeds values(1, 1234, 'namaFile.png', 'isi caption', current_timestamp)`
		_, e := db.Exec(insertStmt)
		CheckError(e)
	*/
}
