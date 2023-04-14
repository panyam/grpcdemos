package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"strings"

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
	server := grpc.NewServer(
		grpc.UnaryInterceptor(EnsureAuthIsValid),
	)
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
	mux := runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {

			//
			// Step 2 - Extend the context
			//
			ctx = metadata.AppendToOutgoingContext(ctx)

			//
			// Step 3 - get the basic auth params
			//
			username, password, ok := request.BasicAuth()
			if !ok {
				return nil
			}
			md := metadata.Pairs()
			md.Append("MusicServiceUsername", username)
			md.Append("MusicServicePassword", password)
			return md
		}))

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(EnsureAuthExists),
	}
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

func EnsureAuthExists(ctx context.Context,
	method string, // Method to be invoked on the service (eg GetAlbums)
	req, // Request payload  (eg GetAlbumsRequest)
	reply interface{}, // Response payload (eg GetAlbumsResponse)
	cc *grpc.ClientConn, // the underlying connection to the service
	invoker grpc.UnaryInvoker, // The next handler
	opts ...grpc.CallOption) error {

	md, ok := metadata.FromOutgoingContext(ctx)
	if ok {
		usernames := md.Get("MusicServiceUsername")
		passwords := md.Get("MusicServicePassword")
		log.Println("UP: ", usernames, passwords)
		if len(usernames) > 0 && len(passwords) > 0 {
			username := strings.TrimSpace(usernames[0])
			password := strings.TrimSpace(passwords[0])
			if len(username) > 0 && len(password) > 0 {
				// All fine - just call the invoker
				return invoker(ctx, method, req, reply, cc, opts...)
			}
		}
	}
	return status.Error(codes.NotFound, "BasicAuth params not found")
}

func EnsureAuthIsValid(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		usernames := md.Get("MusicServiceUsername")
		passwords := md.Get("MusicServicePassword")
		log.Println("UP: ", usernames, passwords)
		if len(usernames) > 0 && len(passwords) > 0 {
			username := strings.TrimSpace(usernames[0])
			password := strings.TrimSpace(passwords[0])

			// Make sure you use better passwords than this!
			if len(username) > 0 && password == fmt.Sprintf("%s123", username) {
				// All fine - just call the invoker
				return handler(ctx, req)
			}
		}
	}
	return nil, status.Error(codes.NotFound, "Invalid username/password")
}

func main() {
	flag.Parse()
	go startGRPCServer(*addr)
	startGatewayServer(*gw_addr, *addr)
}
