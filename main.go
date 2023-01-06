package main
import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id`
	Isbn string `json:"isbn`
	Title string `json:"title"`
	Director *director `json:"director"`
}

type director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"firstname"`
}

var movies []Movie

func getmovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deletemovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,item := range movies{
		if item.ID==params["ID"]{
			movies=append(movies[:index],movies[index+1:]... )
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getmovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _,item:=range movies{
		if item.ID==params["ID"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}


func createmovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_=json.NewDecoder(r.Body).Decode(&movie)
	movie.ID= strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updatemovie(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,item := range movies{
		if item.ID==params["ID"]{
			movies=append(movies[:index],movies[index+1:]... )
			var movie Movie
	_=json.NewDecoder(r.Body).Decode(&movie)
	movie.ID= strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
		}
	}
}

func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1",Isbn: "1111",Title: "interstellar", Director: &director{Firstname: "harish",Lastname: "aadi"}})
	movies = append(movies, Movie{ID: "2",Isbn: "1112",Title: "fightclub", Director: &director{Firstname: "sai",Lastname: "pari"}})
	r.HandleFunc("/movies", getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getmovie).Methods("GET")
	r.HandleFunc("/movies", createmovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updatemovie).Methods("PUT")
	r.HandleFunc("/movie/{id}",deletemovie).Methods("DELETE")
	
	fmt.Printf("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080",r))
}