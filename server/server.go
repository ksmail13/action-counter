package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"../config"
	"../model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	Config *config.Config
}

var countMap map[string]model.Counter = make(map[string]model.Counter)

func (server *Server) Initialize(config *config.Config) {
	server.Config = config
	server.Router = mux.NewRouter()
	server.setRouters()
}

func (server *Server) setRouters() {
	server.Get("/counter/{uuid:[a-z0-9-]+}", server.GetCounter)
	server.Post("/counter", server.CreateCounter)
	server.Put("/counter/{uuid:[a-z0-9-]+}", server.UpdateCounter)
	server.Delete("/counter/{uuid:[a-z0-9-]+}", server.DeleteCounter)
}

func (server *Server) Run(host string) {
	log.Fatal(http.ListenAndServe(host, server.Router))
}

func (server *Server) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	server.Router.HandleFunc(path, f).Methods("GET")
}

func (server *Server) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	server.Router.HandleFunc(path, f).Methods("POST")
}

func (server *Server) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	server.Router.HandleFunc(path, f).Methods("PUT")
}

func (server *Server) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	server.Router.HandleFunc(path, f).Methods("DELETE")
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

func (server *Server) GetCounter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	log.Printf("get UUID : %s", uuid)
	respondJSON(w, http.StatusOK, countMap[uuid])
}

func (server *Server) CreateCounter(w http.ResponseWriter, r *http.Request) {
	// TODO: 새로운 카운터 세션을 생성해서 반
	uuid := uuid.New()
	rawbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	body := new(model.CounterCreate)
	err = json.Unmarshal(rawbody, &body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	counter := model.Counter{UUID: uuid.String(), Count: 1, Life: time.Now().Add(body.Duration)}
	log.Printf("create %s", uuid)
	countMap[uuid.String()] = counter
	respondJSON(w, http.StatusOK, counter)
}

func (server *Server) UpdateCounter(w http.ResponseWriter, r *http.Request) {
	// TODO: 저장소에서 uuid로 카운터를 조회해서 카운터 증가 후 반
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	cnt := countMap[uuid]
	cnt.Count++

	respondJSON(w, http.StatusOK, cnt)
}

func (server *Server) DeleteCounter(w http.ResponseWriter, r *http.Request) {
	// TODO 카운터 삭제
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	cnt := countMap[uuid]
	cnt.Count--
	respondJSON(w, http.StatusOK, cnt)
}
