package service

import (
	"encoding/json"
	"net/http"
	"part3/database"
	"strings"
)

//Ideally this would be an interface, but its overkill for since I don't need the extra abstraction here
type Service struct {
	Data database.DataKind
}

func New(data database.DataKind) *Service {
	return &Service{data}
}

func (s *Service) GetItem(w http.ResponseWriter, r *http.Request) {
	pieces := strings.Split(r.URL.Path, "/")
	if len(pieces) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	key := pieces[len(pieces)-1]
	value, ok := s.Data.Get(key)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(value))
}

func (s *Service) DeleteItem(w http.ResponseWriter, r *http.Request) {
	pieces := strings.Split(r.URL.Path, "/")
	if len(pieces) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	key := pieces[len(pieces)-1]

	err := s.Data.Delete(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

type Poster struct {
	Key   string
	Value string
}

// Accepts a Poster in json
func (s *Service) PostItem(w http.ResponseWriter, r *http.Request) {
	p := Poster{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil || p.Key == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(
			`Please post json of the form:

			{
				Key string 
				Value string
			}`,
		))
		return
	}
	err = s.Data.Update(p.Key, p.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Service) PutItem(w http.ResponseWriter, r *http.Request) {
	p := Poster{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(
			`Please post json of the form:
			{
				Key string 
				Value string
			}`,
		))
		return
	}
	err = s.Data.Update(p.Key, p.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
