package main

import (
	"fmt"
	"net/http"
	"part3/database"
	"part3/service"
)

//For a normal api there would be more endpoint seperation, and I would also use a framework not base http for routing
//For this tiny example though I have everyything going to /api for ease of use
//Post and Put takes a service.Poster json body 
func main() {
	data, err := database.Initialize("./database/db.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s := service.New(data)
	http.Handle("/api/",http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		switch r.Method {
		case http.MethodPost:
			s.PostItem(w,r)
		case http.MethodPut:
			s.PutItem(w,r)
		case http.MethodDelete:
			s.DeleteItem(w,r)
		case http.MethodGet:
			s.GetItem(w,r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}))
	http.ListenAndServe(":8000",nil)
}
