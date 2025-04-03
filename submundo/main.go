package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	pb "submundo/proto/grpc-server/proto" // Importa el paquete generado por protoc

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedComunicacionServer
}

// venderPirataSubmundo implementa la lógica del Submundo
func (s *server) VenderPirataSubmundo(ctx context.Context, req *pb.VentaRequest) (*pb.VentaResponse, error) {
	// Genera un número aleatorio para determinar el resultado de la venta
	rand.Seed(time.Now().UnixNano())
	fraude := rand.Intn(100) < 35 // 35% de probabilidad de fraude

	if fraude {
		return &pb.VentaResponse{
			Exito:   false,
			Pago:    0,
			Mensaje: "Fraude: el Submundo no pagó por el pirata.",
		}, nil
	}

	// Calcula el pago entre 100% y 150% de la recompensa
	pago := 1000000 + rand.Int63n(500000) // Ejemplo de cálculo de pago aleatorio

	return &pb.VentaResponse{
		Exito:   true,
		Pago:    pago,
		Mensaje: "Venta exitosa en el Submundo.",
	}, nil
}

func main() {
	// Inicia el servidor gRPC
	listener, err := net.Listen("tcp", ":50052") // Puerto para el Submundo
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterComunicacionServer(grpcServer, &server{})

	log.Println("Servidor del Submundo escuchando en el puerto 50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
