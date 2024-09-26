package main

import (
	"fmt"

	"github.com/BodeXX/api-rest-go.git/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
