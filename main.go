package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	loads "github.com/go-openapi/loads"
	"github.com/splathon/splathon-server/swagger/restapi"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
)

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewSplathonAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.EnabledListeners = []string{"http"}
	server.Port, _ = strconv.Atoi(os.Getenv("PORT"))

	server.ConfigureAPI()

	http.Handle("/", server.GetHandler())
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
