package main

import (
	"context"
	trfmain "github-connector/pkg/rule/TRF_Main"
	util "github-connector/pkg/util"
)

import (
	_ "github-connector/docs"
)
import (
	"github-connector/pkg/middleware"
)
import (
	"github-connector/pkg/controller"
	"github-connector/pkg/xiLogger"
	"github-connector/sidecar"
	"log"
	"net"
)

func main() {

	cntxt := util.CustomContext{}
	cntxt.AppGoContext = context.Background()
	config := make(map[string]interface{})
	trfmain.TRF_Main(&cntxt, config)

	RunServer()
}

func RunServer() {
	port := middleware.GetGrpcPort()
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := middleware.GetGrpcServerConfigs()

	xiLogger.Log.Info("\n Server listening on port %v \n", ":"+port)
	sidecar.RegisterServiceServer(s, &controller.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
