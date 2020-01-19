package models

import (
	"encoding/json"
	"fmt"

	"github.com/Jon3123/candystore-inventory-server/pkg/server"
)

//Mapping mapping struct
type Mapping struct {
	Source string `json:"Source"`
	Target string `json:"Target"`
	Count  int    `json:"Count"`
}

//NewMapping Create a new mapping
func NewMapping(ctx *server.Context, source string, target string, count int) *Mapping {
	m := &Mapping{
		Source: source,
		Target: target,
		Count:  count,
	}

	val, err := json.Marshal(m)

	if err != nil {
		fmt.Println("ERROR Marshaling")
		return nil
	}
	ctx.Store.AddValue(source, string(val))
	return m
}
