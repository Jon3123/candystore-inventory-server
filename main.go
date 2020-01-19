package main

import (
	"fmt"

	"github.com/Jon3123/candystore-inventory-server/pkg/server"
)

func main() {
	//utils.Init()
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

	server := server.NewServer()
	server.Run()
	fmt.Println(server)

}
