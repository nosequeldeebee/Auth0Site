package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/auth0-samples/auth0-golang-web-app/01-Login/routes/callback"
	"github.com/auth0-samples/auth0-golang-web-app/01-Login/routes/home"
	"github.com/auth0-samples/auth0-golang-web-app/01-Login/routes/middlewares"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/nosequeldeebee/sandbox/auth0/01-login/routes/user"
)

type hospital struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
}

var hospitals []hospital

func StartServer() {

	//for db authentication
	username := flag.String("u", "", "db username")
	pwd := flag.String("p", "", "db password")
	flag.Parse()

	//connect to db
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=test sslmode=disable", *username, *pwd)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal("couldn't connect to db", err)
	}
	defer db.Close()

	//perform query and populate `hospitals` slice
	rows, err := db.Query("SELECT * FROM teachinghospital")
	if err != nil {
		log.Fatal("couldn't query db", err)
	}
	defer rows.Close()

	var (
		name, address, city, state, zip string
	)

	for rows.Next() {
		err := rows.Scan(&name, &address, &city, &state, &zip)
		if err != nil {
			log.Fatal("couldn't scan db results", err)
		}
		hospitals = append(hospitals, hospital{Name: name, Address: address, City: city, State: state, Zip: zip})
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	//server initializations
	r := mux.NewRouter()

	r.HandleFunc("/", home.HomeHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	))

	r.Handle("/hospitals", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(getHospitals)),
	))

	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.Handle("/", r)
	http.ListenAndServe(":9090", nil)
}

func getHospitals(w http.ResponseWriter, r *http.Request) {
	//set headers for CORS
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}

	//prettify api endpoint
	b, err := json.MarshalIndent(hospitals, "", " ")
	if err != nil {
		fmt.Println("couldn't indent json", err)
	}
	io.WriteString(w, string(b))
}
