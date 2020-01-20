package scontext

import "github.com/Jon3123/candystore-inventory-server/storage"

//Context context
type Context struct {
	Store *storage.Storage
}
