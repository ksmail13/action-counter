package server

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/google/uuid"
)

type Server struct {
	Router *mux.Router
	Config *config.Config
}

func (server *Server) Initialize(config *Config) {
	server.Router = mux.NewRouter()
	server.setRouters()
}

func (server *Server) setRouters() {
	server.Router.Get("/uuid", server.CreateUUID)
}

func (server *server) CreateUUID() sdtring {
	return uuid.UUID.String();
}

// setRouters sets the all required routers
/*
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/projects", a.GetAllProjects)
	a.Post("/projects", a.CreateProject)
	a.Get("/projects/{title}", a.GetProject)
	a.Put("/projects/{title}", a.UpdateProject)
	a.Delete("/projects/{title}", a.DeleteProject)
	a.Put("/projects/{title}/archive", a.ArchiveProject)
	a.Delete("/projects/{title}/archive", a.RestoreProject)

	// Routing for handling the tasks
	a.Get("/projects/{title}/tasks", a.GetAllTasks)
	a.Post("/projects/{title}/tasks", a.CreateTask)
	a.Get("/projects/{title}/tasks/{id:[0-9]+}", a.GetTask)
	a.Put("/projects/{title}/tasks/{id:[0-9]+}", a.UpdateTask)
	a.Delete("/projects/{title}/tasks/{id:[0-9]+}", a.DeleteTask)
	a.Put("/projects/{title}/tasks/{id:[0-9]+}/complete", a.CompleteTask)
	a.Delete("/projects/{title}/tasks/{id:[0-9]+}/complete", a.UndoTask)
}
*/