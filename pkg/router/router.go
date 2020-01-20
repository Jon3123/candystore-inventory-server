package router

import (
	"net/http"

	"github.com/Jon3123/candystore-inventory-server/pkg/handler"
	"github.com/Jon3123/candystore-inventory-server/pkg/scontext"
	"github.com/gorilla/mux"
)

//Router router
type Router struct {
	BaseRoute string
	Ctx       *scontext.Context
	MRouter   *mux.Router
}

//NewRouter base
func NewRouter(context *scontext.Context, base string) *Router {
	r := &Router{BaseRoute: base, Ctx: context}
	router := mux.NewRouter().StrictSlash(true)
	r.MRouter = router
	r.Init()
	return r
}

//Init init
func (r *Router) Init() {
	h := &handler.Handler{Ctx: r.Ctx}
	r.MRouter.HandleFunc(r.BaseRoute+"/all", h.HandleAllRequest).Methods("GET")
	//	r.MRouter.HandleFunc(r.BaseRoute+"/Test/{id}", handleTest).Methods("GET")
	http.ListenAndServe(":8080", r.MRouter)
}
