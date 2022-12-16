package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/solnsumei/go-grpc-demo/notes"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	// parse arguments from the command line
	// this lets us define the port for the server
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Instantiate the server
	server := grpc.NewServer()

	// Register server methods
	notes.RegisterNotesServer(server, &NotesServer{})

	log.Printf("Server listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
