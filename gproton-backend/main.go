package main

import (
	"fmt"
	"gproton-backend/grpcclient"
	"gproton-backend/handlers"

	"github.com/labstack/echo/v4"
)

var protoSet *grpcclient.ProtoSet

func main() {
	protoSet = grpcclient.LoadProtoSet("protoset/Uid-proto-20250907.protoset")
	fmt.Println("Successfully loaded protoset")

	e := echo.New()
	e.GET("/services", handlers.ListServiceHandler(protoSet))
	e.POST("/call", handlers.CallServiceHandler(protoSet))
	e.Logger.Fatal(e.Start(":8080"))
}
