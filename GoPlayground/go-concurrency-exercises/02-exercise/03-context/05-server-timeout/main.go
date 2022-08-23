package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func slowQuery() error {
	_, err := db.Exec("SELECT pg_sleep(5)")
	return err
}

func slowHandler(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	err := slowQuery()
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Fprintln(w, "OK")
	fmt.Printf("SlowHandler took: %v\n", time.Since(start))

}

func main() {
	fmt.Println("Inside MAin")
	var err error

	connstr := "host=localhost port=5433 user=postgres password=postgres dbname=test sslmode=disable"

	db, err = sql.Open("postgres", connstr)
	if err != nil {
		fmt.Println("Error in connection ")
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	/*
		// HTTP Handler functions are unaware of these timeouts they run to completion consuming resources.
		//Note: Start server with "time curl http://localhost:9999"
		srv := http.Server{
			Addr:         "localhost:9999",
			WriteTimeout: 2 * time.Second,
			Handler:      http.HandlerFunc(slowHandler),
		}
	*/

	//If input handler runs for longer than it's limit then handler sends the client a "503 service unavailable" error
	// and HTML message
	//Note: Start server with "time curl -i http://localhost:9999"
	//In Output we can see "503 Service Unavailable"
	//
	srv := http.Server{
		Addr:         "localhost:9999",
		WriteTimeout: 2 * time.Second,
		Handler:      http.TimeoutHandler(http.HandlerFunc(slowHandler), 1*time.Second, "Timeout!\n"), //http.HandlerFunc()
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("Server failed: %s\n", err)
	}
}

// --> Installing postgres - macos
// brew install postgresql

// --> start
// pg_ctl -D /usr/local/var/postgres start

// --> create db and user
// psql postgres
// CREATE DATABASE wonderland;
// CREATE USER alice WITH ENCRYPTED PASSWORD 'pa$$word';
// GRANT ALL PRIVILEGES ON DATABASE wonderland TO alice;

// --> stop
// pg_ctl -D /usr/local/var/postgres stop

// --> postgresql download link
// https://www.postgresql.org/download/

// start postgresql - Windows
// pg_ctl -D "C:\Program Files\PostgreSQL\13\data" start

// stop postgresql - Windows
// pg_ctl -D "C:\Program Files\PostgreSQL\13\data" stop

// --> Linux
// sudo apt-get update
// sudo apt-get install postgresql-13

// sudo -u postgres psql -c "ALTER USER alice PASSWORD 'pa$$word';"
// sudo -u postgres psql -c "CREATE DATABASE wonderland;"

// sudo service postgresql start

// sudo service postgresql stop
