package main

import "fmt"
import "encoding/json"
import "github.com/Jon3123/candystore-inventory-server/models"

func main() {
	m := models.NewMapping("Source Test", "Target Test", 12)
	str, err := json.Marshal(m)
	if err == nil {
		fmt.Println(err)
	} else {
		fmt.Println("not error")
		fmt.Println(string(str))
	}
}
