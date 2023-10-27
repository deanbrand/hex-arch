package rpc

import (
	"context"
	"github.com/deanbrand/hex-arch/internal/adapters/app/api"
	"github.com/deanbrand/hex-arch/internal/adapters/core/arithmetic"
	"github.com/deanbrand/hex-arch/internal/adapters/framework/left/grpc/pb"
	"github.com/deanbrand/hex-arch/internal/adapters/framework/right/db"
	"github.com/deanbrand/hex-arch/internal/ports"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"os"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	var err error
	lis = bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()

	// ports
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

	gRPCAdapter = NewAdapter(appAdapter)

	pb.RegisterArithmeticServiceServer(grpcServer, gRPCAdapter)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Error creating grpc server: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func getGRPCConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Error dialing grpc server: %v", err)
	}
	return conn
}

func TestGetAddition(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{
		A: 1,
		B: 2,
	}

	answer, err := client.GetAddition(ctx, params)
	if err != nil {
		t.Fatalf("Error getting addition: %v", err)
	}

	require.Equal(t, answer.Value, int32(3))
}

func TestGetSubtraction(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{
		A: 2,
		B: 1,
	}

	answer, err := client.GetSubtraction(ctx, params)
	if err != nil {
		t.Fatalf("Error getting subtraction: %v", err)
	}

	require.Equal(t, answer.Value, int32(1))
}

func TestGetMultiplication(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{
		A: 2,
		B: 3,
	}

	answer, err := client.GetMultiplication(ctx, params)
	if err != nil {
		t.Fatalf("Error getting multiplication: %v", err)
	}

	require.Equal(t, answer.Value, int32(6))
}

func TestGetDivision(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)

	params := &pb.OperationParameters{
		A: 6,
		B: 3,
	}

	answer, err := client.GetDivision(ctx, params)
	if err != nil {
		t.Fatalf("Error getting division: %v", err)
	}

	require.Equal(t, answer.Value, int32(2))
}
