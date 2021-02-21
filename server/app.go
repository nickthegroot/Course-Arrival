package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/yourbasic/graph"
)

// App default
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, certPath string) {
	connectionString :=
	fmt.Sprintf("postgres://%s:%s@free-tier.gcp-us-central1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full&sslrootcert=%s&options=--cluster=modest-seal-865", user, password, certPath)

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

func (a *App) getPrerequisites(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, found := vars["id"]
	if !found {
		respondWithError(w, http.StatusBadRequest, "Invalid Course ID")
		return
	}

	c := course{ID: id}

	// decoder := json.NewDecoder(r.Body)
	// if err := decoder.Decode(&c); err != nil {
	// 	respondWithError(w, http.StatusBadRequest, "Body must be a valid JSON object")
	// }

	// defer r.Body.Close()

	courses, err := c.getPreqs(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	respondWithJSON(w, http.StatusOK, courses)
}

func (a *App) getFullPrerequisites(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, found := vars["id"]
	if !found {
		respondWithError(w, http.StatusBadRequest, "Invalid Course ID")
		return
	}

	startC := course{ID: id}
	startC.getCourse(a.DB)
	
	queue := []course{startC}
	courseMap := make(map[string]course)
	dependencyMap := make(map[string][]string)
	graphMap := make(map[string]int)
	inverseGraphMap := make(map[int]string)

	i := 0
	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]

		if _, exists := courseMap[next.ID]; exists {
			continue
		}

		courseMap[next.ID] = next
		graphMap[next.ID] = i
		inverseGraphMap[i] = next.ID

		prerequisites, err := next.getPreqs(a.DB)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		}
		for _, prereq := range prerequisites {
			dependencyMap[next.ID] = append(dependencyMap[next.ID], prereq.ID)
			queue = append(queue, prereq)
		}

		i++
	}

	courseGraph := graph.New(len(graphMap))
	for courseID, dependencies := range dependencyMap {
		courseGraphID := graphMap[courseID]
		for _, dependentID := range dependencies {
			dependentGraphID := graphMap[dependentID]
			log.Println(courseID, courseGraphID, dependentGraphID)
			courseGraph.Add(dependentGraphID, courseGraphID)
		}
	}

	log.Println(graphMap)
	log.Println("Acyclic: ", graph.Acyclic(courseGraph))
	order, ok := graph.TopSort(courseGraph)
	if !ok {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	result := []course{}
	for _, graphID := range order {
		courseID := inverseGraphMap[graphID]
		course := courseMap[courseID]
		result = append(result, course)
	}

	respondWithJSON(w, http.StatusOK, result)
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
    a.Router.HandleFunc("/prerequisites/{id}", a.getPrerequisites).Methods("GET")
	a.Router.HandleFunc("/full_prerequisites/{id}", a.getFullPrerequisites).Methods("GET")
}