package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"

	v1 "music.com/musicservice/gen/go/musicservice/v1"
	svc "music.com/musicservice/services"

	// This is needed to enable the use of the grpc_cli tool
	"google.golang.org/grpc/reflection"
)

var (
	addr = flag.String("addr", ":9000", "Address to start the musicservice grpc server on.")
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

func main() {
	flag.Parse()
	startGRPCServer(*addr)
}
