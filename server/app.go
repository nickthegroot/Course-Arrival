package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App default
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password string) {
	connectionString :=
	fmt.Sprintf("postgres://%s:%s@free-tier.gcp-us-central1.cockroachlabs.cloud:26257/defaultdb?sslmode=require", user, password)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) createCourseGraph(w http.ResponseWriter, r *http.Request) {
	var c course
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		respondWithError(w, http.StatusBadRequest, "Body must be a valid JSON object")
	}

	fmt.Println(c)

	defer r.Body.Close()

	courses, err := c.getPreqs(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
	}

	respondWithJSON(w, http.StatusOK, courses)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func (a *App) initializeRoutes() {
    a.Router.HandleFunc("/graph", a.createCourseGraph).Methods("POST")
}