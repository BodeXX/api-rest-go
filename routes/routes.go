package routes

import (
	"log"
	"net/http"

	"github.com/BodeXX/api-rest-go.git/controllers"
)

func HandleRequest() {
	http.HandleFunc("/convert", controllers.Converter)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
