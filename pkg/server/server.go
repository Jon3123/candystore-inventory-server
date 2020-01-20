package server

import (
	"fmt"

	"github.com/Jon3123/candystore-inventory-server/pkg/router"
	"github.com/Jon3123/candystore-inventory-server/pkg/scontext"
	"github.com/Jon3123/candystore-inventory-server/storage"
)

//Server server
type Server struct {
	Ctx     *scontext.Context
	Running bool
}

//NewServer create new Server
func NewServer() *Server {
	s := &Server{Ctx: nil}
	return s
}

func (s *Server) loop() {
	fmt.Println("Looping")
	go router.NewRouter(s.Ctx, "/mapping")
	for s.Running {

	}
}

//Run run server
func (s *Server) Run() {
	//Creating storage
	store := storage.NewStorage("Mapping")

	//Create context
	context := &scontext.Context{
		Store: store}

	s.Ctx = context
	s.Running = true
	s.loop()
}
