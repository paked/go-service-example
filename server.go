package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/paked/configure"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

var (
	DB *runner.DB

	conf      = configure.New()
	dbName    = conf.String("db-name", "postgres", "DB_NAME")
	dbUser    = conf.String("db-user", "postgres", "DB_USER")
	dbPass    = conf.String("db-pass", "postgres", "DB_PASS")
	dbService = conf.String("db-service", "jarvis", "DB_SERVICE")
	dbPort    = conf.String("db-port", "5432", "DB_PORT")
)

func db() {
	db, err := sql.Open("postgres",
		fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
			*dbUser,
			*dbPass,
			*dbService,
			*dbPort,
			*dbName,
		),
	)

	if err != nil {
		fmt.Println("lol")
		panic(err)
	}

	runner.MustPing(db)

	runner.LogQueriesThreshold = 10 * time.Millisecond
	DB = runner.NewDB(db, "postgres")

	DB.DB.MustExec(query)
}

var query = `
CREATE TABLE IF NOT EXISTS pings (
	id serial PRIMARY KEY,
	message text not null
)
`

type Ping struct {
	ID      int64  `db:"id"`
	Message string `db:"message"`
}

func main() {
	// fmt.Println("party")
	conf.Use(configure.NewEnvironment())
	conf.Use(configure.NewFlag())

	conf.Parse()

	db()

	fmt.Println("Welcome to pay.")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "Welcome to vision (the backend)")

		var n int64
		err := DB.
			Select("count(*)").
			From("pings").
			QueryScalar(&n)

		if err != nil {
			fmt.Println("FAILED!")
			return
		}

		fmt.Fprintf(w, "You are visitor #%v. Congratulations!", n)

		var ping Ping
		err = DB.
			InsertInto("pings").
			Columns("message").
			Values("hello!").
			Returning("id", "message").
			QueryStruct(&ping)

		if err != nil {
			fmt.Println("FAILED NOW!")
			return
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
