package main

import (
<<<<<<< HEAD
	"github.com/airstrik/airstrik/src/route"
	"github.com/airstrik/gobase/pkg/server"
=======
	"encoding/json"
	"fmt"
	"github.com/airstrik/airstrik/pkg/core/model/engine"
	engine2 "github.com/airstrik/airstrik/pkg/engine"
	"github.com/tebeka/selenium"
	"io/ioutil"
	"log"
>>>>>>> 2a9a8f22cfa94fe637f4f98462e1c3f5603af100
	"os"
)
var (
	Env string
	Srv *server.Server
)

func init()  {
	Env = os.Getenv("ENV")
	Srv = server.GetServer("Access", 9001)
	//Srv
}


func main()  {
	// Register the Api for the Service from router
	Srv.RegisterRoute("/admin/{VersionId}", route.AdminRoute)
	Srv.RegisterRoute("/signin/{VersionId}", route.SignInRoute)
	Srv.RegisterRoute("/user/{VersionId}", route.UserRoute)
	// Public URL used in the Application
	Srv.RegisterRoute("/account", route.PublicRoute)

	Srv.Start()
}
