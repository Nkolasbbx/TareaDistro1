package main

import (
	"context"
	"log"
	"net"

	pb "gobierno/proto/grpc-server/proto" // Importa el paquete generado por protoc

	"google.golang.org/grpc"
)

// server es usado para implementar el servicio GobiernoServer.
type server struct {
	pb.UnimplementedComunicacionServer
}

// Lista de piratas (simulada).
var piratas = []*pb.Pirata{
	{Id: 1, Nombre: "Arlong", Recompensa: 20000000, NivelDePeligrosidad: "Bajo", Estado: "Buscado"},
	{Id: 2, Nombre: "Crocodile", Recompensa: 181000000, NivelDePeligrosidad: "Media", Estado: "Capturado"},
	{Id: 3, Nombre: "Eustass Kid", Recompensa: 515000000, NivelDePeligrosidad: "Alta", Estado: "Buscado"},
	{Id: 4, Nombre: "Trafalgar D. Water Law", Recompensa: 200000000, NivelDePeligrosidad: "Alta", Estado: "Buscado"},
	{Id: 5, Nombre: "Monkey D. Luffy", Recompensa: 1500000000, NivelDePeligrosidad: "Alta", Estado: "Buscado"},
}

// consultarListaPiratas implementa el método del servicio gRPC.
func (s *server) ConsultarListaPiratas(ctx context.Context, req *pb.Empty) (*pb.Respuesta, error) {
	return &pb.Respuesta{
		Piratas: piratas,
	}, nil
}

func main() {
	// Inicia el servidor gRPC
	listener, err := net.Listen("tcp", ":50051") // Puerto para el Gobierno
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterComunicacionServer(grpcServer, &server{})

	log.Println("Servidor del Gobierno escuchando en el puerto 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
