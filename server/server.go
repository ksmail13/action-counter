package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ksmail13/action-counter/config"
	"github.com/ksmail13/action-counter/errors"
	"github.com/ksmail13/action-counter/model"
	"github.com/ksmail13/action-counter/repository"
)

type Server struct {
	Router *mux.Router
	Config *config.Config
	Repo   repository.Repository
}

func (server *Server) Initialize(config *config.Config) {
	server.Config = config
	server.Router = mux.NewRouter()
	server.setRouters()
}

func (server *Server) setRouters() {
	server.Get("/counter/{uuid:[a-z0-9-]+}", server.getCounter)
	server.Post("/counter", server.createCounter)
	server.Put("/counter/{uuid:[a-z0-9-]+}", server.updateCounter)
	server.Delete("/counter/{uuid:[a-z0-9-]+}", server.deleteCounter)
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

func ResponseJSON(w http.ResponseWriter, status int, payload interface{}) {
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

func RespondError(w http.ResponseWriter, code int, format string, args ...interface{}) {
	errMsg := fmt.Sprintf(format, args...)
	log.Println(errMsg)
	ResponseJSON(w, code, map[string]interface{}{"error": errMsg, "code": code})
}

func errorHandle(w http.ResponseWriter, e error) {
	if err, ok := e.(*errors.CodeError); ok {
		RespondError(w, err.Code(), err.Message())
		return
	}
	log.Printf("unexpect error %s", e)
	RespondError(w, http.StatusInternalServerError, "unexpect error")
}

func response(w http.ResponseWriter, v model.Counter, e error) {

	if e != nil {
		errorHandle(w, e)
		return
	}

	ResponseJSON(w, http.StatusOK, v)
}

func (server *Server) getCounter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	log.Printf("get UUID : %s", uuid)
	v, e := server.Repo.Get(uuid)
	response(w, v, e)
}

func (server *Server) createCounter(w http.ResponseWriter, r *http.Request) {
	rawbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "error while read request body [%s]", err)
		return
	}
	body := new(model.CounterCreate)
	err = json.Unmarshal(rawbody, &body)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "invalid request body [%s]", err)
		return
	}
	v, e := server.Repo.Set(body.Duration)
	response(w, v, e)
}

func (server *Server) updateCounter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]
	v, e := server.Repo.Increase(uuid)
	response(w, v, e)
}

func (server *Server) deleteCounter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	v, e := server.Repo.Decrease(uuid)
	response(w, v, e)
}
