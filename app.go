package main

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/fronzec/rest-api-go/config"
	. "github.com/fronzec/rest-api-go/dao"
	. "github.com/fronzec/rest-api-go/models"
)

var config = Config{}
var dao = MoviesDAO{}

// Handles LIST all movies
func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	movies, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

// Handles Find the detail of a Movie
func FindMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, movie)
}

// Handles POST a new movie
func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	// Next line assign a new ID  to new object then insert it on Mongo instance
	movie.ID = bson.NewObjectId()
	if err := dao.Insert(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, movie)
}

// Handles PUT a existing Movie
func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// Handles Delete a Movie
func DeleteMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	// Creates a new controller for each endpoint
	r := mux.NewRouter()
	r.HandleFunc("/v1/movies", AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/v1/movies", CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/v1/movies", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/v1/movies", DeleteMovieEndpoint).Methods("DELETE")
	r.HandleFunc("/v1/movies", FindMovieEndPoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
