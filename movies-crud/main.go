package main

import (
  "fmt"
  "log"
  "encoding/json"
  "math/rand"
  "net/http"
  "strconv"

  "github.com/gorilla/mux"
)

type Movie struct {
  ID string `json:"id"`
  Isbn string `json:"isbn"`
  Title string `json:"title"`
  Director *Director `json:"director"`
}

type Director struct {
  FirstName string `json:"firstname"`
  LastName string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for _, item := range movies {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
      return
    }
  }
}

func createMovie(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var movie Movie
  _ = json.NewDecoder(r.Body).Decode(&movie)
  movie.ID = strconv.Itoa(rand.Intn(10000000))
  movies = append(movies, movie)
  json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range movies {
    if item.ID == params["id"] {
      movies = append(movies[:index], movies[index+1:]...)
      var movie Movie 
      _ = json.NewDecoder(r.Body).Decode(&movie)
      movie.ID = params["id"]
      movies = append(movies, movie)
      json.NewEncoder(w).Encode(movie)
      return
    }
  }
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)
  for index, item := range movies {
    if item.ID == params["id"] {
      movies = append(movies[:index], movies[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(movies)
}

func main() {
  router := mux.NewRouter()

  movies := append(movies, Movie{ID: "1", Isbn: "438227", Title: "Mission Impossible: Dead Reckoning-I", Director: &Director{FirstName: "Christopher", LastName: "McQuarrie"} })
  movies = append(movies, Movie{ID: "2", Isbn: "562288", Title: "Oppenheimer", Director: &Director{FirstName: "Christopher", LastName: "Nolan"} })
  router.HandleFunc("/movies", getMovies).Methods("GET")
  router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
  router.HandleFunc("/movies", createMovie).Methods("POST")
  router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
  router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

  fmt.Printf("Starting server at port 8000\n")
  log.Fatal(http.ListenAndServe(":8000", router))
}
