package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jon3123/candystore-inventory-server/pkg/scontext"
	"github.com/gorilla/mux"
)

//Handler handle
type Handler struct {
	Ctx *scontext.Context
}

//HandleAllRequest - handle request for all data
func (handle *Handler) HandleAllRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested")
	str, _ := json.Marshal(handle.Ctx.Store.GetAllMappings())
	fmt.Fprintf(w, string(str))
}

//HandleTest - Test
func (handle *Handler) HandleTest(w http.ResponseWriter, r *http.Request) {
	test := mux.Vars(r)["id"]
	fmt.Println(test)
	fmt.Fprintf(w, "Welcome home!")
}
