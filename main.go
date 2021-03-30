package main

import (
	"net/http"

	"github.com/ronrojro/webservice-using-go/controllers"
)

/*
func main() {
	u := models.User{
		ID:        2,
		FirstName: "Yali",
		LastName:  "Carvajal",
	}
	fmt.Println(u)

	port := 3000
	err := startWebServer(port)
	fmt.Println(err)
}

func startWebServer(port int) error {
	fmt.Println("Strating server...")
	//Do great ans brave
	fmt.Println("Server strated on port", port)
	return nil
}

*/

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
