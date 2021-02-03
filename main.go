package main

import (
	//"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func functionHandler(rw http.ResponseWriter, req *http.Request) {

	//jString := json.NewDecoder(req.Body)
	vars := mux.Vars(req)
	fmt.Println("Function name: ", vars["function"])
	if req.Method == "POST" {

	} else {

	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{function}", functionHandler)
	http.Handle("/", router)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}
