package main

import (
	"github.com/deanbrand/hex-arch/internal/adapters/app/api"
	"github.com/deanbrand/hex-arch/internal/adapters/core/arithmetic"
	gRPC "github.com/deanbrand/hex-arch/internal/adapters/framework/left/grpc"
	"github.com/deanbrand/hex-arch/internal/adapters/framework/right/db"
	"github.com/deanbrand/hex-arch/internal/ports"
	"log"
	"os"
)

func main() {
	var err error

	//	ports
	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DB_SOURCE_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("Error creating db adapter: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
