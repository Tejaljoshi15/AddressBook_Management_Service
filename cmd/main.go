package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"sync"

	"infoBlox/config"
	"infoBlox/pkg/db"
	"infoBlox/pkg/db/storage"
	pb "infoBlox/pkg/gen/proto"
	"infoBlox/pkg/server"
	"infoBlox/pkg/usecases/user"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func startGRPCServer(port string, srv *server.AddressBookServer) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to set up TCP listener")
	}
	s := grpc.NewServer()
	pb.RegisterAddressBookServer(s, srv)

	log.Info().Msgf("gRPC server started listening on %v", port)
	if err := s.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("Failed to serve gRPC")
	}
}

func startHTTPGateway(httpAddress string, grpcAddress string) {
	ctx := context.Background()
	mux := http.NewServeMux()

	grpcGatewayMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterAddressBookHandlerFromEndpoint(ctx, grpcGatewayMux, grpcAddress, opts)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to register gRPC-Gateway handler")
	}

	mux.Handle("/", grpcGatewayMux)

	swaggerDir := http.Dir("C:/Users/tejal_joshi/Documents/Info_blox_Task2/infoblox_bootcamp/web/swaggerui/dist")
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(swaggerDir)))

	mux.Handle("/swagger.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "C:/Users/tejal_joshi/Documents/Info_blox_Task2/infoblox_bootcamp/pkg/gen/pkg/proto/addressbook.swagger.json")
	}))

	server := &http.Server{
		Addr:    httpAddress,
		Handler: mux,
	}
	log.Info().Msgf("HTTP gateway started listening on %v", httpAddress)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Failed to start HTTP server")
	}
}

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})

	config.LoadConfig()

	grpcPort := viper.GetString("grpc.port")
	httpPort := viper.GetString("http.port")

	storage := storage.New()
	dbInstance := db.New(storage)
	userService := user.New(dbInstance)
	srv := server.NewAddressBookServer(userService)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer recoverFromPanic()
		startGRPCServer(":"+grpcPort, srv)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer recoverFromPanic()
		startHTTPGateway(":"+httpPort, "localhost:"+grpcPort)
	}()

	wg.Wait()
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		log.Error().Msgf("Recovered from panic: %v", r)
	}
}
