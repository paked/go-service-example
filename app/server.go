package main

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/paked/configure"
	"github.com/paked/pay/models"
	"github.com/paked/restrict"
)

var (
	conf = configure.New()

	dbName    = conf.String("db-name", "postgres", "DB_NAME")
	dbUser    = conf.String("db-user", "postgres", "DB_USER")
	dbPass    = conf.String("db-pass", "postgres", "DB_PASS")
	dbService = conf.String("db-service", "jarvis", "DB_SERVICE")
	dbPort    = conf.String("db-port", "5432", "DB_PORT")

	crypto = conf.String("crypto", "/crypto/app.rsa", "Your crypto")
)

type Ping struct {
	ID      int64  `db:"id"`
	Message string `db:"message"`
}

func main() {
	conf.Use(configure.NewEnvironment())
	conf.Use(configure.NewFlag())

	conf.Parse()

	restrict.ReadCryptoKey(*crypto)

	models.Init(
		*dbUser,
		*dbPass,
		*dbService,
		*dbPort,
		*dbName,
	)

	fmt.Println("Welcome to pay.")

	r := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(true)

	r.HandleFunc("/ping", restrict.R(pingHandler)).
		Methods("GET")

	r.HandleFunc("/users", registerHandler).
		Methods("POST")
	r.HandleFunc("/users", loginHandler).
		Methods("GET")

	fmt.Println(http.ListenAndServe(":8080", r))
}
func pingHandler(w http.ResponseWriter, r *http.Request, t *jwt.Token) {
	fmt.Fprintln(w, "Welcome to vision (the backend)")

	var n int64
	err := models.DB.
		Select("count(*)").
		From("pings").
		QueryScalar(&n)

	if err != nil {
		fmt.Println("FAILED!")
		return
	}

	fmt.Fprintf(w, "You are visitor #%v. Congratulations!", n)

	var ping Ping
	err = models.DB.
		InsertInto("pings").
		Columns("message").
		Values("hello!").
		Returning("id", "message").
		QueryStruct(&ping)

	if err != nil {
		fmt.Println("FAILED NOW!")
		return
	}
}
