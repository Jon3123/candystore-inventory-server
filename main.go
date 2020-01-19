package main

import (
	"fmt"

	"github.com/Jon3123/candystore-inventory-server/storage"
)

func main() {
	// m := models.NewMapping("Source Test", "Target Test", 12)
	// fmt.Println(m)
	// str, err := json.Marshal(m)
	// if err != nil {
	// 	fmt.Println("ERROR")
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("not error")
	// 	fmt.Println(string(str))
	// }

	// data := []byte(`{"Source":"Source 3","Target":"Target 2","Count":124}`)

	// var mapping models.Mapping
	// json.Unmarshal(data, &mapping)

	// fmt.Println(mapping)

	store := storage.NewStorage("Test")

	store.AddValue("Hi2", "bye2")

	fmt.Println(store.GetValue("Hi3"))

}
