package server

import (
	"encoding/json"
	"log"
	"net/http"

	"../config"
	"../model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	Config *config.Config
}

var countMap map[string]int = make(map[string]int)

func (server *Server) Initialize(config *config.Config) {
	server.Config = config
	server.Router = mux.NewRouter()
	server.setRouters()
}

func (server *Server) setRouters() {
	server.Get("/counter", server.CreateCounter)
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

func (server *Server) CreateCounter(w http.ResponseWriter, r *http.Request) {
	// TODO: 새로운 카운터 세션을 생성해서 반
	uuid := uuid.New()
	counter := model.Counter{UUID: uuid.String(), Count: 1}
	log.Printf("create %s", uuid)
	countMap[uuid.String()] = 1
	respondJSON(w, http.StatusOK, counter)
}

func (server *Server) UpdateCounter(w http.ResponseWriter, r *http.Request) {
	// TODO: 저장소에서 uuid로 카운터를 조회해서 카운터 증가 후 반
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	cnt := countMap[uuid]
	countMap[uuid] = cnt + 1
	counter := model.Counter{UUID: uuid, Count: cnt + 1}
	respondJSON(w, http.StatusOK, counter)
}

func (server *Server) DeleteCounter(w http.ResponseWriter, r *http.Request) {
	// TODO 카운터 삭제
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	cnt := countMap[uuid]
	countMap[uuid] = cnt - 1
	counter := model.Counter{UUID: uuid, Count: cnt - 1}
	respondJSON(w, http.StatusOK, counter)
}
