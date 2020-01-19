package main

import "fmt"
import "encoding/json"
import "github.com/Jon3123/candy-store-inventory-server/models"

func main() {
	m := models.NewMapping("Source Test", "Target Test", 12)
	str, err := json.Marshal(m)
	fmt.Println("Hello, world.")
}
