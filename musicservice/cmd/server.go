package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "music.com/musicservice/gen/go/musicservice/v1"
	svc "music.com/musicservice/services"

	// This is needed to enable the use of the grpc_cli tool
	"google.golang.org/grpc/reflection"
)

var (
	addr    = flag.String("addr", ":9000", "Address to start the grpc server on.")
	gw_addr = flag.String("gw_addr", ":8080", "Address to start the grpc gateway server on.")
)

func startGRPCServer(addr string) {
	// create new gRPC server
	server := grpc.NewServer()
	v1.RegisterSongServiceServer(server, svc.NewSongService(nil))
	v1.RegisterAlbumServiceServer(server, svc.NewAlbumService(nil))
	v1.RegisterArtistServiceServer(server, svc.NewArtistService(nil))
	v1.RegisterLabelServiceServer(server, svc.NewLabelService(nil))
	if l, err := net.Listen("tcp", addr); err != nil {
		log.Fatalf("error in listening on port %s: %v", addr, err)
	} else {
		// the gRPC server
		log.Printf("Starting grpc endpoint on %s:", addr)
		reflection.Register(server)
		if err := server.Serve(l); err != nil {
			log.Fatal("unable to start server", err)
		}
	}
}

func startGatewayServer(gw_addr, grpc_addr string) {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := v1.RegisterSongServiceHandlerFromEndpoint(ctx, mux, grpc_addr, opts)
	if err != nil {
		log.Fatal(err)
	}
	if err := v1.RegisterAlbumServiceHandlerFromEndpoint(ctx, mux, grpc_addr, opts); err != nil {
		log.Fatal(err)
	}
	if err := v1.RegisterArtistServiceHandlerFromEndpoint(ctx, mux, grpc_addr, opts); err != nil {
		log.Fatal(err)
	}
	if err := v1.RegisterLabelServiceHandlerFromEndpoint(ctx, mux, grpc_addr, opts); err != nil {
		log.Fatal(err)
	}

	log.Println("Starting grpc gateway server on: ", gw_addr)
	http.ListenAndServe(gw_addr, mux)
}

func main() {
	flag.Parse()
	go startGRPCServer(*addr)
	startGatewayServer(*gw_addr, *addr)
}
