package main
import (
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json" 
)
type Post struct {
  ID string `json:"id"`
  Title string `json:"title"`
  Body string `json:"body"`
}
var posts []Post
func main() {
  router := mux.NewRouter()
  router.HandleFunc("/posts", getPosts).Methods("GET")

http.ListenAndServe(":8000", router)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("betinho")
}