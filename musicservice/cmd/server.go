package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"

	v1 "music.com/musicservice/gen/go/musicservice/v1"
	svc "music.com/musicservice/services"
)

var (
	addr = flag.String("addr", ":9000", "Address to start the musicservice grpc server on.")
)

func startGRPCServer(addr string) {
	// create new gRPC server
	server := grpc.NewServer()
	v1.RegisterSongServiceServer(server, svc.NewSongServiceServer())
	v1.RegisterAlbumServiceServer(server, svc.NewAlbumServiceServer())
	v1.RegisterArtistServiceServer(server, svc.NewArtistServiceServer())
	v1.RegisterLabelServiceServer(server, svc.NewLabelServiceServer())
	if l, err := net.Listen("tcp", addr); err != nil {
		log.Fatalf("error in listening on port %s: %v", addr, err)
	} else {
		// the gRPC server
		log.Printf("Starting grpc endpoint on %s:", addr)
		if err := server.Serve(l); err != nil {
			log.Fatal("unable to start server", err)
		}
	}
}

func main() {
	flag.Parse()
	startGRPCServer(*addr)
}
