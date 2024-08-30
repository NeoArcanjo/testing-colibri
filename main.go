package main

import (
	"github.com/NeoArcanjo/testing-colibri/src/application/controllers"
	"github.com/colibri-project-io/colibri-sdk-go"
	"github.com/colibri-project-io/colibri-sdk-go/pkg/web/restserver"
)

func init() {
	colibri.InitializeApp()
	//storage.Initialize() // uncomment if you use storage
	//cacheDB.Initialize() // uncomment if you use cache
	//sqlDB.Initialize() // uncomment if you use sql database
	//messaging.Initialize() // uncomment if you use messaging
	// testing-colibri
}

func main() {
	restserver.AddRoutes(controllers.NewDemoController().Routes())
	restserver.ListenAndServe()
}
