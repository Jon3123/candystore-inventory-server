package server

import (
	"github.com/Jon3123/candystore-inventory-server/storage"
)

type Server struct {
	context *Context
}

type Context struct {
	Store *storage.Storage
}

//NewServer create new Server
func NewServer() *Server {
	s := &Server{context: nil}
	return s
}

//Run run server
func (s *Server) Run() {
	//Creating storage
	store := storage.NewStorage("Mapping")

	//Create context
	context := &Context{
		Store: store}

	s.context = context
}
